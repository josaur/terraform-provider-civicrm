package provider

import (
	"context"
	"fmt"
	"strconv"

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
	_ resource.Resource                = &CustomFieldResource{}
	_ resource.ResourceWithConfigure   = &CustomFieldResource{}
	_ resource.ResourceWithImportState = &CustomFieldResource{}
)

// CustomFieldResource manages custom fields in CiviCRM.
type CustomFieldResource struct {
	client *Client
}

type CustomFieldResourceModel struct {
	ID              types.Int64  `tfsdk:"id"`
	CustomGroupID   types.Int64  `tfsdk:"custom_group_id"`
	Name            types.String `tfsdk:"name"`
	Label           types.String `tfsdk:"label"`
	DataType        types.String `tfsdk:"data_type"`
	HtmlType        types.String `tfsdk:"html_type"`
	DefaultValue    types.String `tfsdk:"default_value"`
	IsRequired      types.Bool   `tfsdk:"is_required"`
	IsSearchable    types.Bool   `tfsdk:"is_searchable"`
	IsSearchRange   types.Bool   `tfsdk:"is_search_range"`
	Weight          types.Int64  `tfsdk:"weight"`
	HelpPre         types.String `tfsdk:"help_pre"`
	HelpPost        types.String `tfsdk:"help_post"`
	Attributes      types.String `tfsdk:"attributes"`
	IsActive        types.Bool   `tfsdk:"is_active"`
	IsView          types.Bool   `tfsdk:"is_view"`
	OptionsPerLine  types.Int64  `tfsdk:"options_per_line"`
	TextLength      types.Int64  `tfsdk:"text_length"`
	StartDateYears  types.Int64  `tfsdk:"start_date_years"`
	EndDateYears    types.Int64  `tfsdk:"end_date_years"`
	DateFormat      types.String `tfsdk:"date_format"`
	TimeFormat      types.Int64  `tfsdk:"time_format"`
	NoteColumns     types.Int64  `tfsdk:"note_columns"`
	NoteRows        types.Int64  `tfsdk:"note_rows"`
	ColumnName      types.String `tfsdk:"column_name"`
	OptionGroupID   types.Int64  `tfsdk:"option_group_id"`
	Serialize       types.Int64  `tfsdk:"serialize"`
	Filter          types.String `tfsdk:"filter"`
	InSelector      types.Bool   `tfsdk:"in_selector"`
	FkEntity        types.String `tfsdk:"fk_entity"`
	FkEntityOnDelete types.String `tfsdk:"fk_entity_on_delete"`
}

func NewCustomFieldResource() resource.Resource {
	return &CustomFieldResource{}
}

func (r *CustomFieldResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_field"
}

func (r *CustomFieldResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages CiviCRM Custom Fields within Custom Groups.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier of the custom field.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"custom_group_id": schema.Int64Attribute{
				Description: "The ID of the custom group this field belongs to.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The machine name of the custom field (must be unique within the group).",
				Required:    true,
			},
			"label": schema.StringAttribute{
				Description: "The display label of the custom field.",
				Required:    true,
			},
			"data_type": schema.StringAttribute{
				Description: "The data type. Options: 'String', 'Int', 'Float', 'Money', 'Memo', 'Date', 'Boolean', 'StateProvince', 'Country', 'File', 'Link', 'ContactReference', 'EntityReference'.",
				Required:    true,
			},
			"html_type": schema.StringAttribute{
				Description: "The HTML input type. Options: 'Text', 'TextArea', 'Select', 'Multi-Select', 'AdvMulti-Select', 'Radio', 'CheckBox', 'Select Date', 'Select State/Province', 'Select Country', 'File', 'Link', 'RichTextEditor', 'Autocomplete-Select', 'EntityRef'.",
				Required:    true,
			},
			"default_value": schema.StringAttribute{
				Description: "The default value for the field.",
				Optional:    true,
			},
			"is_required": schema.BoolAttribute{
				Description: "Whether the field is required. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"is_searchable": schema.BoolAttribute{
				Description: "Whether the field is searchable. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"is_search_range": schema.BoolAttribute{
				Description: "Whether to enable range search for this field. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"weight": schema.Int64Attribute{
				Description: "The display order weight. Default: 1.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
			},
			"help_pre": schema.StringAttribute{
				Description: "Help text displayed before the field.",
				Optional:    true,
			},
			"help_post": schema.StringAttribute{
				Description: "Help text displayed after the field.",
				Optional:    true,
			},
			"attributes": schema.StringAttribute{
				Description: "Additional HTML attributes for the field.",
				Optional:    true,
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether the field is active. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"is_view": schema.BoolAttribute{
				Description: "Whether the field is view-only. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"options_per_line": schema.Int64Attribute{
				Description: "Number of options to display per line (for Radio/CheckBox).",
				Optional:    true,
			},
			"text_length": schema.Int64Attribute{
				Description: "Maximum text length for text fields. Default: 255.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(255),
			},
			"start_date_years": schema.Int64Attribute{
				Description: "Number of years before current date for date picker start.",
				Optional:    true,
			},
			"end_date_years": schema.Int64Attribute{
				Description: "Number of years after current date for date picker end.",
				Optional:    true,
			},
			"date_format": schema.StringAttribute{
				Description: "The date format string.",
				Optional:    true,
			},
			"time_format": schema.Int64Attribute{
				Description: "The time format (1 for 12-hour, 2 for 24-hour).",
				Optional:    true,
			},
			"note_columns": schema.Int64Attribute{
				Description: "Number of columns for note/textarea fields. Default: 60.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(60),
			},
			"note_rows": schema.Int64Attribute{
				Description: "Number of rows for note/textarea fields. Default: 4.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(4),
			},
			"column_name": schema.StringAttribute{
				Description: "The database column name. Auto-generated if not specified.",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"option_group_id": schema.Int64Attribute{
				Description: "The ID of the option group for Select/Radio/CheckBox fields.",
				Optional:    true,
			},
			"serialize": schema.Int64Attribute{
				Description: "Serialization method (0 for none, 1 for separator). Default: 0.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
			},
			"filter": schema.StringAttribute{
				Description: "Filter for entity reference fields.",
				Optional:    true,
			},
			"in_selector": schema.BoolAttribute{
				Description: "Whether to include in selector. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"fk_entity": schema.StringAttribute{
				Description: "Foreign key entity for EntityReference fields.",
				Optional:    true,
			},
			"fk_entity_on_delete": schema.StringAttribute{
				Description: "Action on delete for foreign key. Options: 'cascade', 'set_null'. Default: 'set_null'.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("set_null"),
			},
		},
	}
}

func (r *CustomFieldResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CustomFieldResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CustomFieldResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating custom field", map[string]any{
		"name":            plan.Name.ValueString(),
		"custom_group_id": plan.CustomGroupID.ValueInt64(),
	})

	// Build values for API call
	values := map[string]any{
		"custom_group_id":  plan.CustomGroupID.ValueInt64(),
		"name":             plan.Name.ValueString(),
		"label":            plan.Label.ValueString(),
		"data_type":        plan.DataType.ValueString(),
		"html_type":        plan.HtmlType.ValueString(),
		"is_required":      plan.IsRequired.ValueBool(),
		"is_searchable":    plan.IsSearchable.ValueBool(),
		"is_search_range":  plan.IsSearchRange.ValueBool(),
		"weight":           plan.Weight.ValueInt64(),
		"is_active":        plan.IsActive.ValueBool(),
		"is_view":          plan.IsView.ValueBool(),
		"text_length":      plan.TextLength.ValueInt64(),
		"note_columns":     plan.NoteColumns.ValueInt64(),
		"note_rows":        plan.NoteRows.ValueInt64(),
		"serialize":        plan.Serialize.ValueInt64(),
		"in_selector":      plan.InSelector.ValueBool(),
		"fk_entity_on_delete": plan.FkEntityOnDelete.ValueString(),
	}

	if !plan.DefaultValue.IsNull() {
		values["default_value"] = plan.DefaultValue.ValueString()
	}

	if !plan.HelpPre.IsNull() {
		values["help_pre"] = plan.HelpPre.ValueString()
	}

	if !plan.HelpPost.IsNull() {
		values["help_post"] = plan.HelpPost.ValueString()
	}

	if !plan.Attributes.IsNull() {
		values["attributes"] = plan.Attributes.ValueString()
	}

	if !plan.OptionsPerLine.IsNull() {
		values["options_per_line"] = plan.OptionsPerLine.ValueInt64()
	}

	if !plan.StartDateYears.IsNull() {
		values["start_date_years"] = plan.StartDateYears.ValueInt64()
	}

	if !plan.EndDateYears.IsNull() {
		values["end_date_years"] = plan.EndDateYears.ValueInt64()
	}

	if !plan.DateFormat.IsNull() {
		values["date_format"] = plan.DateFormat.ValueString()
	}

	if !plan.TimeFormat.IsNull() {
		values["time_format"] = plan.TimeFormat.ValueInt64()
	}

	if !plan.ColumnName.IsNull() {
		values["column_name"] = plan.ColumnName.ValueString()
	}

	if !plan.OptionGroupID.IsNull() {
		values["option_group_id"] = plan.OptionGroupID.ValueInt64()
	}

	if !plan.Filter.IsNull() {
		values["filter"] = plan.Filter.ValueString()
	}

	if !plan.FkEntity.IsNull() {
		values["fk_entity"] = plan.FkEntity.ValueString()
	}

	// Call API
	result, err := r.client.Create("CustomField", values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating custom field",
			"Could not create custom field, unexpected error: "+err.Error(),
		)
		return
	}

	// Update state with response
	r.mapResponseToModel(result, &plan)

	tflog.Debug(ctx, "Created custom field", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CustomFieldResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CustomFieldResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading custom field", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	result, err := r.client.GetByID("CustomField", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading custom field",
			"Could not read custom field ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Update state
	r.mapResponseToModel(result, &state)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *CustomFieldResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CustomFieldResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state CustomFieldResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating custom field", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	// Build values for API call
	values := map[string]any{
		"custom_group_id":  plan.CustomGroupID.ValueInt64(),
		"name":             plan.Name.ValueString(),
		"label":            plan.Label.ValueString(),
		"data_type":        plan.DataType.ValueString(),
		"html_type":        plan.HtmlType.ValueString(),
		"is_required":      plan.IsRequired.ValueBool(),
		"is_searchable":    plan.IsSearchable.ValueBool(),
		"is_search_range":  plan.IsSearchRange.ValueBool(),
		"weight":           plan.Weight.ValueInt64(),
		"is_active":        plan.IsActive.ValueBool(),
		"is_view":          plan.IsView.ValueBool(),
		"text_length":      plan.TextLength.ValueInt64(),
		"note_columns":     plan.NoteColumns.ValueInt64(),
		"note_rows":        plan.NoteRows.ValueInt64(),
		"serialize":        plan.Serialize.ValueInt64(),
		"in_selector":      plan.InSelector.ValueBool(),
		"fk_entity_on_delete": plan.FkEntityOnDelete.ValueString(),
	}

	if !plan.DefaultValue.IsNull() {
		values["default_value"] = plan.DefaultValue.ValueString()
	} else {
		values["default_value"] = nil
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

	if !plan.Attributes.IsNull() {
		values["attributes"] = plan.Attributes.ValueString()
	} else {
		values["attributes"] = nil
	}

	if !plan.OptionsPerLine.IsNull() {
		values["options_per_line"] = plan.OptionsPerLine.ValueInt64()
	} else {
		values["options_per_line"] = nil
	}

	if !plan.StartDateYears.IsNull() {
		values["start_date_years"] = plan.StartDateYears.ValueInt64()
	} else {
		values["start_date_years"] = nil
	}

	if !plan.EndDateYears.IsNull() {
		values["end_date_years"] = plan.EndDateYears.ValueInt64()
	} else {
		values["end_date_years"] = nil
	}

	if !plan.DateFormat.IsNull() {
		values["date_format"] = plan.DateFormat.ValueString()
	} else {
		values["date_format"] = nil
	}

	if !plan.TimeFormat.IsNull() {
		values["time_format"] = plan.TimeFormat.ValueInt64()
	} else {
		values["time_format"] = nil
	}

	if !plan.OptionGroupID.IsNull() {
		values["option_group_id"] = plan.OptionGroupID.ValueInt64()
	} else {
		values["option_group_id"] = nil
	}

	if !plan.Filter.IsNull() {
		values["filter"] = plan.Filter.ValueString()
	} else {
		values["filter"] = nil
	}

	if !plan.FkEntity.IsNull() {
		values["fk_entity"] = plan.FkEntity.ValueString()
	} else {
		values["fk_entity"] = nil
	}

	// Call API
	result, err := r.client.Update("CustomField", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating custom field",
			"Could not update custom field ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Update state
	plan.ID = state.ID
	r.mapResponseToModel(result, &plan)

	tflog.Debug(ctx, "Updated custom field", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CustomFieldResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CustomFieldResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting custom field", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	err := r.client.Delete("CustomField", state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting custom field",
			"Could not delete custom field ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "Deleted custom field", map[string]any{
		"id": state.ID.ValueInt64(),
	})
}

func (r *CustomFieldResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (r *CustomFieldResource) mapResponseToModel(result map[string]any, model *CustomFieldResourceModel) {
	if id, ok := GetInt64(result, "id"); ok {
		model.ID = types.Int64Value(id)
	}

	if customGroupID, ok := GetInt64(result, "custom_group_id"); ok {
		model.CustomGroupID = types.Int64Value(customGroupID)
	}

	if name, ok := GetString(result, "name"); ok {
		model.Name = types.StringValue(name)
	}

	if label, ok := GetString(result, "label"); ok {
		model.Label = types.StringValue(label)
	}

	if dataType, ok := GetString(result, "data_type"); ok {
		model.DataType = types.StringValue(dataType)
	}

	if htmlType, ok := GetString(result, "html_type"); ok {
		model.HtmlType = types.StringValue(htmlType)
	}

	if defaultValue, ok := GetString(result, "default_value"); ok && defaultValue != "" {
		model.DefaultValue = types.StringValue(defaultValue)
	} else {
		model.DefaultValue = types.StringNull()
	}

	if isRequired, ok := GetBool(result, "is_required"); ok {
		model.IsRequired = types.BoolValue(isRequired)
	}

	if isSearchable, ok := GetBool(result, "is_searchable"); ok {
		model.IsSearchable = types.BoolValue(isSearchable)
	}

	if isSearchRange, ok := GetBool(result, "is_search_range"); ok {
		model.IsSearchRange = types.BoolValue(isSearchRange)
	}

	if weight, ok := GetInt64(result, "weight"); ok {
		model.Weight = types.Int64Value(weight)
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

	if attributes, ok := GetString(result, "attributes"); ok && attributes != "" {
		model.Attributes = types.StringValue(attributes)
	} else {
		model.Attributes = types.StringNull()
	}

	if isActive, ok := GetBool(result, "is_active"); ok {
		model.IsActive = types.BoolValue(isActive)
	}

	if isView, ok := GetBool(result, "is_view"); ok {
		model.IsView = types.BoolValue(isView)
	}

	if optionsPerLine, ok := GetInt64(result, "options_per_line"); ok {
		model.OptionsPerLine = types.Int64Value(optionsPerLine)
	} else {
		model.OptionsPerLine = types.Int64Null()
	}

	if textLength, ok := GetInt64(result, "text_length"); ok {
		model.TextLength = types.Int64Value(textLength)
	}

	if startDateYears, ok := GetInt64(result, "start_date_years"); ok {
		model.StartDateYears = types.Int64Value(startDateYears)
	} else {
		model.StartDateYears = types.Int64Null()
	}

	if endDateYears, ok := GetInt64(result, "end_date_years"); ok {
		model.EndDateYears = types.Int64Value(endDateYears)
	} else {
		model.EndDateYears = types.Int64Null()
	}

	if dateFormat, ok := GetString(result, "date_format"); ok && dateFormat != "" {
		model.DateFormat = types.StringValue(dateFormat)
	} else {
		model.DateFormat = types.StringNull()
	}

	if timeFormat, ok := GetInt64(result, "time_format"); ok {
		model.TimeFormat = types.Int64Value(timeFormat)
	} else {
		model.TimeFormat = types.Int64Null()
	}

	if noteColumns, ok := GetInt64(result, "note_columns"); ok {
		model.NoteColumns = types.Int64Value(noteColumns)
	}

	if noteRows, ok := GetInt64(result, "note_rows"); ok {
		model.NoteRows = types.Int64Value(noteRows)
	}

	if columnName, ok := GetString(result, "column_name"); ok {
		model.ColumnName = types.StringValue(columnName)
	}

	if optionGroupID, ok := GetInt64(result, "option_group_id"); ok {
		model.OptionGroupID = types.Int64Value(optionGroupID)
	} else {
		model.OptionGroupID = types.Int64Null()
	}

	if serialize, ok := GetInt64(result, "serialize"); ok {
		model.Serialize = types.Int64Value(serialize)
	}

	if filter, ok := GetString(result, "filter"); ok && filter != "" {
		model.Filter = types.StringValue(filter)
	} else {
		model.Filter = types.StringNull()
	}

	if inSelector, ok := GetBool(result, "in_selector"); ok {
		model.InSelector = types.BoolValue(inSelector)
	}

	if fkEntity, ok := GetString(result, "fk_entity"); ok && fkEntity != "" {
		model.FkEntity = types.StringValue(fkEntity)
	} else {
		model.FkEntity = types.StringNull()
	}

	if fkEntityOnDelete, ok := GetString(result, "fk_entity_on_delete"); ok {
		model.FkEntityOnDelete = types.StringValue(fkEntityOnDelete)
	}
}
