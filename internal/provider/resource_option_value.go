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
	_ resource.Resource                = &OptionValueResource{}
	_ resource.ResourceWithConfigure   = &OptionValueResource{}
	_ resource.ResourceWithImportState = &OptionValueResource{}
)

// OptionValueResource manages a single OptionValue inside an OptionGroup.
// OptionValues back the choices of Radio/Select/Multi-Select custom fields
// (via option_group_id on civicrm_custom_field) as well as many built-in
// CiviCRM enumerations. For the case_status option group there is a
// dedicated civicrm_case_status resource with a fixed group.
type OptionValueResource struct {
	client *Client
}

type OptionValueResourceModel struct {
	ID            types.Int64  `tfsdk:"id"`
	OptionGroupID types.Int64  `tfsdk:"option_group_id"`
	Label         types.String `tfsdk:"label"`
	Value         types.String `tfsdk:"value"`
	Name          types.String `tfsdk:"name"`
	Description   types.String `tfsdk:"description"`
	Weight        types.Int64  `tfsdk:"weight"`
	IsActive      types.Bool   `tfsdk:"is_active"`
	IsReserved    types.Bool   `tfsdk:"is_reserved"`
	IsDefault     types.Bool   `tfsdk:"is_default"`
	Grouping      types.String `tfsdk:"grouping"`
	Color         types.String `tfsdk:"color"`
	Icon          types.String `tfsdk:"icon"`
}

func NewOptionValueResource() resource.Resource {
	return &OptionValueResource{}
}

func (r *OptionValueResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_option_value"
}

func (r *OptionValueResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a CiviCRM OptionValue inside an OptionGroup. Used to declare the choices of Radio/Select custom fields and enumerations.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Unique ID of the option value.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"option_group_id": schema.Int64Attribute{
				Description: "ID of the parent OptionGroup.",
				Required:    true,
			},
			"label": schema.StringAttribute{
				Description: "Display label shown to users.",
				Required:    true,
			},
			"value": schema.StringAttribute{
				Description: "Stored value (what the custom field records). Must be unique within the group.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Machine name. Defaults to `value` if not set.",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Optional description.",
				Optional:    true,
			},
			"weight": schema.Int64Attribute{
				Description: "Sort weight. Controls display order.",
				Optional:    true,
				Computed:    true,
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether the option value is active. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"is_reserved": schema.BoolAttribute{
				Description: "Whether the option value is reserved (protected from deletion). Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"is_default": schema.BoolAttribute{
				Description: "Whether this is the default option in the group. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"grouping": schema.StringAttribute{
				Description: "Optional category bucket used by some CiviCRM contexts (e.g. case_status uses \"Opened\"/\"Closed\").",
				Optional:    true,
			},
			"color": schema.StringAttribute{
				Description: "Optional color in hex format (e.g. `#ff0000`).",
				Optional:    true,
			},
			"icon": schema.StringAttribute{
				Description: "Optional icon (CSS class name).",
				Optional:    true,
			},
		},
	}
}

func (r *OptionValueResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OptionValueResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan OptionValueResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating OptionValue", map[string]any{
		"option_group_id": plan.OptionGroupID.ValueInt64(),
		"value":           plan.Value.ValueString(),
	})

	values := map[string]any{
		"option_group_id": plan.OptionGroupID.ValueInt64(),
		"label":           plan.Label.ValueString(),
		"value":           plan.Value.ValueString(),
		"is_active":       plan.IsActive.ValueBool(),
		"is_reserved":     plan.IsReserved.ValueBool(),
		"is_default":      plan.IsDefault.ValueBool(),
	}
	if !plan.Name.IsNull() && !plan.Name.IsUnknown() && plan.Name.ValueString() != "" {
		values["name"] = plan.Name.ValueString()
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		values["description"] = plan.Description.ValueString()
	}
	if !plan.Weight.IsNull() && !plan.Weight.IsUnknown() {
		values["weight"] = plan.Weight.ValueInt64()
	}
	if !plan.Grouping.IsNull() && !plan.Grouping.IsUnknown() && plan.Grouping.ValueString() != "" {
		values["grouping"] = plan.Grouping.ValueString()
	}
	if !plan.Color.IsNull() && !plan.Color.IsUnknown() && plan.Color.ValueString() != "" {
		values["color"] = plan.Color.ValueString()
	}
	if !plan.Icon.IsNull() && !plan.Icon.IsUnknown() && plan.Icon.ValueString() != "" {
		values["icon"] = plan.Icon.ValueString()
	}

	result, err := r.client.Create("OptionValue", values)
	if err != nil {
		resp.Diagnostics.AddError("Error creating OptionValue", err.Error())
		return
	}

	if createdID, ok := GetInt64(result, "id"); ok {
		if fullResult, err2 := r.client.GetByID("OptionValue", createdID, nil); err2 == nil {
			result = fullResult
		}
	}
	r.mapResultToState(result, &plan)
	tflog.Debug(ctx, "Created OptionValue", map[string]any{"id": plan.ID.ValueInt64()})
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *OptionValueResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OptionValueResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.GetByID("OptionValue", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading OptionValue",
			"Could not read OptionValue ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	r.mapResultToState(result, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *OptionValueResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan OptionValueResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var state OptionValueResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	values := map[string]any{
		"option_group_id": plan.OptionGroupID.ValueInt64(),
		"label":           plan.Label.ValueString(),
		"value":           plan.Value.ValueString(),
		"is_active":       plan.IsActive.ValueBool(),
		"is_reserved":     plan.IsReserved.ValueBool(),
		"is_default":      plan.IsDefault.ValueBool(),
	}
	if !plan.Name.IsNull() && !plan.Name.IsUnknown() && plan.Name.ValueString() != "" {
		values["name"] = plan.Name.ValueString()
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		values["description"] = plan.Description.ValueString()
	} else {
		values["description"] = nil
	}
	if !plan.Weight.IsNull() && !plan.Weight.IsUnknown() {
		values["weight"] = plan.Weight.ValueInt64()
	}
	if !plan.Grouping.IsNull() && !plan.Grouping.IsUnknown() && plan.Grouping.ValueString() != "" {
		values["grouping"] = plan.Grouping.ValueString()
	} else {
		values["grouping"] = nil
	}
	if !plan.Color.IsNull() && !plan.Color.IsUnknown() && plan.Color.ValueString() != "" {
		values["color"] = plan.Color.ValueString()
	} else {
		values["color"] = nil
	}
	if !plan.Icon.IsNull() && !plan.Icon.IsUnknown() && plan.Icon.ValueString() != "" {
		values["icon"] = plan.Icon.ValueString()
	} else {
		values["icon"] = nil
	}

	_, err := r.client.Update("OptionValue", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating OptionValue",
			"Could not update OptionValue ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	plan.ID = state.ID

	result, err := r.client.GetByID("OptionValue", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading OptionValue after update",
			"Could not re-read OptionValue ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}
	r.mapResultToState(result, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *OptionValueResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OptionValueResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.Delete("OptionValue", state.ID.ValueInt64()); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting OptionValue",
			"Could not delete OptionValue ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
	}
}

func (r *OptionValueResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", "Could not parse import ID as integer: "+err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (r *OptionValueResource) mapResultToState(result map[string]any, model *OptionValueResourceModel) {
	if id, ok := GetInt64(result, "id"); ok {
		model.ID = types.Int64Value(id)
	}
	if ogid, ok := GetInt64(result, "option_group_id"); ok {
		model.OptionGroupID = types.Int64Value(ogid)
	}
	if label, ok := GetString(result, "label"); ok {
		model.Label = types.StringValue(label)
	}
	if value, ok := GetString(result, "value"); ok {
		model.Value = types.StringValue(value)
	}
	if name, ok := GetString(result, "name"); ok {
		model.Name = types.StringValue(name)
	}
	if desc, ok := GetString(result, "description"); ok && desc != "" {
		model.Description = types.StringValue(desc)
	} else {
		model.Description = types.StringNull()
	}
	if weight, ok := GetInt64(result, "weight"); ok {
		model.Weight = types.Int64Value(weight)
	}
	if active, ok := GetBool(result, "is_active"); ok {
		model.IsActive = types.BoolValue(active)
	}
	if reserved, ok := GetBool(result, "is_reserved"); ok {
		model.IsReserved = types.BoolValue(reserved)
	}
	if def, ok := GetBool(result, "is_default"); ok {
		model.IsDefault = types.BoolValue(def)
	}
	if grouping, ok := GetString(result, "grouping"); ok && grouping != "" {
		model.Grouping = types.StringValue(grouping)
	} else {
		model.Grouping = types.StringNull()
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
