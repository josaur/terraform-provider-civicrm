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
	_ resource.Resource                = &GroupResource{}
	_ resource.ResourceWithConfigure   = &GroupResource{}
	_ resource.ResourceWithImportState = &GroupResource{}
)

type GroupResource struct {
	client *Client
}

type GroupResourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Title       types.String `tfsdk:"title"`
	Description types.String `tfsdk:"description"`
	IsActive    types.Bool   `tfsdk:"is_active"`
	Visibility  types.String `tfsdk:"visibility"`
	GroupType   types.List   `tfsdk:"group_type"`
}

func NewGroupResource() resource.Resource {
	return &GroupResource{}
}

func (r *GroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

func (r *GroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a CiviCRM Group. Groups are collections of contacts that can be used for ACL role assignments.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier of the group.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The machine name of the group (must be unique).",
				Required:    true,
			},
			"title": schema.StringAttribute{
				Description: "The display title of the group.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the group.",
				Optional:    true,
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether the group is active. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"visibility": schema.StringAttribute{
				Description: "The visibility of the group. Options: 'User and User Admin Only', 'Public Pages'. Default: 'User and User Admin Only'.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("User and User Admin Only"),
			},
			"group_type": schema.ListAttribute{
				Description: "The types of the group (e.g., 'Mailing List', 'Access Control').",
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *GroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan GroupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating group", map[string]any{
		"name":  plan.Name.ValueString(),
		"title": plan.Title.ValueString(),
	})

	// Build values for API call
	values := map[string]any{
		"name":       plan.Name.ValueString(),
		"title":      plan.Title.ValueString(),
		"is_active":  plan.IsActive.ValueBool(),
		"visibility": plan.Visibility.ValueString(),
	}

	if !plan.Description.IsNull() {
		values["description"] = plan.Description.ValueString()
	}

	if !plan.GroupType.IsNull() {
		var groupTypes []string
		diags = plan.GroupType.ElementsAs(ctx, &groupTypes, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		values["group_type"] = groupTypes
	}

	// Call API
	result, err := r.client.Create("Group", values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating group",
			"Could not create group, unexpected error: "+err.Error(),
		)
		return
	}

	// Update state with response
	if id, ok := GetInt64(result, "id"); ok {
		plan.ID = types.Int64Value(id)
	}

	if name, ok := GetString(result, "name"); ok {
		plan.Name = types.StringValue(name)
	}

	if title, ok := GetString(result, "title"); ok {
		plan.Title = types.StringValue(title)
	}

	if desc, ok := GetString(result, "description"); ok {
		plan.Description = types.StringValue(desc)
	} else {
		plan.Description = types.StringNull()
	}

	if active, ok := GetBool(result, "is_active"); ok {
		plan.IsActive = types.BoolValue(active)
	}

	if visibility, ok := GetString(result, "visibility"); ok {
		plan.Visibility = types.StringValue(visibility)
	}

	tflog.Debug(ctx, "Created group", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *GroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state GroupResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading group", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	result, err := r.client.GetByID("Group", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading group",
			"Could not read group ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Update state
	if name, ok := GetString(result, "name"); ok {
		state.Name = types.StringValue(name)
	}

	if title, ok := GetString(result, "title"); ok {
		state.Title = types.StringValue(title)
	}

	if desc, ok := GetString(result, "description"); ok && desc != "" {
		state.Description = types.StringValue(desc)
	} else {
		state.Description = types.StringNull()
	}

	if active, ok := GetBool(result, "is_active"); ok {
		state.IsActive = types.BoolValue(active)
	}

	if visibility, ok := GetString(result, "visibility"); ok {
		state.Visibility = types.StringValue(visibility)
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *GroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan GroupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state GroupResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating group", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	// Build values for API call
	values := map[string]any{
		"name":       plan.Name.ValueString(),
		"title":      plan.Title.ValueString(),
		"is_active":  plan.IsActive.ValueBool(),
		"visibility": plan.Visibility.ValueString(),
	}

	if !plan.Description.IsNull() {
		values["description"] = plan.Description.ValueString()
	} else {
		values["description"] = nil
	}

	if !plan.GroupType.IsNull() {
		var groupTypes []string
		diags = plan.GroupType.ElementsAs(ctx, &groupTypes, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		values["group_type"] = groupTypes
	}

	// Call API
	result, err := r.client.Update("Group", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating group",
			"Could not update group ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Update state
	plan.ID = state.ID

	if name, ok := GetString(result, "name"); ok {
		plan.Name = types.StringValue(name)
	}

	if title, ok := GetString(result, "title"); ok {
		plan.Title = types.StringValue(title)
	}

	if desc, ok := GetString(result, "description"); ok && desc != "" {
		plan.Description = types.StringValue(desc)
	} else {
		plan.Description = types.StringNull()
	}

	if active, ok := GetBool(result, "is_active"); ok {
		plan.IsActive = types.BoolValue(active)
	}

	if visibility, ok := GetString(result, "visibility"); ok {
		plan.Visibility = types.StringValue(visibility)
	}

	tflog.Debug(ctx, "Updated group", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *GroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state GroupResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting group", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	err := r.client.Delete("Group", state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting group",
			"Could not delete group ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "Deleted group", map[string]any{
		"id": state.ID.ValueInt64(),
	})
}

func (r *GroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
