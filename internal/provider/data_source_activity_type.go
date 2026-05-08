package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &ActivityTypeDataSource{}
var _ datasource.DataSourceWithConfigure = &ActivityTypeDataSource{}

type ActivityTypeDataSource struct {
	client *Client
}

type ActivityTypeDataSourceModel struct {
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

func NewActivityTypeDataSource() datasource.DataSource {
	return &ActivityTypeDataSource{}
}

func (d *ActivityTypeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_activity_type"
}

func (d *ActivityTypeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches a CiviCRM Activity Type (OptionValue in activity_type group) by ID or name.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Unique ID of the OptionValue. Specify either id or name.",
				Optional:    true,
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Machine name of the activity type. Specify either id or name.",
				Optional:    true,
				Computed:    true,
			},
			"label":       schema.StringAttribute{Description: "Display label.", Computed: true},
			"description": schema.StringAttribute{Description: "Description of the activity type.", Computed: true},
			"is_active":   schema.BoolAttribute{Description: "Whether active.", Computed: true},
			"is_reserved": schema.BoolAttribute{Description: "Whether reserved by system.", Computed: true},
			"weight":      schema.Int64Attribute{Description: "Sort weight.", Computed: true},
			"value": schema.StringAttribute{
				Description: "Internal numeric value used by CiviCRM to identify this activity type.",
				Computed:    true,
			},
			"color": schema.StringAttribute{Description: "Hex color code for calendar/timeline display.", Computed: true},
			"icon":  schema.StringAttribute{Description: "CSS icon class.", Computed: true},
		},
	}
}

func (d *ActivityTypeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *Client, got: %T.", req.ProviderData))
		return
	}
	d.client = client
}

func (d *ActivityTypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ActivityTypeDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.ID.IsNull() && config.Name.IsNull() {
		resp.Diagnostics.AddError("Missing Filter", "At least one of 'id' or 'name' must be specified.")
		return
	}

	where := [][]any{
		{"option_group_id:name", "=", "activity_type"},
	}
	if !config.ID.IsNull() {
		where = append(where, []any{"id", "=", config.ID.ValueInt64()})
	}
	if !config.Name.IsNull() {
		where = append(where, []any{"name", "=", config.Name.ValueString()})
	}

	tflog.Debug(ctx, "Reading ActivityType data source", map[string]any{"where": where})

	results, err := d.client.Get("OptionValue", where, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading ActivityType", err.Error())
		return
	}
	if len(results) == 0 {
		resp.Diagnostics.AddError("ActivityType not found", "No activity type found matching the specified criteria.")
		return
	}

	result := results[0]

	if id, ok := GetInt64(result, "id"); ok {
		config.ID = types.Int64Value(id)
	}
	if v, ok := GetString(result, "name"); ok {
		config.Name = types.StringValue(v)
	}
	if v, ok := GetString(result, "label"); ok {
		config.Label = types.StringValue(v)
	}
	if v, ok := GetString(result, "description"); ok && v != "" {
		config.Description = types.StringValue(v)
	} else {
		config.Description = types.StringNull()
	}
	if v, ok := GetBool(result, "is_active"); ok {
		config.IsActive = types.BoolValue(v)
	}
	if v, ok := GetBool(result, "is_reserved"); ok {
		config.IsReserved = types.BoolValue(v)
	}
	if v, ok := GetInt64(result, "weight"); ok {
		config.Weight = types.Int64Value(v)
	}
	if v, ok := GetString(result, "value"); ok && v != "" {
		config.Value = types.StringValue(v)
	} else {
		config.Value = types.StringNull()
	}
	if v, ok := GetString(result, "color"); ok && v != "" {
		config.Color = types.StringValue(v)
	} else {
		config.Color = types.StringNull()
	}
	if v, ok := GetString(result, "icon"); ok && v != "" {
		config.Icon = types.StringValue(v)
	} else {
		config.Icon = types.StringNull()
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
