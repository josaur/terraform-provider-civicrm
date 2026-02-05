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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &CustomGroupResource{}
	_ resource.ResourceWithConfigure   = &CustomGroupResource{}
	_ resource.ResourceWithImportState = &CustomGroupResource{}
)

// CustomGroupResource manages custom field groups in CiviCRM.
type CustomGroupResource struct {
	client *Client
}

type CustomGroupResourceModel struct {
	ID                        types.Int64  `tfsdk:"id"`
	Name                      types.String `tfsdk:"name"`
	Title                     types.String `tfsdk:"title"`
	Extends                   types.String `tfsdk:"extends"`
	ExtendsEntityColumnID     types.Int64  `tfsdk:"extends_entity_column_id"`
	ExtendsEntityColumnValue  types.List   `tfsdk:"extends_entity_column_value"`
	Style                     types.String `tfsdk:"style"`
	CollapseDisplay           types.Bool   `tfsdk:"collapse_display"`
	HelpPre                   types.String `tfsdk:"help_pre"`
	HelpPost                  types.String `tfsdk:"help_post"`
	Weight                    types.Int64  `tfsdk:"weight"`
	IsActive                  types.Bool   `tfsdk:"is_active"`
	TableName                 types.String `tfsdk:"table_name"`
	IsMultiple                types.Bool   `tfsdk:"is_multiple"`
	MinMultiple               types.Int64  `tfsdk:"min_multiple"`
	MaxMultiple               types.Int64  `tfsdk:"max_multiple"`
	CollapseAdvDisplay        types.Bool   `tfsdk:"collapse_adv_display"`
	IsReserved                types.Bool   `tfsdk:"is_reserved"`
	IsPublic                  types.Bool   `tfsdk:"is_public"`
	Icon                      types.String `tfsdk:"icon"`
}

func NewCustomGroupResource() resource.Resource {
	return &CustomGroupResource{}
}

func (r *CustomGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_group"
}

func (r *CustomGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages CiviCRM Custom Field Groups. Custom groups organize custom fields that extend CiviCRM entities.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier of the custom group.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The machine name of the custom group (must be unique).",
				Required:    true,
			},
			"title": schema.StringAttribute{
				Description: "The display title of the custom group.",
				Required:    true,
			},
			"extends": schema.StringAttribute{
				Description: "The entity type this custom group extends (e.g., 'Contact', 'Organization', 'Individual', 'Household', 'Activity', 'Contribution', etc.).",
				Required:    true,
			},
			"extends_entity_column_id": schema.Int64Attribute{
				Description: "For extending specific subtypes, the column ID.",
				Optional:    true,
			},
			"extends_entity_column_value": schema.ListAttribute{
				Description: "For extending specific subtypes, the allowed values.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"style": schema.StringAttribute{
				Description: "The display style. Options: 'Inline', 'Tab', 'Tab with table'. Default: 'Inline'.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("Inline"),
			},
			"collapse_display": schema.BoolAttribute{
				Description: "Whether to collapse the group display by default. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"help_pre": schema.StringAttribute{
				Description: "Help text displayed before the custom fields.",
				Optional:    true,
			},
			"help_post": schema.StringAttribute{
				Description: "Help text displayed after the custom fields.",
				Optional:    true,
			},
			"weight": schema.Int64Attribute{
				Description: "The display order weight. Default: 1.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether the custom group is active. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"table_name": schema.StringAttribute{
				Description: "The database table name for storing custom field values. Auto-generated if not specified.",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"is_multiple": schema.BoolAttribute{
				Description: "Whether multiple records can be stored per entity. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"min_multiple": schema.Int64Attribute{
				Description: "Minimum number of multiple records (if is_multiple is true).",
				Optional:    true,
			},
			"max_multiple": schema.Int64Attribute{
				Description: "Maximum number of multiple records (if is_multiple is true).",
				Optional:    true,
			},
			"collapse_adv_display": schema.BoolAttribute{
				Description: "Whether to collapse in advanced search display. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"is_reserved": schema.BoolAttribute{
				Description: "Whether this is a reserved system group. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"is_public": schema.BoolAttribute{
				Description: "Whether this group is visible on public forms. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"icon": schema.StringAttribute{
				Description: "The icon for the custom group (CSS class name).",
				Optional:    true,
			},
		},
	}
}

func (r *CustomGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CustomGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CustomGroupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating custom group", map[string]any{
		"name":  plan.Name.ValueString(),
		"title": plan.Title.ValueString(),
	})

	// Build values for API call
	values := map[string]any{
		"name":                 plan.Name.ValueString(),
		"title":                plan.Title.ValueString(),
		"extends":              plan.Extends.ValueString(),
		"style":                plan.Style.ValueString(),
		"collapse_display":     plan.CollapseDisplay.ValueBool(),
		"weight":               plan.Weight.ValueInt64(),
		"is_active":            plan.IsActive.ValueBool(),
		"is_multiple":          plan.IsMultiple.ValueBool(),
		"collapse_adv_display": plan.CollapseAdvDisplay.ValueBool(),
		"is_reserved":          plan.IsReserved.ValueBool(),
		"is_public":            plan.IsPublic.ValueBool(),
	}

	if !plan.ExtendsEntityColumnID.IsNull() {
		values["extends_entity_column_id"] = plan.ExtendsEntityColumnID.ValueInt64()
	}

	if !plan.ExtendsEntityColumnValue.IsNull() {
		var columnValues []string
		diags = plan.ExtendsEntityColumnValue.ElementsAs(ctx, &columnValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		values["extends_entity_column_value"] = columnValues
	}

	if !plan.HelpPre.IsNull() {
		values["help_pre"] = plan.HelpPre.ValueString()
	}

	if !plan.HelpPost.IsNull() {
		values["help_post"] = plan.HelpPost.ValueString()
	}

	if !plan.TableName.IsNull() {
		values["table_name"] = plan.TableName.ValueString()
	}

	if !plan.MinMultiple.IsNull() {
		values["min_multiple"] = plan.MinMultiple.ValueInt64()
	}

	if !plan.MaxMultiple.IsNull() {
		values["max_multiple"] = plan.MaxMultiple.ValueInt64()
	}

	if !plan.Icon.IsNull() {
		values["icon"] = plan.Icon.ValueString()
	}

	// Call API
	result, err := r.client.Create("CustomGroup", values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating custom group",
			"Could not create custom group, unexpected error: "+err.Error(),
		)
		return
	}

	// Update state with response
	var d diag.Diagnostics
	r.mapResponseToModel(ctx, result, &plan, &d)
	resp.Diagnostics.Append(d...)

	tflog.Debug(ctx, "Created custom group", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CustomGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CustomGroupResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading custom group", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	result, err := r.client.GetByID("CustomGroup", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading custom group",
			"Could not read custom group ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
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

func (r *CustomGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CustomGroupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state CustomGroupResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating custom group", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	// Build values for API call
	values := map[string]any{
		"name":                 plan.Name.ValueString(),
		"title":                plan.Title.ValueString(),
		"extends":              plan.Extends.ValueString(),
		"style":                plan.Style.ValueString(),
		"collapse_display":     plan.CollapseDisplay.ValueBool(),
		"weight":               plan.Weight.ValueInt64(),
		"is_active":            plan.IsActive.ValueBool(),
		"is_multiple":          plan.IsMultiple.ValueBool(),
		"collapse_adv_display": plan.CollapseAdvDisplay.ValueBool(),
		"is_reserved":          plan.IsReserved.ValueBool(),
		"is_public":            plan.IsPublic.ValueBool(),
	}

	if !plan.ExtendsEntityColumnID.IsNull() {
		values["extends_entity_column_id"] = plan.ExtendsEntityColumnID.ValueInt64()
	} else {
		values["extends_entity_column_id"] = nil
	}

	if !plan.ExtendsEntityColumnValue.IsNull() {
		var columnValues []string
		diags = plan.ExtendsEntityColumnValue.ElementsAs(ctx, &columnValues, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		values["extends_entity_column_value"] = columnValues
	} else {
		values["extends_entity_column_value"] = nil
	}

	if !plan.HelpPre.IsNull() {
		values["help_pre"] = plan.HelpPre.ValueString()
	} else {
		values["help_pre"] = nil
	}

	if !plan.HelpPost.IsNull() {
		values["help_post"] = plan.HelpPost.ValueString()
	} else {
		values["help_post"] = nil
	}

	if !plan.MinMultiple.IsNull() {
		values["min_multiple"] = plan.MinMultiple.ValueInt64()
	} else {
		values["min_multiple"] = nil
	}

	if !plan.MaxMultiple.IsNull() {
		values["max_multiple"] = plan.MaxMultiple.ValueInt64()
	} else {
		values["max_multiple"] = nil
	}

	if !plan.Icon.IsNull() {
		values["icon"] = plan.Icon.ValueString()
	} else {
		values["icon"] = nil
	}

	// Call API
	result, err := r.client.Update("CustomGroup", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating custom group",
			"Could not update custom group ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Update state
	plan.ID = state.ID
	var d diag.Diagnostics
	r.mapResponseToModel(ctx, result, &plan, &d)
	resp.Diagnostics.Append(d...)

	tflog.Debug(ctx, "Updated custom group", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CustomGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CustomGroupResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting custom group", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	err := r.client.Delete("CustomGroup", state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting custom group",
			"Could not delete custom group ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "Deleted custom group", map[string]any{
		"id": state.ID.ValueInt64(),
	})
}

func (r *CustomGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (r *CustomGroupResource) mapResponseToModel(ctx context.Context, result map[string]any, model *CustomGroupResourceModel, diags *diag.Diagnostics) {
	if id, ok := GetInt64(result, "id"); ok {
		model.ID = types.Int64Value(id)
	}

	if name, ok := GetString(result, "name"); ok {
		model.Name = types.StringValue(name)
	}

	if title, ok := GetString(result, "title"); ok {
		model.Title = types.StringValue(title)
	}

	if extends, ok := GetString(result, "extends"); ok {
		model.Extends = types.StringValue(extends)
	}

	if columnID, ok := GetInt64(result, "extends_entity_column_id"); ok {
		model.ExtendsEntityColumnID = types.Int64Value(columnID)
	} else {
		model.ExtendsEntityColumnID = types.Int64Null()
	}

	// Handle extends_entity_column_value
	if columnValueRaw, ok := result["extends_entity_column_value"]; ok && columnValueRaw != nil {
		if columnValueSlice, ok := columnValueRaw.([]any); ok {
			values := make([]string, 0, len(columnValueSlice))
			for _, v := range columnValueSlice {
				if s, ok := v.(string); ok {
					values = append(values, s)
				}
			}
			if len(values) > 0 {
				valueList, d := types.ListValueFrom(ctx, types.StringType, values)
				diags.Append(d...)
				model.ExtendsEntityColumnValue = valueList
			} else {
				model.ExtendsEntityColumnValue = types.ListNull(types.StringType)
			}
		} else {
			model.ExtendsEntityColumnValue = types.ListNull(types.StringType)
		}
	} else {
		model.ExtendsEntityColumnValue = types.ListNull(types.StringType)
	}

	if style, ok := GetString(result, "style"); ok {
		model.Style = types.StringValue(style)
	}

	if collapseDisplay, ok := GetBool(result, "collapse_display"); ok {
		model.CollapseDisplay = types.BoolValue(collapseDisplay)
	}

	if helpPre, ok := GetString(result, "help_pre"); ok && helpPre != "" {
		model.HelpPre = types.StringValue(helpPre)
	} else {
		model.HelpPre = types.StringNull()
	}

	if helpPost, ok := GetString(result, "help_post"); ok && helpPost != "" {
		model.HelpPost = types.StringValue(helpPost)
	} else {
		model.HelpPost = types.StringNull()
	}

	if weight, ok := GetInt64(result, "weight"); ok {
		model.Weight = types.Int64Value(weight)
	}

	if isActive, ok := GetBool(result, "is_active"); ok {
		model.IsActive = types.BoolValue(isActive)
	}

	if tableName, ok := GetString(result, "table_name"); ok {
		model.TableName = types.StringValue(tableName)
	}

	if isMultiple, ok := GetBool(result, "is_multiple"); ok {
		model.IsMultiple = types.BoolValue(isMultiple)
	}

	if minMultiple, ok := GetInt64(result, "min_multiple"); ok {
		model.MinMultiple = types.Int64Value(minMultiple)
	} else {
		model.MinMultiple = types.Int64Null()
	}

	if maxMultiple, ok := GetInt64(result, "max_multiple"); ok {
		model.MaxMultiple = types.Int64Value(maxMultiple)
	} else {
		model.MaxMultiple = types.Int64Null()
	}

	if collapseAdvDisplay, ok := GetBool(result, "collapse_adv_display"); ok {
		model.CollapseAdvDisplay = types.BoolValue(collapseAdvDisplay)
	}

	if isReserved, ok := GetBool(result, "is_reserved"); ok {
		model.IsReserved = types.BoolValue(isReserved)
	}

	if isPublic, ok := GetBool(result, "is_public"); ok {
		model.IsPublic = types.BoolValue(isPublic)
	}

	if icon, ok := GetString(result, "icon"); ok && icon != "" {
		model.Icon = types.StringValue(icon)
	} else {
		model.Icon = types.StringNull()
	}
}
