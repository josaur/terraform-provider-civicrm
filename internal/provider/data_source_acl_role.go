package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &ACLRoleDataSource{}
var _ datasource.DataSourceWithConfigure = &ACLRoleDataSource{}

type ACLRoleDataSource struct {
	client *Client
}

type ACLRoleDataSourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Label       types.String `tfsdk:"label"`
	Description types.String `tfsdk:"description"`
	IsActive    types.Bool   `tfsdk:"is_active"`
	Weight      types.Int64  `tfsdk:"weight"`
	Value       types.String `tfsdk:"value"`
}

func NewACLRoleDataSource() datasource.DataSource {
	return &ACLRoleDataSource{}
}

func (d *ACLRoleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acl_role"
}

func (d *ACLRoleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches a CiviCRM ACL Role by ID or name.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier of the ACL role. Specify either id or name.",
				Optional:    true,
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The machine name of the ACL role. Specify either id or name.",
				Optional:    true,
				Computed:    true,
			},
			"label": schema.StringAttribute{
				Description: "The display label of the ACL role.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the ACL role.",
				Computed:    true,
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether the ACL role is active.",
				Computed:    true,
			},
			"weight": schema.Int64Attribute{
				Description: "The sort weight of the ACL role.",
				Computed:    true,
			},
			"value": schema.StringAttribute{
				Description: "The value of the ACL role (used internally by CiviCRM).",
				Computed:    true,
			},
		},
	}
}

func (d *ACLRoleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ACLRoleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ACLRoleDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build where clause based on provided filters
	// ACL Roles are stored as OptionValues in the acl_role option group
	where := [][]any{
		{"option_group_id:name", "=", "acl_role"},
	}
	if !config.ID.IsNull() {
		where = append(where, []any{"id", "=", config.ID.ValueInt64()})
	}
	if !config.Name.IsNull() {
		where = append(where, []any{"name", "=", config.Name.ValueString()})
	}

	if config.ID.IsNull() && config.Name.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Filter",
			"At least one of 'id' or 'name' must be specified.",
		)
		return
	}

	tflog.Debug(ctx, "Reading ACL role data source", map[string]any{
		"filters": where,
	})

	results, err := d.client.Get("OptionValue", where, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading ACL role",
			"Could not read ACL role: "+err.Error(),
		)
		return
	}

	if len(results) == 0 {
		resp.Diagnostics.AddError(
			"ACL role not found",
			"No ACL role found matching the specified criteria.",
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

	if label, ok := GetString(result, "label"); ok {
		config.Label = types.StringValue(label)
	}

	if desc, ok := GetString(result, "description"); ok && desc != "" {
		config.Description = types.StringValue(desc)
	} else {
		config.Description = types.StringNull()
	}

	if active, ok := GetBool(result, "is_active"); ok {
		config.IsActive = types.BoolValue(active)
	}

	if weight, ok := GetInt64(result, "weight"); ok {
		config.Weight = types.Int64Value(weight)
	}

	if value, ok := GetString(result, "value"); ok {
		config.Value = types.StringValue(value)
	}

	diags = resp.State.Set(ctx, config)
	resp.Diagnostics.Append(diags...)
}
