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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &OptionGroupResource{}
	_ resource.ResourceWithConfigure   = &OptionGroupResource{}
	_ resource.ResourceWithImportState = &OptionGroupResource{}
)

// OptionGroupResource manages CiviCRM OptionGroups. An OptionGroup is a named
// container for OptionValues used to back Radio/Select/Multi-Select custom
// fields as well as many built-in CiviCRM enumerations (e.g. gender,
// activity_type). Pair with civicrm_option_value.
type OptionGroupResource struct {
	client *Client
}

type OptionGroupResourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Title       types.String `tfsdk:"title"`
	Description types.String `tfsdk:"description"`
	DataType    types.String `tfsdk:"data_type"`
	IsActive    types.Bool   `tfsdk:"is_active"`
	IsReserved  types.Bool   `tfsdk:"is_reserved"`
}

func NewOptionGroupResource() resource.Resource {
	return &OptionGroupResource{}
}

func (r *OptionGroupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_option_group"
}

func (r *OptionGroupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a CiviCRM OptionGroup. An OptionGroup is a named container for OptionValues used to back Radio/Select/Multi-Select custom fields and built-in enumerations. Pair with civicrm_option_value to declare the members.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Unique ID of the option group.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Machine name (must be unique, no spaces).",
				Required:    true,
			},
			"title": schema.StringAttribute{
				Description: "Display title shown in the admin UI.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Optional description.",
				Optional:    true,
			},
			"data_type": schema.StringAttribute{
				Description: "Data type of the values (e.g. \"String\", \"Integer\"). Default: \"String\".",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("String"),
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether the option group is active. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"is_reserved": schema.BoolAttribute{
				Description: "Whether the option group is reserved (protected from deletion by CiviCRM). Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
		},
	}
}

func (r *OptionGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OptionGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan OptionGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating OptionGroup", map[string]any{"name": plan.Name.ValueString()})

	values := map[string]any{
		"name":        plan.Name.ValueString(),
		"title":       plan.Title.ValueString(),
		"data_type":   plan.DataType.ValueString(),
		"is_active":   plan.IsActive.ValueBool(),
		"is_reserved": plan.IsReserved.ValueBool(),
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		values["description"] = plan.Description.ValueString()
	}

	result, err := r.client.Create("OptionGroup", values)
	if err != nil {
		resp.Diagnostics.AddError("Error creating OptionGroup", err.Error())
		return
	}

	if createdID, ok := GetInt64(result, "id"); ok {
		if fullResult, err2 := r.client.GetByID("OptionGroup", createdID, nil); err2 == nil {
			result = fullResult
		}
	}
	r.mapResultToState(result, &plan)
	tflog.Debug(ctx, "Created OptionGroup", map[string]any{"id": plan.ID.ValueInt64()})
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *OptionGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OptionGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.GetByID("OptionGroup", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading OptionGroup",
			"Could not read OptionGroup ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	r.mapResultToState(result, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *OptionGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan OptionGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var state OptionGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	values := map[string]any{
		"name":        plan.Name.ValueString(),
		"title":       plan.Title.ValueString(),
		"data_type":   plan.DataType.ValueString(),
		"is_active":   plan.IsActive.ValueBool(),
		"is_reserved": plan.IsReserved.ValueBool(),
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		values["description"] = plan.Description.ValueString()
	} else {
		values["description"] = nil
	}

	_, err := r.client.Update("OptionGroup", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating OptionGroup",
			"Could not update OptionGroup ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	plan.ID = state.ID

	result, err := r.client.GetByID("OptionGroup", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading OptionGroup after update",
			"Could not re-read OptionGroup ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}
	r.mapResultToState(result, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *OptionGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OptionGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.Delete("OptionGroup", state.ID.ValueInt64()); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting OptionGroup",
			"Could not delete OptionGroup ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
	}
}

func (r *OptionGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", "Could not parse import ID as integer: "+err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (r *OptionGroupResource) mapResultToState(result map[string]any, model *OptionGroupResourceModel) {
	if id, ok := GetInt64(result, "id"); ok {
		model.ID = types.Int64Value(id)
	}
	if name, ok := GetString(result, "name"); ok {
		model.Name = types.StringValue(name)
	}
	if title, ok := GetString(result, "title"); ok {
		model.Title = types.StringValue(title)
	}
	if desc, ok := GetString(result, "description"); ok && desc != "" {
		model.Description = types.StringValue(desc)
	} else {
		model.Description = types.StringNull()
	}
	if dt, ok := GetString(result, "data_type"); ok && dt != "" {
		model.DataType = types.StringValue(dt)
	} else {
		model.DataType = types.StringValue("String")
	}
	if active, ok := GetBool(result, "is_active"); ok {
		model.IsActive = types.BoolValue(active)
	}
	if reserved, ok := GetBool(result, "is_reserved"); ok {
		model.IsReserved = types.BoolValue(reserved)
	}
}
