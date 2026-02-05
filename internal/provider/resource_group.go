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

// Group type mappings between human-readable names and CiviCRM API values
var groupTypeNameToID = map[string]string{
	"Access Control": "1",
	"Mailing List":   "2",
}

var groupTypeIDToName = map[string]string{
	"1": "Access Control",
	"2": "Mailing List",
}

// convertGroupTypesToIDs converts human-readable group type names to API IDs
func convertGroupTypesToIDs(names []string) []string {
	ids := make([]string, 0, len(names))
	for _, name := range names {
		if id, ok := groupTypeNameToID[name]; ok {
			ids = append(ids, id)
		}
	}
	return ids
}

// convertGroupTypeIDsToNames converts API IDs to human-readable group type names
func convertGroupTypeIDsToNames(ids []string) []string {
	names := make([]string, 0, len(ids))
	for _, id := range ids {
		if name, ok := groupTypeIDToName[id]; ok {
			names = append(names, name)
		}
	}
	return names
}

type GroupResource struct {
	client *Client
}

type GroupResourceModel struct {
	ID                  types.Int64  `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	Title               types.String `tfsdk:"title"`
	Description         types.String `tfsdk:"description"`
	IsActive            types.Bool   `tfsdk:"is_active"`
	Visibility          types.String `tfsdk:"visibility"`
	GroupType           types.List   `tfsdk:"group_type"`
	IsHidden            types.Bool   `tfsdk:"is_hidden"`
	IsReserved          types.Bool   `tfsdk:"is_reserved"`
	FrontendTitle       types.String `tfsdk:"frontend_title"`
	FrontendDescription types.String `tfsdk:"frontend_description"`
	Parents             types.List   `tfsdk:"parents"`
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
				Description: "The types of the group. Valid values: 'Access Control', 'Mailing List'.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"is_hidden": schema.BoolAttribute{
				Description: "Whether the group is hidden from the user interface. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"is_reserved": schema.BoolAttribute{
				Description: "Whether the group is reserved (system group). Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"frontend_title": schema.StringAttribute{
				Description: "The public title of the group shown on frontend pages.",
				Optional:    true,
			},
			"frontend_description": schema.StringAttribute{
				Description: "The public description of the group shown on frontend pages.",
				Optional:    true,
			},
			"parents": schema.ListAttribute{
				Description: "List of parent group IDs for nested groups.",
				Optional:    true,
				ElementType: types.Int64Type,
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
		"name":        plan.Name.ValueString(),
		"title":       plan.Title.ValueString(),
		"is_active":   plan.IsActive.ValueBool(),
		"visibility":  plan.Visibility.ValueString(),
		"is_hidden":   plan.IsHidden.ValueBool(),
		"is_reserved": plan.IsReserved.ValueBool(),
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
		// Convert human-readable names to API IDs
		values["group_type"] = convertGroupTypesToIDs(groupTypes)
	}

	if !plan.FrontendTitle.IsNull() {
		values["frontend_title"] = plan.FrontendTitle.ValueString()
	}

	if !plan.FrontendDescription.IsNull() {
		values["frontend_description"] = plan.FrontendDescription.ValueString()
	}

	if !plan.Parents.IsNull() {
		var parents []int64
		diags = plan.Parents.ElementsAs(ctx, &parents, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		values["parents"] = parents
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

	// Handle group_type from API response
	if groupTypeRaw, ok := result["group_type"]; ok && groupTypeRaw != nil {
		if groupTypeSlice, ok := groupTypeRaw.([]any); ok {
			ids := make([]string, 0, len(groupTypeSlice))
			for _, v := range groupTypeSlice {
				if s, ok := v.(string); ok {
					ids = append(ids, s)
				}
			}
			names := convertGroupTypeIDsToNames(ids)
			groupTypeList, diags := types.ListValueFrom(ctx, types.StringType, names)
			resp.Diagnostics.Append(diags...)
			if !resp.Diagnostics.HasError() {
				plan.GroupType = groupTypeList
			}
		}
	}

	if hidden, ok := GetBool(result, "is_hidden"); ok {
		plan.IsHidden = types.BoolValue(hidden)
	}

	if reserved, ok := GetBool(result, "is_reserved"); ok {
		plan.IsReserved = types.BoolValue(reserved)
	}

	if frontendTitle, ok := GetString(result, "frontend_title"); ok && frontendTitle != "" {
		plan.FrontendTitle = types.StringValue(frontendTitle)
	} else {
		plan.FrontendTitle = types.StringNull()
	}

	if frontendDesc, ok := GetString(result, "frontend_description"); ok && frontendDesc != "" {
		plan.FrontendDescription = types.StringValue(frontendDesc)
	} else {
		plan.FrontendDescription = types.StringNull()
	}

	// Handle parents from API response
	if parentsRaw, ok := result["parents"]; ok && parentsRaw != nil {
		if parentsSlice, ok := parentsRaw.([]any); ok {
			parentIDs := make([]int64, 0, len(parentsSlice))
			for _, v := range parentsSlice {
				if id, ok := v.(float64); ok {
					parentIDs = append(parentIDs, int64(id))
				} else if id, ok := v.(int64); ok {
					parentIDs = append(parentIDs, id)
				}
			}
			if len(parentIDs) > 0 {
				parentsList, diags := types.ListValueFrom(ctx, types.Int64Type, parentIDs)
				resp.Diagnostics.Append(diags...)
				if !resp.Diagnostics.HasError() {
					plan.Parents = parentsList
				}
			}
		}
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

	// Handle group_type from API response
	if groupTypeRaw, ok := result["group_type"]; ok && groupTypeRaw != nil {
		if groupTypeSlice, ok := groupTypeRaw.([]any); ok {
			ids := make([]string, 0, len(groupTypeSlice))
			for _, v := range groupTypeSlice {
				if s, ok := v.(string); ok {
					ids = append(ids, s)
				}
			}
			names := convertGroupTypeIDsToNames(ids)
			groupTypeList, diags := types.ListValueFrom(ctx, types.StringType, names)
			resp.Diagnostics.Append(diags...)
			if !resp.Diagnostics.HasError() {
				state.GroupType = groupTypeList
			}
		}
	}

	if hidden, ok := GetBool(result, "is_hidden"); ok {
		state.IsHidden = types.BoolValue(hidden)
	}

	if reserved, ok := GetBool(result, "is_reserved"); ok {
		state.IsReserved = types.BoolValue(reserved)
	}

	if frontendTitle, ok := GetString(result, "frontend_title"); ok && frontendTitle != "" {
		state.FrontendTitle = types.StringValue(frontendTitle)
	} else {
		state.FrontendTitle = types.StringNull()
	}

	if frontendDesc, ok := GetString(result, "frontend_description"); ok && frontendDesc != "" {
		state.FrontendDescription = types.StringValue(frontendDesc)
	} else {
		state.FrontendDescription = types.StringNull()
	}

	// Handle parents from API response
	if parentsRaw, ok := result["parents"]; ok && parentsRaw != nil {
		if parentsSlice, ok := parentsRaw.([]any); ok {
			parentIDs := make([]int64, 0, len(parentsSlice))
			for _, v := range parentsSlice {
				if id, ok := v.(float64); ok {
					parentIDs = append(parentIDs, int64(id))
				} else if id, ok := v.(int64); ok {
					parentIDs = append(parentIDs, id)
				}
			}
			if len(parentIDs) > 0 {
				parentsList, diags := types.ListValueFrom(ctx, types.Int64Type, parentIDs)
				resp.Diagnostics.Append(diags...)
				if !resp.Diagnostics.HasError() {
					state.Parents = parentsList
				}
			}
		}
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
		"name":        plan.Name.ValueString(),
		"title":       plan.Title.ValueString(),
		"is_active":   plan.IsActive.ValueBool(),
		"visibility":  plan.Visibility.ValueString(),
		"is_hidden":   plan.IsHidden.ValueBool(),
		"is_reserved": plan.IsReserved.ValueBool(),
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
		// Convert human-readable names to API IDs
		values["group_type"] = convertGroupTypesToIDs(groupTypes)
	}

	if !plan.FrontendTitle.IsNull() {
		values["frontend_title"] = plan.FrontendTitle.ValueString()
	} else {
		values["frontend_title"] = nil
	}

	if !plan.FrontendDescription.IsNull() {
		values["frontend_description"] = plan.FrontendDescription.ValueString()
	} else {
		values["frontend_description"] = nil
	}

	if !plan.Parents.IsNull() {
		var parents []int64
		diags = plan.Parents.ElementsAs(ctx, &parents, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		values["parents"] = parents
	} else {
		values["parents"] = nil
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

	// Handle group_type from API response
	if groupTypeRaw, ok := result["group_type"]; ok && groupTypeRaw != nil {
		if groupTypeSlice, ok := groupTypeRaw.([]any); ok {
			ids := make([]string, 0, len(groupTypeSlice))
			for _, v := range groupTypeSlice {
				if s, ok := v.(string); ok {
					ids = append(ids, s)
				}
			}
			names := convertGroupTypeIDsToNames(ids)
			groupTypeList, diags := types.ListValueFrom(ctx, types.StringType, names)
			resp.Diagnostics.Append(diags...)
			if !resp.Diagnostics.HasError() {
				plan.GroupType = groupTypeList
			}
		}
	}

	if hidden, ok := GetBool(result, "is_hidden"); ok {
		plan.IsHidden = types.BoolValue(hidden)
	}

	if reserved, ok := GetBool(result, "is_reserved"); ok {
		plan.IsReserved = types.BoolValue(reserved)
	}

	if frontendTitle, ok := GetString(result, "frontend_title"); ok && frontendTitle != "" {
		plan.FrontendTitle = types.StringValue(frontendTitle)
	} else {
		plan.FrontendTitle = types.StringNull()
	}

	if frontendDesc, ok := GetString(result, "frontend_description"); ok && frontendDesc != "" {
		plan.FrontendDescription = types.StringValue(frontendDesc)
	} else {
		plan.FrontendDescription = types.StringNull()
	}

	// Handle parents from API response
	if parentsRaw, ok := result["parents"]; ok && parentsRaw != nil {
		if parentsSlice, ok := parentsRaw.([]any); ok {
			parentIDs := make([]int64, 0, len(parentsSlice))
			for _, v := range parentsSlice {
				if id, ok := v.(float64); ok {
					parentIDs = append(parentIDs, int64(id))
				} else if id, ok := v.(int64); ok {
					parentIDs = append(parentIDs, id)
				}
			}
			if len(parentIDs) > 0 {
				parentsList, diags := types.ListValueFrom(ctx, types.Int64Type, parentIDs)
				resp.Diagnostics.Append(diags...)
				if !resp.Diagnostics.HasError() {
					plan.Parents = parentsList
				}
			}
		}
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
