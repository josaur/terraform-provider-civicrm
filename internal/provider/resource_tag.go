package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	_ resource.Resource                = &TagResource{}
	_ resource.ResourceWithConfigure   = &TagResource{}
	_ resource.ResourceWithImportState = &TagResource{}
)

// TagResource manages tags in CiviCRM.
type TagResource struct {
	client *Client
}

type TagResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Label        types.String `tfsdk:"label"`
	Description  types.String `tfsdk:"description"`
	ParentID     types.Int64  `tfsdk:"parent_id"`
	IsSelectable types.Bool   `tfsdk:"is_selectable"`
	IsReserved   types.Bool   `tfsdk:"is_reserved"`
	IsTagset     types.Bool   `tfsdk:"is_tagset"`
	UsedFor      types.List   `tfsdk:"used_for"`
	Color        types.String `tfsdk:"color"`
}

func NewTagResource() resource.Resource {
	return &TagResource{}
}

func (r *TagResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tag"
}

func (r *TagResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages CiviCRM Tags for categorizing contacts and other entities.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier of the tag.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The machine name of the tag (must be unique, no spaces).",
				Required:    true,
			},
			"label": schema.StringAttribute{
				Description: "The display label of the tag.",
				Optional:    true,
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the tag.",
				Optional:    true,
			},
			"parent_id": schema.Int64Attribute{
				Description: "The parent tag ID for hierarchical tags.",
				Optional:    true,
			},
			"is_selectable": schema.BoolAttribute{
				Description: "Whether this tag can be selected. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"is_reserved": schema.BoolAttribute{
				Description: "Whether this is a reserved system tag. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"is_tagset": schema.BoolAttribute{
				Description: "Whether this is a tagset (container for other tags). Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"used_for": schema.ListAttribute{
				Description: "Entity types this tag can be used for (e.g., 'civicrm_contact', 'civicrm_activity').",
				Optional:    true,
				ElementType: types.StringType,
			},
			"color": schema.StringAttribute{
				Description: "The color for the tag in hex format (e.g., '#ff0000').",
				Optional:    true,
			},
		},
	}
}

func (r *TagResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TagResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TagResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tag", map[string]any{
		"name": plan.Name.ValueString(),
	})

	// Build values for API call
	values := map[string]any{
		"name":          plan.Name.ValueString(),
		"is_selectable": plan.IsSelectable.ValueBool(),
		"is_reserved":   plan.IsReserved.ValueBool(),
		"is_tagset":     plan.IsTagset.ValueBool(),
	}

	if !plan.Label.IsNull() {
		values["label"] = plan.Label.ValueString()
	}

	if !plan.Description.IsNull() {
		values["description"] = plan.Description.ValueString()
	}

	if !plan.ParentID.IsNull() {
		values["parent_id"] = plan.ParentID.ValueInt64()
	}

	if !plan.UsedFor.IsNull() {
		var usedFor []string
		diags = plan.UsedFor.ElementsAs(ctx, &usedFor, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		values["used_for"] = usedFor
	}

	if !plan.Color.IsNull() {
		values["color"] = plan.Color.ValueString()
	}

	// Call API
	result, err := r.client.Create("Tag", values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating tag",
			"Could not create tag, unexpected error: "+err.Error(),
		)
		return
	}

	// Update state with response
	var d diag.Diagnostics
	r.mapResponseToModel(ctx, result, &plan, &d)
	resp.Diagnostics.Append(d...)

	tflog.Debug(ctx, "Created tag", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TagResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TagResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tag", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	result, err := r.client.GetByID("Tag", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading tag",
			"Could not read tag ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Update state
	var d diag.Diagnostics
	r.mapResponseToModel(ctx, result, &state, &d)
	resp.Diagnostics.Append(d...)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *TagResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan TagResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state TagResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating tag", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	// Build values for API call
	values := map[string]any{
		"name":          plan.Name.ValueString(),
		"is_selectable": plan.IsSelectable.ValueBool(),
		"is_reserved":   plan.IsReserved.ValueBool(),
		"is_tagset":     plan.IsTagset.ValueBool(),
	}

	if !plan.Label.IsNull() {
		values["label"] = plan.Label.ValueString()
	} else {
		values["label"] = nil
	}

	if !plan.Description.IsNull() {
		values["description"] = plan.Description.ValueString()
	} else {
		values["description"] = nil
	}

	if !plan.ParentID.IsNull() {
		values["parent_id"] = plan.ParentID.ValueInt64()
	} else {
		values["parent_id"] = nil
	}

	if !plan.UsedFor.IsNull() {
		var usedFor []string
		diags = plan.UsedFor.ElementsAs(ctx, &usedFor, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		values["used_for"] = usedFor
	} else {
		values["used_for"] = nil
	}

	if !plan.Color.IsNull() {
		values["color"] = plan.Color.ValueString()
	} else {
		values["color"] = nil
	}

	// Call API
	result, err := r.client.Update("Tag", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating tag",
			"Could not update tag ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Update state
	plan.ID = state.ID
	var d diag.Diagnostics
	r.mapResponseToModel(ctx, result, &plan, &d)
	resp.Diagnostics.Append(d...)

	tflog.Debug(ctx, "Updated tag", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TagResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TagResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tag", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	err := r.client.Delete("Tag", state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting tag",
			"Could not delete tag ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "Deleted tag", map[string]any{
		"id": state.ID.ValueInt64(),
	})
}

func (r *TagResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (r *TagResource) mapResponseToModel(ctx context.Context, result map[string]any, model *TagResourceModel, diags *diag.Diagnostics) {
	if id, ok := GetInt64(result, "id"); ok {
		model.ID = types.Int64Value(id)
	}

	if name, ok := GetString(result, "name"); ok {
		model.Name = types.StringValue(name)
	}

	if label, ok := GetString(result, "label"); ok && label != "" {
		model.Label = types.StringValue(label)
	} else {
		// If label is empty, use name as label
		if name, ok := GetString(result, "name"); ok {
			model.Label = types.StringValue(name)
		}
	}

	if description, ok := GetString(result, "description"); ok && description != "" {
		model.Description = types.StringValue(description)
	} else {
		model.Description = types.StringNull()
	}

	if parentID, ok := GetInt64(result, "parent_id"); ok {
		model.ParentID = types.Int64Value(parentID)
	} else {
		model.ParentID = types.Int64Null()
	}

	if isSelectable, ok := GetBool(result, "is_selectable"); ok {
		model.IsSelectable = types.BoolValue(isSelectable)
	}

	if isReserved, ok := GetBool(result, "is_reserved"); ok {
		model.IsReserved = types.BoolValue(isReserved)
	}

	if isTagset, ok := GetBool(result, "is_tagset"); ok {
		model.IsTagset = types.BoolValue(isTagset)
	}

	// Handle used_for
	if usedForRaw, ok := result["used_for"]; ok && usedForRaw != nil {
		if usedForSlice, ok := usedForRaw.([]any); ok {
			values := make([]string, 0, len(usedForSlice))
			for _, v := range usedForSlice {
				if s, ok := v.(string); ok {
					values = append(values, s)
				}
			}
			if len(values) > 0 {
				valueList, d := types.ListValueFrom(ctx, types.StringType, values)
				diags.Append(d...)
				model.UsedFor = valueList
			} else {
				model.UsedFor = types.ListNull(types.StringType)
			}
		} else {
			model.UsedFor = types.ListNull(types.StringType)
		}
	} else {
		model.UsedFor = types.ListNull(types.StringType)
	}

	if color, ok := GetString(result, "color"); ok && color != "" {
		model.Color = types.StringValue(color)
	} else {
		model.Color = types.StringNull()
	}
}
