package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &NavigationDataSource{}
var _ datasource.DataSourceWithConfigure = &NavigationDataSource{}

type NavigationDataSource struct {
	client *Client
}

type NavigationDataSourceModel struct {
	ID                 types.Int64  `tfsdk:"id"`
	Name               types.String `tfsdk:"name"`
	Label              types.String `tfsdk:"label"`
	URL                types.String `tfsdk:"url"`
	Icon               types.String `tfsdk:"icon"`
	Permission         types.List   `tfsdk:"permission"`
	PermissionOperator types.String `tfsdk:"permission_operator"`
	ParentID           types.Int64  `tfsdk:"parent_id"`
	IsActive           types.Bool   `tfsdk:"is_active"`
	HasSeparator       types.Int64  `tfsdk:"has_separator"`
	Weight             types.Int64  `tfsdk:"weight"`
}

func NewNavigationDataSource() datasource.DataSource {
	return &NavigationDataSource{}
}

func (d *NavigationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigation"
}

func (d *NavigationDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches a CiviCRM Navigation menu entry by id or by name. Lets an existing " +
			"entry — including CiviCRM's own built-in top-level items (e.g. \"Search\", " +
			"\"Administer\") — be resolved by name, e.g. as a parent_id for a new " +
			"civicrm_navigation resource, or adopted via the import block pattern documented for " +
			"civicrm_message_template instead of being recreated as a duplicate.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier. Specify id, or name.",
				Optional:    true,
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Internal machine name of the menu entry. Specify id, or name.",
				Optional:    true,
				Computed:    true,
			},
			"label":               schema.StringAttribute{Description: "Menu title shown in the UI.", Computed: true},
			"url":                 schema.StringAttribute{Description: "Target URL, for custom links.", Computed: true},
			"icon":                schema.StringAttribute{Description: "CSS class of the icon shown next to the label.", Computed: true},
			"permission":          schema.ListAttribute{Description: "Permissions required to see this menu entry.", Computed: true, ElementType: types.StringType},
			"permission_operator": schema.StringAttribute{Description: "How multiple permission entries combine: AND or OR.", Computed: true},
			"parent_id":           schema.Int64Attribute{Description: "FK to the parent civicrm_navigation entry, if nested.", Computed: true},
			"is_active":           schema.BoolAttribute{Description: "Whether the entry is active.", Computed: true},
			"has_separator":       schema.Int64Attribute{Description: "Separator line around this entry: 0 = none, 1 = after, 2 = before.", Computed: true},
			"weight":              schema.Int64Attribute{Description: "Ordering of the entry among its siblings.", Computed: true},
		},
	}
}

func (d *NavigationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *NavigationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NavigationDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.ID.IsNull() && config.Name.IsNull() {
		resp.Diagnostics.AddError("Missing Filter", "At least one of 'id' or 'name' must be specified.")
		return
	}

	where := [][]any{}
	if !config.ID.IsNull() {
		where = append(where, []any{"id", "=", config.ID.ValueInt64()})
	}
	if !config.Name.IsNull() {
		where = append(where, []any{"name", "=", config.Name.ValueString()})
	}

	tflog.Debug(ctx, "Reading Navigation data source", map[string]any{"where": where})

	results, err := d.client.Get("Navigation", where, navigationSelect)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Navigation", err.Error())
		return
	}
	if len(results) == 0 {
		resp.Diagnostics.AddError("Navigation not found", "No navigation entry found matching the specified criteria.")
		return
	}
	if len(results) > 1 {
		resp.Diagnostics.AddError(
			"Multiple Navigation entries found",
			fmt.Sprintf("Found %d navigation entries matching the specified criteria; expected exactly one. "+
				"Narrow the filter, e.g. with id.", len(results)),
		)
		return
	}

	result := results[0]

	if id, ok := GetInt64(result, "id"); ok {
		config.ID = types.Int64Value(id)
	}
	if v, ok := GetString(result, "name"); ok && v != "" {
		config.Name = types.StringValue(v)
	} else {
		config.Name = types.StringNull()
	}
	if v, ok := GetString(result, "label"); ok && v != "" {
		config.Label = types.StringValue(v)
	} else {
		config.Label = types.StringNull()
	}
	if v, ok := GetString(result, "url"); ok && v != "" {
		config.URL = types.StringValue(v)
	} else {
		config.URL = types.StringNull()
	}
	if v, ok := GetString(result, "icon"); ok && v != "" {
		config.Icon = types.StringValue(v)
	} else {
		config.Icon = types.StringNull()
	}
	if permRaw, ok := result["permission"]; ok && permRaw != nil {
		if permSlice, ok := permRaw.([]any); ok {
			values := make([]string, 0, len(permSlice))
			for _, v := range permSlice {
				if s, ok := v.(string); ok {
					values = append(values, s)
				}
			}
			if len(values) > 0 {
				permList, d := types.ListValueFrom(ctx, types.StringType, values)
				resp.Diagnostics.Append(d...)
				config.Permission = permList
			} else {
				config.Permission = types.ListNull(types.StringType)
			}
		} else {
			config.Permission = types.ListNull(types.StringType)
		}
	} else {
		config.Permission = types.ListNull(types.StringType)
	}
	if v, ok := GetString(result, "permission_operator"); ok && v != "" {
		config.PermissionOperator = types.StringValue(v)
	} else {
		config.PermissionOperator = types.StringNull()
	}
	if v, ok := GetInt64(result, "parent_id"); ok {
		config.ParentID = types.Int64Value(v)
	} else {
		config.ParentID = types.Int64Null()
	}
	if v, ok := GetBool(result, "is_active"); ok {
		config.IsActive = types.BoolValue(v)
	}
	if v, ok := GetInt64(result, "has_separator"); ok {
		config.HasSeparator = types.Int64Value(v)
	} else {
		config.HasSeparator = types.Int64Value(0)
	}
	if v, ok := GetInt64(result, "weight"); ok {
		config.Weight = types.Int64Value(v)
	} else {
		config.Weight = types.Int64Null()
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
