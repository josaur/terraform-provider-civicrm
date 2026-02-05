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
	_ resource.Resource                = &RelationshipTypeResource{}
	_ resource.ResourceWithConfigure   = &RelationshipTypeResource{}
	_ resource.ResourceWithImportState = &RelationshipTypeResource{}
)

// RelationshipTypeResource manages relationship types in CiviCRM.
type RelationshipTypeResource struct {
	client *Client
}

type RelationshipTypeResourceModel struct {
	ID               types.Int64  `tfsdk:"id"`
	NameAB           types.String `tfsdk:"name_a_b"`
	LabelAB          types.String `tfsdk:"label_a_b"`
	NameBA           types.String `tfsdk:"name_b_a"`
	LabelBA          types.String `tfsdk:"label_b_a"`
	Description      types.String `tfsdk:"description"`
	ContactTypeA     types.String `tfsdk:"contact_type_a"`
	ContactTypeB     types.String `tfsdk:"contact_type_b"`
	ContactSubTypeA  types.String `tfsdk:"contact_sub_type_a"`
	ContactSubTypeB  types.String `tfsdk:"contact_sub_type_b"`
	IsReserved       types.Bool   `tfsdk:"is_reserved"`
	IsActive         types.Bool   `tfsdk:"is_active"`
}

func NewRelationshipTypeResource() resource.Resource {
	return &RelationshipTypeResource{}
}

func (r *RelationshipTypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_relationship_type"
}

func (r *RelationshipTypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages CiviCRM Relationship Types that define how contacts can be related to each other.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier of the relationship type.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name_a_b": schema.StringAttribute{
				Description: "The relationship name from A to B perspective (e.g., 'Child of').",
				Required:    true,
			},
			"label_a_b": schema.StringAttribute{
				Description: "The display label from A to B perspective.",
				Required:    true,
			},
			"name_b_a": schema.StringAttribute{
				Description: "The relationship name from B to A perspective (e.g., 'Parent of').",
				Required:    true,
			},
			"label_b_a": schema.StringAttribute{
				Description: "The display label from B to A perspective.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the relationship type.",
				Optional:    true,
			},
			"contact_type_a": schema.StringAttribute{
				Description: "The contact type for side A (e.g., 'Individual', 'Organization', 'Household'). Leave empty for any type.",
				Optional:    true,
			},
			"contact_type_b": schema.StringAttribute{
				Description: "The contact type for side B (e.g., 'Individual', 'Organization', 'Household'). Leave empty for any type.",
				Optional:    true,
			},
			"contact_sub_type_a": schema.StringAttribute{
				Description: "The contact subtype for side A.",
				Optional:    true,
			},
			"contact_sub_type_b": schema.StringAttribute{
				Description: "The contact subtype for side B.",
				Optional:    true,
			},
			"is_reserved": schema.BoolAttribute{
				Description: "Whether this is a reserved system relationship type. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether the relationship type is active. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
		},
	}
}

func (r *RelationshipTypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RelationshipTypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RelationshipTypeResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating relationship type", map[string]any{
		"name_a_b": plan.NameAB.ValueString(),
	})

	// Build values for API call
	values := map[string]any{
		"name_a_b":    plan.NameAB.ValueString(),
		"label_a_b":   plan.LabelAB.ValueString(),
		"name_b_a":    plan.NameBA.ValueString(),
		"label_b_a":   plan.LabelBA.ValueString(),
		"is_reserved": plan.IsReserved.ValueBool(),
		"is_active":   plan.IsActive.ValueBool(),
	}

	if !plan.Description.IsNull() {
		values["description"] = plan.Description.ValueString()
	}

	if !plan.ContactTypeA.IsNull() {
		values["contact_type_a"] = plan.ContactTypeA.ValueString()
	}

	if !plan.ContactTypeB.IsNull() {
		values["contact_type_b"] = plan.ContactTypeB.ValueString()
	}

	if !plan.ContactSubTypeA.IsNull() {
		values["contact_sub_type_a"] = plan.ContactSubTypeA.ValueString()
	}

	if !plan.ContactSubTypeB.IsNull() {
		values["contact_sub_type_b"] = plan.ContactSubTypeB.ValueString()
	}

	// Call API
	result, err := r.client.Create("RelationshipType", values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating relationship type",
			"Could not create relationship type, unexpected error: "+err.Error(),
		)
		return
	}

	// Update state with response
	r.mapResponseToModel(result, &plan)

	tflog.Debug(ctx, "Created relationship type", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *RelationshipTypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RelationshipTypeResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading relationship type", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	result, err := r.client.GetByID("RelationshipType", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading relationship type",
			"Could not read relationship type ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Update state
	r.mapResponseToModel(result, &state)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *RelationshipTypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan RelationshipTypeResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state RelationshipTypeResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating relationship type", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	// Build values for API call
	values := map[string]any{
		"name_a_b":    plan.NameAB.ValueString(),
		"label_a_b":   plan.LabelAB.ValueString(),
		"name_b_a":    plan.NameBA.ValueString(),
		"label_b_a":   plan.LabelBA.ValueString(),
		"is_reserved": plan.IsReserved.ValueBool(),
		"is_active":   plan.IsActive.ValueBool(),
	}

	if !plan.Description.IsNull() {
		values["description"] = plan.Description.ValueString()
	} else {
		values["description"] = nil
	}

	if !plan.ContactTypeA.IsNull() {
		values["contact_type_a"] = plan.ContactTypeA.ValueString()
	} else {
		values["contact_type_a"] = nil
	}

	if !plan.ContactTypeB.IsNull() {
		values["contact_type_b"] = plan.ContactTypeB.ValueString()
	} else {
		values["contact_type_b"] = nil
	}

	if !plan.ContactSubTypeA.IsNull() {
		values["contact_sub_type_a"] = plan.ContactSubTypeA.ValueString()
	} else {
		values["contact_sub_type_a"] = nil
	}

	if !plan.ContactSubTypeB.IsNull() {
		values["contact_sub_type_b"] = plan.ContactSubTypeB.ValueString()
	} else {
		values["contact_sub_type_b"] = nil
	}

	// Call API
	result, err := r.client.Update("RelationshipType", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating relationship type",
			"Could not update relationship type ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Update state
	plan.ID = state.ID
	r.mapResponseToModel(result, &plan)

	tflog.Debug(ctx, "Updated relationship type", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *RelationshipTypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RelationshipTypeResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting relationship type", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	err := r.client.Delete("RelationshipType", state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting relationship type",
			"Could not delete relationship type ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "Deleted relationship type", map[string]any{
		"id": state.ID.ValueInt64(),
	})
}

func (r *RelationshipTypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (r *RelationshipTypeResource) mapResponseToModel(result map[string]any, model *RelationshipTypeResourceModel) {
	if id, ok := GetInt64(result, "id"); ok {
		model.ID = types.Int64Value(id)
	}

	if nameAB, ok := GetString(result, "name_a_b"); ok {
		model.NameAB = types.StringValue(nameAB)
	}

	if labelAB, ok := GetString(result, "label_a_b"); ok {
		model.LabelAB = types.StringValue(labelAB)
	}

	if nameBA, ok := GetString(result, "name_b_a"); ok {
		model.NameBA = types.StringValue(nameBA)
	}

	if labelBA, ok := GetString(result, "label_b_a"); ok {
		model.LabelBA = types.StringValue(labelBA)
	}

	if description, ok := GetString(result, "description"); ok && description != "" {
		model.Description = types.StringValue(description)
	} else {
		model.Description = types.StringNull()
	}

	if contactTypeA, ok := GetString(result, "contact_type_a"); ok && contactTypeA != "" {
		model.ContactTypeA = types.StringValue(contactTypeA)
	} else {
		model.ContactTypeA = types.StringNull()
	}

	if contactTypeB, ok := GetString(result, "contact_type_b"); ok && contactTypeB != "" {
		model.ContactTypeB = types.StringValue(contactTypeB)
	} else {
		model.ContactTypeB = types.StringNull()
	}

	if contactSubTypeA, ok := GetString(result, "contact_sub_type_a"); ok && contactSubTypeA != "" {
		model.ContactSubTypeA = types.StringValue(contactSubTypeA)
	} else {
		model.ContactSubTypeA = types.StringNull()
	}

	if contactSubTypeB, ok := GetString(result, "contact_sub_type_b"); ok && contactSubTypeB != "" {
		model.ContactSubTypeB = types.StringValue(contactSubTypeB)
	} else {
		model.ContactSubTypeB = types.StringNull()
	}

	if isReserved, ok := GetBool(result, "is_reserved"); ok {
		model.IsReserved = types.BoolValue(isReserved)
	}

	if isActive, ok := GetBool(result, "is_active"); ok {
		model.IsActive = types.BoolValue(isActive)
	}
}
