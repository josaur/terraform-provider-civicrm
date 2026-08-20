package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// decodeJSONAttribute unmarshals a JSON-string attribute value (e.g. from
// jsonencode(...) in a Terraform config) into a Go structure suitable for
// passing to CiviCRM API4. SavedSearch.api_params (SERIALIZE_JSON) and
// form_values (SERIALIZE_PHP, but likewise API4-JSON-shaped) both expect
// the *decoded* structure, not a pre-encoded JSON string: API4 performs its
// own encoding, so passing an already-encoded string double-encodes it and
// CiviCRM stores a JSON array containing a JSON string instead of the
// object SearchKit expects (see GitHub issue #10).
func decodeJSONAttribute(raw string) (any, error) {
	var decoded any
	if err := json.Unmarshal([]byte(raw), &decoded); err != nil {
		return nil, err
	}
	return decoded, nil
}

// encodeJSONAttribute marshals a value returned by CiviCRM API4 (already a
// native Go structure, not a JSON string) back into the JSON string stored
// in Terraform state, so config written with jsonencode(...) round-trips
// without drift.
func encodeJSONAttribute(v any) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

var (
	_ resource.Resource                = &SavedSearchResource{}
	_ resource.ResourceWithConfigure   = &SavedSearchResource{}
	_ resource.ResourceWithImportState = &SavedSearchResource{}
)

type SavedSearchResource struct {
	client *Client
}

type SavedSearchResourceModel struct {
	ID             types.Int64          `tfsdk:"id"`
	Name           types.String         `tfsdk:"name"`
	Label          types.String         `tfsdk:"label"`
	FormValues     jsontypes.Normalized `tfsdk:"form_values"`
	MappingID      types.Int64          `tfsdk:"mapping_id"`
	SearchCustomID types.Int64          `tfsdk:"search_custom_id"`
	APIEntity      types.String         `tfsdk:"api_entity"`
	APIParams      jsontypes.Normalized `tfsdk:"api_params"`
	CreatedID      types.Int64          `tfsdk:"created_id"`
	ModifiedID     types.Int64          `tfsdk:"modified_id"`
	ExpiresDate    types.String         `tfsdk:"expires_date"`
	CreatedDate    types.String         `tfsdk:"created_date"`
	ModifiedDate   types.String         `tfsdk:"modified_date"`
	Description    types.String         `tfsdk:"description"`
	IsTemplate     types.Bool           `tfsdk:"is_template"`
}

func NewSavedSearchResource() resource.Resource {
	return &SavedSearchResource{}
}

func (r *SavedSearchResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_saved_search"
}

func (r *SavedSearchResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a CiviCRM SavedSearch.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier of the saved_search.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Unique name of saved search.",
				Required:    true,
			},
			"label": schema.StringAttribute{
				Description: "Administrative label for search.",
				Optional:    true,
				Computed:    true,
			},
			"form_values": schema.StringAttribute{
				Description: "Submitted form values for this search, as a JSON string (e.g. from jsonencode(...)). " +
					"CiviCRM API4 expects the decoded structure, not a pre-encoded string; this attribute handles " +
					"that encoding/decoding automatically.",
				Optional:   true,
				Computed:   true,
				CustomType: jsontypes.NormalizedType{},
			},
			"mapping_id": schema.Int64Attribute{
				Description: "Foreign key to civicrm_mapping used for saved search-builder searches..",
				Optional:    true,
				Computed:    true,
			},
			"search_custom_id": schema.Int64Attribute{
				Description: "Foreign key to civicrm_option value table used for saved custom searches..",
				Optional:    true,
				Computed:    true,
			},
			"api_entity": schema.StringAttribute{
				Description: "Entity name for API based search.",
				Optional:    true,
				Computed:    true,
			},
			"api_params": schema.StringAttribute{
				Description: "Parameters for API based search (select/where/join/etc.), as a JSON string (e.g. " +
					"from jsonencode(...)). CiviCRM API4 expects the decoded structure, not a pre-encoded string; " +
					"this attribute handles that encoding/decoding automatically.",
				Optional:   true,
				Computed:   true,
				CustomType: jsontypes.NormalizedType{},
			},
			"created_id": schema.Int64Attribute{
				Description: "FK to contact table..",
				Optional:    true,
				Computed:    true,
			},
			"modified_id": schema.Int64Attribute{
				Description: "FK to contact table..",
				Optional:    true,
				Computed:    true,
			},
			"expires_date": schema.StringAttribute{
				Description: "Optional date after which the search is not needed.",
				Optional:    true,
				Computed:    true,
			},
			"created_date": schema.StringAttribute{
				Description: "When the search was created..",
				Optional:    true,
				Computed:    true,
			},
			"modified_date": schema.StringAttribute{
				Description: "When the search was last modified..",
				Optional:    true,
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "Saved Search Description.",
				Optional:    true,
				Computed:    true,
			},
			"is_template": schema.BoolAttribute{
				Description: "Search templates are used as a starting point for building new searches.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
		},
	}
}

func (r *SavedSearchResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.client = client
}

func (r *SavedSearchResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SavedSearchResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating SavedSearch", map[string]any{"name": plan.Name.ValueString()})

	values := map[string]any{
		"name":        plan.Name.ValueString(),
		"is_template": plan.IsTemplate.ValueBool(),
	}

	if !plan.Label.IsNull() && !plan.Label.IsUnknown() {
		values["label"] = plan.Label.ValueString()
	}
	if !plan.FormValues.IsNull() && !plan.FormValues.IsUnknown() {
		decoded, err := decodeJSONAttribute(plan.FormValues.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("form_values"),
				"Invalid form_values JSON",
				"form_values must be valid JSON (e.g. produced by jsonencode(...)): "+err.Error(),
			)
			return
		}
		values["form_values"] = decoded
	}
	if !plan.MappingID.IsNull() && !plan.MappingID.IsUnknown() {
		values["mapping_id"] = plan.MappingID.ValueInt64()
	}
	if !plan.SearchCustomID.IsNull() && !plan.SearchCustomID.IsUnknown() {
		values["search_custom_id"] = plan.SearchCustomID.ValueInt64()
	}
	if !plan.APIEntity.IsNull() && !plan.APIEntity.IsUnknown() {
		values["api_entity"] = plan.APIEntity.ValueString()
	}
	if !plan.APIParams.IsNull() && !plan.APIParams.IsUnknown() {
		decoded, err := decodeJSONAttribute(plan.APIParams.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("api_params"),
				"Invalid api_params JSON",
				"api_params must be valid JSON (e.g. produced by jsonencode(...)): "+err.Error(),
			)
			return
		}
		values["api_params"] = decoded
	}
	if !plan.CreatedID.IsNull() && !plan.CreatedID.IsUnknown() {
		values["created_id"] = plan.CreatedID.ValueInt64()
	}
	if !plan.ModifiedID.IsNull() && !plan.ModifiedID.IsUnknown() {
		values["modified_id"] = plan.ModifiedID.ValueInt64()
	}
	if !plan.ExpiresDate.IsNull() && !plan.ExpiresDate.IsUnknown() {
		values["expires_date"] = plan.ExpiresDate.ValueString()
	}
	if !plan.CreatedDate.IsNull() && !plan.CreatedDate.IsUnknown() {
		values["created_date"] = plan.CreatedDate.ValueString()
	}
	if !plan.ModifiedDate.IsNull() && !plan.ModifiedDate.IsUnknown() {
		values["modified_date"] = plan.ModifiedDate.ValueString()
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		values["description"] = plan.Description.ValueString()
	}

	result, err := r.client.Create("SavedSearch", values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating SavedSearch",
			"Could not create SavedSearch: "+err.Error(),
		)
		return
	}

	if createdID, ok := GetInt64(result, "id"); ok {
		if fullResult, err2 := r.client.GetByID("SavedSearch", createdID, nil); err2 == nil {
			result = fullResult
		}
	}

	r.mapResultToState(result, &plan)

	tflog.Debug(ctx, "Created SavedSearch", map[string]any{"id": plan.ID.ValueInt64()})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SavedSearchResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SavedSearchResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading SavedSearch", map[string]any{"id": state.ID.ValueInt64()})

	result, err := r.client.GetByID("SavedSearch", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading SavedSearch",
			"Could not read SavedSearch ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	r.mapResultToState(result, &state)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *SavedSearchResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SavedSearchResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var state SavedSearchResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating SavedSearch", map[string]any{"id": state.ID.ValueInt64()})

	values := map[string]any{
		"name":        plan.Name.ValueString(),
		"is_template": plan.IsTemplate.ValueBool(),
	}

	if !plan.Label.IsNull() && !plan.Label.IsUnknown() {
		values["label"] = plan.Label.ValueString()
	} else {
		values["label"] = nil
	}
	if !plan.FormValues.IsNull() && !plan.FormValues.IsUnknown() {
		decoded, err := decodeJSONAttribute(plan.FormValues.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("form_values"),
				"Invalid form_values JSON",
				"form_values must be valid JSON (e.g. produced by jsonencode(...)): "+err.Error(),
			)
			return
		}
		values["form_values"] = decoded
	} else {
		values["form_values"] = nil
	}
	if !plan.MappingID.IsNull() && !plan.MappingID.IsUnknown() {
		values["mapping_id"] = plan.MappingID.ValueInt64()
	} else {
		values["mapping_id"] = nil
	}
	if !plan.SearchCustomID.IsNull() && !plan.SearchCustomID.IsUnknown() {
		values["search_custom_id"] = plan.SearchCustomID.ValueInt64()
	} else {
		values["search_custom_id"] = nil
	}
	if !plan.APIEntity.IsNull() && !plan.APIEntity.IsUnknown() {
		values["api_entity"] = plan.APIEntity.ValueString()
	} else {
		values["api_entity"] = nil
	}
	if !plan.APIParams.IsNull() && !plan.APIParams.IsUnknown() {
		decoded, err := decodeJSONAttribute(plan.APIParams.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("api_params"),
				"Invalid api_params JSON",
				"api_params must be valid JSON (e.g. produced by jsonencode(...)): "+err.Error(),
			)
			return
		}
		values["api_params"] = decoded
	} else {
		values["api_params"] = nil
	}
	if !plan.CreatedID.IsNull() && !plan.CreatedID.IsUnknown() {
		values["created_id"] = plan.CreatedID.ValueInt64()
	} else {
		values["created_id"] = nil
	}
	if !plan.ModifiedID.IsNull() && !plan.ModifiedID.IsUnknown() {
		values["modified_id"] = plan.ModifiedID.ValueInt64()
	} else {
		values["modified_id"] = nil
	}
	if !plan.ExpiresDate.IsNull() && !plan.ExpiresDate.IsUnknown() {
		values["expires_date"] = plan.ExpiresDate.ValueString()
	} else {
		values["expires_date"] = nil
	}
	if !plan.CreatedDate.IsNull() && !plan.CreatedDate.IsUnknown() {
		values["created_date"] = plan.CreatedDate.ValueString()
	} else {
		values["created_date"] = nil
	}
	if !plan.ModifiedDate.IsNull() && !plan.ModifiedDate.IsUnknown() {
		values["modified_date"] = plan.ModifiedDate.ValueString()
	} else {
		values["modified_date"] = nil
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		values["description"] = plan.Description.ValueString()
	} else {
		values["description"] = nil
	}

	_, err := r.client.Update("SavedSearch", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating SavedSearch",
			"Could not update SavedSearch ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	plan.ID = state.ID

	result, err := r.client.GetByID("SavedSearch", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading SavedSearch after update",
			"Could not re-read SavedSearch ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}
	r.mapResultToState(result, &plan)

	tflog.Debug(ctx, "Updated SavedSearch", map[string]any{"id": plan.ID.ValueInt64()})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SavedSearchResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SavedSearchResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting SavedSearch", map[string]any{"id": state.ID.ValueInt64()})

	err := r.client.Delete("SavedSearch", state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting SavedSearch",
			"Could not delete SavedSearch ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "Deleted SavedSearch", map[string]any{"id": state.ID.ValueInt64()})
}

func (r *SavedSearchResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid import ID",
			"Could not parse import ID as integer: "+err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (r *SavedSearchResource) mapResultToState(result map[string]any, model *SavedSearchResourceModel) {
	if id, ok := GetInt64(result, "id"); ok {
		model.ID = types.Int64Value(id)
	}

	if v, ok := GetString(result, "name"); ok && v != "" {
		model.Name = types.StringValue(v)
	} else {
		model.Name = types.StringNull()
	}

	if v, ok := GetString(result, "label"); ok && v != "" {
		model.Label = types.StringValue(v)
	} else {
		model.Label = types.StringNull()
	}

	if raw, ok := result["form_values"]; ok && raw != nil {
		encoded, err := encodeJSONAttribute(raw)
		if err == nil && encoded != "" && encoded != "null" {
			model.FormValues = jsontypes.NewNormalizedValue(encoded)
		} else {
			model.FormValues = jsontypes.NewNormalizedNull()
		}
	} else {
		model.FormValues = jsontypes.NewNormalizedNull()
	}

	if v, ok := GetInt64(result, "mapping_id"); ok {
		model.MappingID = types.Int64Value(v)
	} else {
		model.MappingID = types.Int64Null()
	}

	if v, ok := GetInt64(result, "search_custom_id"); ok {
		model.SearchCustomID = types.Int64Value(v)
	} else {
		model.SearchCustomID = types.Int64Null()
	}

	if v, ok := GetString(result, "api_entity"); ok && v != "" {
		model.APIEntity = types.StringValue(v)
	} else {
		model.APIEntity = types.StringNull()
	}

	if raw, ok := result["api_params"]; ok && raw != nil {
		encoded, err := encodeJSONAttribute(raw)
		if err == nil && encoded != "" && encoded != "null" {
			model.APIParams = jsontypes.NewNormalizedValue(encoded)
		} else {
			model.APIParams = jsontypes.NewNormalizedNull()
		}
	} else {
		model.APIParams = jsontypes.NewNormalizedNull()
	}

	if v, ok := GetInt64(result, "created_id"); ok {
		model.CreatedID = types.Int64Value(v)
	} else {
		model.CreatedID = types.Int64Null()
	}

	if v, ok := GetInt64(result, "modified_id"); ok {
		model.ModifiedID = types.Int64Value(v)
	} else {
		model.ModifiedID = types.Int64Null()
	}

	if v, ok := GetString(result, "expires_date"); ok && v != "" {
		model.ExpiresDate = types.StringValue(v)
	} else {
		model.ExpiresDate = types.StringNull()
	}

	if v, ok := GetString(result, "created_date"); ok && v != "" {
		model.CreatedDate = types.StringValue(v)
	} else {
		model.CreatedDate = types.StringNull()
	}

	if v, ok := GetString(result, "modified_date"); ok && v != "" {
		model.ModifiedDate = types.StringValue(v)
	} else {
		model.ModifiedDate = types.StringNull()
	}

	if v, ok := GetString(result, "description"); ok && v != "" {
		model.Description = types.StringValue(v)
	} else {
		model.Description = types.StringNull()
	}

	if v, ok := GetBool(result, "is_template"); ok {
		model.IsTemplate = types.BoolValue(v)
	}

}
