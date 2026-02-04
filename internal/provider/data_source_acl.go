package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &ACLDataSource{}
var _ datasource.DataSourceWithConfigure = &ACLDataSource{}

type ACLDataSource struct {
	client *Client
}

type ACLDataSourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	EntityTable types.String `tfsdk:"entity_table"`
	EntityID    types.Int64  `tfsdk:"entity_id"`
	Operation   types.String `tfsdk:"operation"`
	ObjectTable types.String `tfsdk:"object_table"`
	ObjectID    types.Int64  `tfsdk:"object_id"`
	IsActive    types.Bool   `tfsdk:"is_active"`
	Deny        types.Bool   `tfsdk:"deny"`
	Priority    types.Int64  `tfsdk:"priority"`
}

func NewACLDataSource() datasource.DataSource {
	return &ACLDataSource{}
}

func (d *ACLDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acl"
}

func (d *ACLDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches a CiviCRM ACL rule by ID or name.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier of the ACL. Specify either id or name.",
				Optional:    true,
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the ACL rule. Specify either id or name.",
				Optional:    true,
				Computed:    true,
			},
			"entity_table": schema.StringAttribute{
				Description: "The entity table that owns this ACL.",
				Computed:    true,
			},
			"entity_id": schema.Int64Attribute{
				Description: "The ID of the ACL role this rule belongs to.",
				Computed:    true,
			},
			"operation": schema.StringAttribute{
				Description: "The operation this ACL grants.",
				Computed:    true,
			},
			"object_table": schema.StringAttribute{
				Description: "The type of object being permissioned.",
				Computed:    true,
			},
			"object_id": schema.Int64Attribute{
				Description: "The ID of the specific object being permissioned.",
				Computed:    true,
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether the ACL rule is active.",
				Computed:    true,
			},
			"deny": schema.BoolAttribute{
				Description: "Whether this ACL denies rather than allows access.",
				Computed:    true,
			},
			"priority": schema.Int64Attribute{
				Description: "The priority of the ACL rule.",
				Computed:    true,
			},
		},
	}
}

func (d *ACLDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ACLDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ACLDataSourceModel
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

	tflog.Debug(ctx, "Reading ACL data source", map[string]any{
		"filters": where,
	})

	results, err := d.client.Get("ACL", where, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading ACL",
			"Could not read ACL: "+err.Error(),
		)
		return
	}

	if len(results) == 0 {
		resp.Diagnostics.AddError(
			"ACL not found",
			"No ACL found matching the specified criteria.",
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

	if entityTable, ok := GetString(result, "entity_table"); ok {
		config.EntityTable = types.StringValue(entityTable)
	}

	if entityID, ok := GetInt64(result, "entity_id"); ok {
		config.EntityID = types.Int64Value(entityID)
	}

	if operation, ok := GetString(result, "operation"); ok {
		config.Operation = types.StringValue(operation)
	}

	if objectTable, ok := GetString(result, "object_table"); ok {
		config.ObjectTable = types.StringValue(objectTable)
	}

	if objectID, ok := GetInt64(result, "object_id"); ok {
		config.ObjectID = types.Int64Value(objectID)
	} else {
		config.ObjectID = types.Int64Null()
	}

	if active, ok := GetBool(result, "is_active"); ok {
		config.IsActive = types.BoolValue(active)
	}

	if deny, ok := GetBool(result, "deny"); ok {
		config.Deny = types.BoolValue(deny)
	}

	if priority, ok := GetInt64(result, "priority"); ok {
		config.Priority = types.Int64Value(priority)
	}

	diags = resp.State.Set(ctx, config)
	resp.Diagnostics.Append(diags...)
}
