package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &SearchDisplayResource{}
	_ resource.ResourceWithConfigure   = &SearchDisplayResource{}
	_ resource.ResourceWithImportState = &SearchDisplayResource{}
)

// searchDisplayTypes lists every SearchKit display type, from
// SearchDisplay.getFields(loadOptions) on a live instance. They differ only
// in the shape of `settings`, so a single resource with `settings` as a JSON
// string covers all of them.
var searchDisplayTypes = []string{"table", "list", "grid", "tree", "autocomplete", "entity", "batch"}

// searchDisplaySelect is passed to every GetByID call for this resource.
// is_autocomplete_default is declared with data_type "Extra" in
// SearchDisplay.getFields — verified against a live instance that CiviCRM
// omits "Extra" fields from the response entirely unless explicitly
// selected, even though the field itself always has a value (false by
// default). Passing select=nil (the default for other resources in this
// provider, which return every field without it) silently drops the
// attribute from every Create/Read/Update response, which read as
// permanent post-import drift.
var searchDisplaySelect = []string{
	"id", "name", "label", "saved_search_id", "type", "settings", "acl_bypass", "is_autocomplete_default",
}

// stripColumnSpecs removes the "spec" key CiviCRM auto-generates on each
// settings.columns[] entry (field metadata: data_type, input_type, fk
// target, etc.) for some display types (observed on batch and entity;
// verified absent on table). Config never supplies "spec" — the resource's
// own decodeJSONAttribute call sends exactly what the user wrote — so
// leaving it in the value read back from CiviCRM makes every plan after a
// batch/entity create look like unmanaged drift. Only "spec" is touched;
// any other CiviCRM-added top-level or per-column keys pass through
// unchanged, since only this one key is a known, reliably-reproduced
// generated artifact rather than user-configurable data.
func stripColumnSpecs(settings any) any {
	obj, ok := settings.(map[string]any)
	if !ok {
		return settings
	}
	columnsRaw, ok := obj["columns"]
	if !ok {
		return settings
	}
	columns, ok := columnsRaw.([]any)
	if !ok {
		return settings
	}
	cleaned := make([]any, len(columns))
	for i, colRaw := range columns {
		col, ok := colRaw.(map[string]any)
		if !ok {
			cleaned[i] = colRaw
			continue
		}
		if _, hasSpec := col["spec"]; !hasSpec {
			cleaned[i] = colRaw
			continue
		}
		colCopy := make(map[string]any, len(col)-1)
		for k, v := range col {
			if k != "spec" {
				colCopy[k] = v
			}
		}
		cleaned[i] = colCopy
	}
	objCopy := make(map[string]any, len(obj))
	for k, v := range obj {
		objCopy[k] = v
	}
	objCopy["columns"] = cleaned
	return objCopy
}

type SearchDisplayResource struct {
	client *Client
}

type SearchDisplayResourceModel struct {
	ID                    types.Int64          `tfsdk:"id"`
	Name                  types.String         `tfsdk:"name"`
	Label                 types.String         `tfsdk:"label"`
	SavedSearchID         types.Int64          `tfsdk:"saved_search_id"`
	Type                  types.String         `tfsdk:"type"`
	Settings              jsontypes.Normalized `tfsdk:"settings"`
	ACLBypass             types.Bool           `tfsdk:"acl_bypass"`
	IsAutocompleteDefault types.Bool           `tfsdk:"is_autocomplete_default"`
}

func NewSearchDisplayResource() resource.Resource {
	return &SearchDisplayResource{}
}

func (r *SearchDisplayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_search_display"
}

func (r *SearchDisplayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a CiviCRM SearchKit SearchDisplay — the columns, labels, sort order, " +
			"pager, row links and action buttons shown for a civicrm_saved_search. Covers all SearchKit " +
			"display types (table, list, grid, tree, autocomplete, entity, batch): they differ only in " +
			"the shape of `settings`, which this resource stores as a JSON string.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier of the search display.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Machine name of the search display. If omitted, CiviCRM assigns one derived " +
					"from the label.",
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"label": schema.StringAttribute{
				Description: "Administrative label for the display.",
				Required:    true,
			},
			"saved_search_id": schema.Int64Attribute{
				Description: "ID of the civicrm_saved_search this display renders.",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "Display type. One of: " + fmt.Sprintf("%q", searchDisplayTypes) + ".",
				Required:    true,
				Validators: []validator.String{
					stringOneOf(searchDisplayTypes...),
				},
			},
			"settings": schema.StringAttribute{
				Description: "Display settings (columns, sort, limit, pager, row links, action buttons, " +
					"etc.) as a JSON string (e.g. from jsonencode(...)). The exact shape depends on `type` — " +
					"see the CiviCRM SearchKit documentation and the resource docs for worked examples. " +
					"CiviCRM API4 expects the decoded structure, not a pre-encoded string; this attribute " +
					"handles that encoding/decoding automatically. CiviCRM augments each `columns[]` entry " +
					"with a generated `spec` sub-object server-side for some display types (e.g. batch, " +
					"entity); this attribute strips `spec` back out when reading state, so config and state " +
					"stay comparable.",
				Optional:   true,
				Computed:   true,
				CustomType: jsontypes.NormalizedType{},
			},
			"acl_bypass": schema.BoolAttribute{
				Description: "Whether this display bypasses ACL permission checks. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"is_autocomplete_default": schema.BoolAttribute{
				Description: "Whether this is the default autocomplete display for its saved search's entity.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
		},
	}
}

func (r *SearchDisplayResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SearchDisplayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SearchDisplayResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating SearchDisplay", map[string]any{"label": plan.Label.ValueString()})

	// The "entity" display type registers a virtual DB entity named after
	// this display, so CiviCRM requires an explicit name for it rather than
	// deriving one from label as it does for other types — omitting it fails
	// deep inside CiviCRM with an unhelpful TypeError ("$displayName must be
	// of type string, null given"). Catch it here with a clear message.
	if plan.Type.ValueString() == "entity" && (plan.Name.IsNull() || plan.Name.IsUnknown() || plan.Name.ValueString() == "") {
		resp.Diagnostics.AddAttributeError(
			path.Root("name"),
			"name is required for type \"entity\"",
			"SearchDisplay type \"entity\" registers a virtual DB entity named after this display, "+
				"so CiviCRM requires an explicit name (it is not derived from label as it is for other "+
				"display types). Set the name attribute.",
		)
		return
	}

	values := map[string]any{
		"label":                   plan.Label.ValueString(),
		"saved_search_id":         plan.SavedSearchID.ValueInt64(),
		"type":                    plan.Type.ValueString(),
		"acl_bypass":              plan.ACLBypass.ValueBool(),
		"is_autocomplete_default": plan.IsAutocompleteDefault.ValueBool(),
	}

	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		values["name"] = plan.Name.ValueString()
	}
	if !plan.Settings.IsNull() && !plan.Settings.IsUnknown() {
		decoded, err := decodeJSONAttribute(plan.Settings.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("settings"),
				"Invalid settings JSON",
				"settings must be valid JSON (e.g. produced by jsonencode(...)): "+err.Error(),
			)
			return
		}
		values["settings"] = decoded
	}

	result, err := r.client.Create("SearchDisplay", values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating SearchDisplay",
			"Could not create SearchDisplay: "+err.Error(),
		)
		return
	}

	if createdID, ok := GetInt64(result, "id"); ok {
		if fullResult, err2 := r.client.GetByID("SearchDisplay", createdID, searchDisplaySelect); err2 == nil {
			result = fullResult
		}
	}

	r.mapResultToState(result, &plan)

	tflog.Debug(ctx, "Created SearchDisplay", map[string]any{"id": plan.ID.ValueInt64()})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SearchDisplayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SearchDisplayResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading SearchDisplay", map[string]any{"id": state.ID.ValueInt64()})

	result, err := r.client.GetByID("SearchDisplay", state.ID.ValueInt64(), searchDisplaySelect)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading SearchDisplay",
			"Could not read SearchDisplay ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	r.mapResultToState(result, &state)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *SearchDisplayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SearchDisplayResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var state SearchDisplayResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating SearchDisplay", map[string]any{"id": state.ID.ValueInt64()})

	values := map[string]any{
		"label":                   plan.Label.ValueString(),
		"saved_search_id":         plan.SavedSearchID.ValueInt64(),
		"type":                    plan.Type.ValueString(),
		"acl_bypass":              plan.ACLBypass.ValueBool(),
		"is_autocomplete_default": plan.IsAutocompleteDefault.ValueBool(),
	}

	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		values["name"] = plan.Name.ValueString()
	} else {
		values["name"] = nil
	}
	if !plan.Settings.IsNull() && !plan.Settings.IsUnknown() {
		decoded, err := decodeJSONAttribute(plan.Settings.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("settings"),
				"Invalid settings JSON",
				"settings must be valid JSON (e.g. produced by jsonencode(...)): "+err.Error(),
			)
			return
		}
		values["settings"] = decoded
	} else {
		values["settings"] = nil
	}

	_, err := r.client.Update("SearchDisplay", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating SearchDisplay",
			"Could not update SearchDisplay ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	plan.ID = state.ID

	result, err := r.client.GetByID("SearchDisplay", state.ID.ValueInt64(), searchDisplaySelect)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading SearchDisplay after update",
			"Could not re-read SearchDisplay ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}
	r.mapResultToState(result, &plan)

	tflog.Debug(ctx, "Updated SearchDisplay", map[string]any{"id": plan.ID.ValueInt64()})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SearchDisplayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SearchDisplayResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting SearchDisplay", map[string]any{"id": state.ID.ValueInt64()})

	err := r.client.Delete("SearchDisplay", state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting SearchDisplay",
			"Could not delete SearchDisplay ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "Deleted SearchDisplay", map[string]any{"id": state.ID.ValueInt64()})
}

func (r *SearchDisplayResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (r *SearchDisplayResource) mapResultToState(result map[string]any, model *SearchDisplayResourceModel) {
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

	if v, ok := GetInt64(result, "saved_search_id"); ok {
		model.SavedSearchID = types.Int64Value(v)
	} else {
		model.SavedSearchID = types.Int64Null()
	}

	if v, ok := GetString(result, "type"); ok && v != "" {
		model.Type = types.StringValue(v)
	} else {
		model.Type = types.StringNull()
	}

	if raw, ok := result["settings"]; ok && raw != nil {
		encoded, err := encodeJSONAttribute(stripColumnSpecs(raw))
		if err == nil && encoded != "" && encoded != "null" {
			model.Settings = jsontypes.NewNormalizedValue(encoded)
		} else {
			model.Settings = jsontypes.NewNormalizedNull()
		}
	} else {
		model.Settings = jsontypes.NewNormalizedNull()
	}

	if v, ok := GetBool(result, "acl_bypass"); ok {
		model.ACLBypass = types.BoolValue(v)
	}

	if v, ok := GetBool(result, "is_autocomplete_default"); ok {
		model.IsAutocompleteDefault = types.BoolValue(v)
	}
}
