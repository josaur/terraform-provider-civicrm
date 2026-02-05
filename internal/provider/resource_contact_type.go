package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &ContactTypeResource{}
	_ resource.ResourceWithConfigure   = &ContactTypeResource{}
	_ resource.ResourceWithImportState = &ContactTypeResource{}
)

// ContactTypeResource manages contact types in CiviCRM.
type ContactTypeResource struct {
	client *Client
}

type ContactTypeResourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Label       types.String `tfsdk:"label"`
	Description types.String `tfsdk:"description"`
	ImageURL    types.String `tfsdk:"image_url"`
	Icon        types.String `tfsdk:"icon"`
	ParentID    types.Int64  `tfsdk:"parent_id"`
	IsActive    types.Bool   `tfsdk:"is_active"`
	IsReserved  types.Bool   `tfsdk:"is_reserved"`
}

func NewContactTypeResource() resource.Resource {
	return &ContactTypeResource{}
}

func (r *ContactTypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contact_type"
}

func (r *ContactTypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages CiviCRM Contact Types and subtypes (Individual, Organization, Household, and custom subtypes).",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier of the contact type.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The machine name of the contact type (must be unique).",
				Required:    true,
			},
			"label": schema.StringAttribute{
				Description: "The display label of the contact type.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the contact type.",
				Optional:    true,
			},
			"image_url": schema.StringAttribute{
				Description: "URL to an image for this contact type.",
				Optional:    true,
			},
			"icon": schema.StringAttribute{
				Description: "FontAwesome icon class (e.g., 'fa-user', 'fa-building').",
				Optional:    true,
			},
			"parent_id": schema.Int64Attribute{
				Description: "The parent contact type ID. Use 1 for Individual subtypes, 2 for Household subtypes, 3 for Organization subtypes.",
				Optional:    true,
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether the contact type is active. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"is_reserved": schema.BoolAttribute{
				Description: "Whether this is a reserved system contact type. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
		},
	}
}

func (r *ContactTypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ContactTypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ContactTypeResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating contact type", map[string]any{
		"name": plan.Name.ValueString(),
	})

	// Build values for API call
	values := map[string]any{
		"name":        plan.Name.ValueString(),
		"label":       plan.Label.ValueString(),
		"is_active":   plan.IsActive.ValueBool(),
		"is_reserved": plan.IsReserved.ValueBool(),
	}

	if !plan.Description.IsNull() {
		values["description"] = plan.Description.ValueString()
	}

	if !plan.ImageURL.IsNull() {
		values["image_URL"] = plan.ImageURL.ValueString()
	}

	if !plan.Icon.IsNull() {
		values["icon"] = plan.Icon.ValueString()
	}

	if !plan.ParentID.IsNull() {
		values["parent_id"] = plan.ParentID.ValueInt64()
	}

	// Call API
	result, err := r.client.Create("ContactType", values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating contact type",
			"Could not create contact type, unexpected error: "+err.Error(),
		)
		return
	}

	// Update state with response
	r.mapResponseToModel(result, &plan)

	tflog.Debug(ctx, "Created contact type", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ContactTypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ContactTypeResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading contact type", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	result, err := r.client.GetByID("ContactType", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading contact type",
			"Could not read contact type ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Update state
	r.mapResponseToModel(result, &state)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *ContactTypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ContactTypeResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ContactTypeResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating contact type", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	// Build values for API call
	values := map[string]any{
		"name":        plan.Name.ValueString(),
		"label":       plan.Label.ValueString(),
		"is_active":   plan.IsActive.ValueBool(),
		"is_reserved": plan.IsReserved.ValueBool(),
	}

	if !plan.Description.IsNull() {
		values["description"] = plan.Description.ValueString()
	} else {
		values["description"] = nil
	}

	if !plan.ImageURL.IsNull() {
		values["image_URL"] = plan.ImageURL.ValueString()
	} else {
		values["image_URL"] = nil
	}

	if !plan.Icon.IsNull() {
		values["icon"] = plan.Icon.ValueString()
	} else {
		values["icon"] = nil
	}

	if !plan.ParentID.IsNull() {
		values["parent_id"] = plan.ParentID.ValueInt64()
	} else {
		values["parent_id"] = nil
	}

	// Call API
	result, err := r.client.Update("ContactType", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating contact type",
			"Could not update contact type ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Update state
	plan.ID = state.ID
	r.mapResponseToModel(result, &plan)

	tflog.Debug(ctx, "Updated contact type", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ContactTypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ContactTypeResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting contact type", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	err := r.client.Delete("ContactType", state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting contact type",
			"Could not delete contact type ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "Deleted contact type", map[string]any{
		"id": state.ID.ValueInt64(),
	})
}

func (r *ContactTypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (r *ContactTypeResource) mapResponseToModel(result map[string]any, model *ContactTypeResourceModel) {
	if id, ok := GetInt64(result, "id"); ok {
		model.ID = types.Int64Value(id)
	}

	if name, ok := GetString(result, "name"); ok {
		model.Name = types.StringValue(name)
	}

	if label, ok := GetString(result, "label"); ok {
		model.Label = types.StringValue(label)
	}

	if description, ok := GetString(result, "description"); ok && description != "" {
		model.Description = types.StringValue(description)
	} else {
		model.Description = types.StringNull()
	}

	if imageURL, ok := GetString(result, "image_URL"); ok && imageURL != "" {
		model.ImageURL = types.StringValue(imageURL)
	} else {
		model.ImageURL = types.StringNull()
	}

	if icon, ok := GetString(result, "icon"); ok && icon != "" {
		model.Icon = types.StringValue(icon)
	} else {
		model.Icon = types.StringNull()
	}

	if parentID, ok := GetInt64(result, "parent_id"); ok {
		model.ParentID = types.Int64Value(parentID)
	} else {
		model.ParentID = types.Int64Null()
	}

	if isActive, ok := GetBool(result, "is_active"); ok {
		model.IsActive = types.BoolValue(isActive)
	}

	if isReserved, ok := GetBool(result, "is_reserved"); ok {
		model.IsReserved = types.BoolValue(isReserved)
	}
}
