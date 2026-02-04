package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &GroupDataSource{}
var _ datasource.DataSourceWithConfigure = &GroupDataSource{}

type GroupDataSource struct {
	client *Client
}

type GroupDataSourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Title       types.String `tfsdk:"title"`
	Description types.String `tfsdk:"description"`
	IsActive    types.Bool   `tfsdk:"is_active"`
	Visibility  types.String `tfsdk:"visibility"`
}

func NewGroupDataSource() datasource.DataSource {
	return &GroupDataSource{}
}

func (d *GroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

func (d *GroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches a CiviCRM Group by ID or name.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier of the group. Specify either id or name.",
				Optional:    true,
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The machine name of the group. Specify either id or name.",
				Optional:    true,
				Computed:    true,
			},
			"title": schema.StringAttribute{
				Description: "The display title of the group.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the group.",
				Computed:    true,
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether the group is active.",
				Computed:    true,
			},
			"visibility": schema.StringAttribute{
				Description: "The visibility of the group.",
				Computed:    true,
			},
		},
	}
}

func (d *GroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}

func (d *GroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config GroupDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build where clause based on provided filters
	var where [][]any
	if !config.ID.IsNull() {
		where = append(where, []any{"id", "=", config.ID.ValueInt64()})
	}
	if !config.Name.IsNull() {
		where = append(where, []any{"name", "=", config.Name.ValueString()})
	}

	if len(where) == 0 {
		resp.Diagnostics.AddError(
			"Missing Filter",
			"At least one of 'id' or 'name' must be specified.",
		)
		return
	}

	tflog.Debug(ctx, "Reading group data source", map[string]any{
		"filters": where,
	})

	results, err := d.client.Get("Group", where, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading group",
			"Could not read group: "+err.Error(),
		)
		return
	}

	if len(results) == 0 {
		resp.Diagnostics.AddError(
			"Group not found",
			"No group found matching the specified criteria.",
		)
		return
	}

	result := results[0]

	// Update state
	if id, ok := GetInt64(result, "id"); ok {
		config.ID = types.Int64Value(id)
	}

	if name, ok := GetString(result, "name"); ok {
		config.Name = types.StringValue(name)
	}

	if title, ok := GetString(result, "title"); ok {
		config.Title = types.StringValue(title)
	}

	if desc, ok := GetString(result, "description"); ok && desc != "" {
		config.Description = types.StringValue(desc)
	} else {
		config.Description = types.StringNull()
	}

	if active, ok := GetBool(result, "is_active"); ok {
		config.IsActive = types.BoolValue(active)
	}

	if visibility, ok := GetString(result, "visibility"); ok {
		config.Visibility = types.StringValue(visibility)
	}

	diags = resp.State.Set(ctx, config)
	resp.Diagnostics.Append(diags...)
}
