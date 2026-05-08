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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &ActivityTypeResource{}
	_ resource.ResourceWithConfigure   = &ActivityTypeResource{}
	_ resource.ResourceWithImportState = &ActivityTypeResource{}
)

// ActivityTypeResource manages CiviCRM activity types.
// Activity types are OptionValues in the "activity_type" OptionGroup.
type ActivityTypeResource struct {
	client *Client
}

type ActivityTypeResourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Label       types.String `tfsdk:"label"`
	Description types.String `tfsdk:"description"`
	IsActive    types.Bool   `tfsdk:"is_active"`
	IsReserved  types.Bool   `tfsdk:"is_reserved"`
	Weight      types.Int64  `tfsdk:"weight"`
	Value       types.String `tfsdk:"value"`
	Color       types.String `tfsdk:"color"`
	Icon        types.String `tfsdk:"icon"`
}

func NewActivityTypeResource() resource.Resource {
	return &ActivityTypeResource{}
}

func (r *ActivityTypeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_activity_type"
}

func (r *ActivityTypeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a CiviCRM Activity Type. Activity types are OptionValues in the 'activity_type' option group and define the kinds of activities that can be recorded against contacts.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Unique ID of the OptionValue.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Machine name of the activity type (e.g. 'phone_call', 'housing_assessment'). Must be unique within the activity_type group.",
				Required:    true,
			},
			"label": schema.StringAttribute{
				Description: "Display label shown in the UI (e.g. 'Phone Call', 'Housing Assessment').",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Optional description of the activity type.",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether the activity type is active. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"is_reserved": schema.BoolAttribute{
				Description: "Whether the activity type is reserved (protected from deletion by CiviCRM). Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"weight": schema.Int64Attribute{
				Description: "Sort weight. Controls display order in dropdowns.",
				Optional:    true,
				Computed:    true,
			},
			"value": schema.StringAttribute{
				Description: "Internal numeric value used by CiviCRM to identify this activity type. Auto-generated if not set.",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"color": schema.StringAttribute{
				Description: "Hex color code for the activity type (e.g. '#FF0000'). Used for calendar and timeline display.",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"icon": schema.StringAttribute{
				Description: "CSS icon class for the activity type (e.g. 'crm-i fa-phone').",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *ActivityTypeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *Client, got: %T.", req.ProviderData),
		)
		return
	}
	r.client = client
}

func (r *ActivityTypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ActivityTypeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ActivityType", map[string]any{"name": plan.Name.ValueString()})

	optionGroupID, err := r.client.GetOptionGroupID("activity_type")
	if err != nil {
		resp.Diagnostics.AddError("Error looking up option group", "Could not find activity_type option group: "+err.Error())
		return
	}

	values := map[string]any{
		"option_group_id": optionGroupID,
		"name":            plan.Name.ValueString(),
		"label":           plan.Label.ValueString(),
		"is_active":       plan.IsActive.ValueBool(),
		"is_reserved":     plan.IsReserved.ValueBool(),
	}
	if !plan.Description.IsNull() && plan.Description.ValueString() != "" {
		values["description"] = plan.Description.ValueString()
	}
	if !plan.Weight.IsNull() {
		values["weight"] = plan.Weight.ValueInt64()
	}
	if !plan.Value.IsNull() && plan.Value.ValueString() != "" {
		values["value"] = plan.Value.ValueString()
	}
	if !plan.Color.IsNull() && plan.Color.ValueString() != "" {
		values["color"] = plan.Color.ValueString()
	}
	if !plan.Icon.IsNull() && plan.Icon.ValueString() != "" {
		values["icon"] = plan.Icon.ValueString()
	}

	result, err := r.client.Create("OptionValue", values)
	if err != nil {
		resp.Diagnostics.AddError("Error creating ActivityType", err.Error())
		return
	}

	r.mapResultToState(result, &plan)
	tflog.Debug(ctx, "Created ActivityType", map[string]any{"id": plan.ID.ValueInt64()})
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *ActivityTypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ActivityTypeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.GetByID("OptionValue", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading ActivityType",
			"Could not read ActivityType ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	r.mapResultToState(result, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *ActivityTypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ActivityTypeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var state ActivityTypeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	values := map[string]any{
		"name":        plan.Name.ValueString(),
		"label":       plan.Label.ValueString(),
		"is_active":   plan.IsActive.ValueBool(),
		"is_reserved": plan.IsReserved.ValueBool(),
	}
	if !plan.Description.IsNull() && plan.Description.ValueString() != "" {
		values["description"] = plan.Description.ValueString()
	} else {
		values["description"] = nil
	}
	if !plan.Weight.IsNull() {
		values["weight"] = plan.Weight.ValueInt64()
	}
	if !plan.Value.IsNull() && plan.Value.ValueString() != "" {
		values["value"] = plan.Value.ValueString()
	}
	if !plan.Color.IsNull() && plan.Color.ValueString() != "" {
		values["color"] = plan.Color.ValueString()
	} else {
		values["color"] = nil
	}
	if !plan.Icon.IsNull() && plan.Icon.ValueString() != "" {
		values["icon"] = plan.Icon.ValueString()
	} else {
		values["icon"] = nil
	}

	result, err := r.client.Update("OptionValue", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating ActivityType",
			"Could not update ActivityType ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	plan.ID = state.ID
	r.mapResultToState(result, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *ActivityTypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ActivityTypeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.Delete("OptionValue", state.ID.ValueInt64()); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting ActivityType",
			"Could not delete ActivityType ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
	}
}

func (r *ActivityTypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", "Could not parse import ID as integer: "+err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (r *ActivityTypeResource) mapResultToState(result map[string]any, model *ActivityTypeResourceModel) {
	if id, ok := GetInt64(result, "id"); ok {
		model.ID = types.Int64Value(id)
	}
	if name, ok := GetString(result, "name"); ok {
		model.Name = types.StringValue(name)
	}
	if label, ok := GetString(result, "label"); ok {
		model.Label = types.StringValue(label)
	}
	if desc, ok := GetString(result, "description"); ok && desc != "" {
		model.Description = types.StringValue(desc)
	} else {
		model.Description = types.StringNull()
	}
	if active, ok := GetBool(result, "is_active"); ok {
		model.IsActive = types.BoolValue(active)
	}
	if reserved, ok := GetBool(result, "is_reserved"); ok {
		model.IsReserved = types.BoolValue(reserved)
	}
	if weight, ok := GetInt64(result, "weight"); ok {
		model.Weight = types.Int64Value(weight)
	}
	if value, ok := GetString(result, "value"); ok && value != "" {
		model.Value = types.StringValue(value)
	} else {
		model.Value = types.StringNull()
	}
	if color, ok := GetString(result, "color"); ok && color != "" {
		model.Color = types.StringValue(color)
	} else {
		model.Color = types.StringNull()
	}
	if icon, ok := GetString(result, "icon"); ok && icon != "" {
		model.Icon = types.StringValue(icon)
	} else {
		model.Icon = types.StringNull()
	}
}
