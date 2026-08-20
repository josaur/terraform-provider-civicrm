package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &SearchDisplayDataSource{}
var _ datasource.DataSourceWithConfigure = &SearchDisplayDataSource{}

type SearchDisplayDataSource struct {
	client *Client
}

type SearchDisplayDataSourceModel struct {
	ID                    types.Int64  `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	Label                 types.String `tfsdk:"label"`
	SavedSearchID         types.Int64  `tfsdk:"saved_search_id"`
	Type                  types.String `tfsdk:"type"`
	Settings              types.String `tfsdk:"settings"`
	ACLBypass             types.Bool   `tfsdk:"acl_bypass"`
	IsAutocompleteDefault types.Bool   `tfsdk:"is_autocomplete_default"`
}

func NewSearchDisplayDataSource() datasource.DataSource {
	return &SearchDisplayDataSource{}
}

func (d *SearchDisplayDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_search_display"
}

func (d *SearchDisplayDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches a CiviCRM SearchKit SearchDisplay by ID, or by name (and/or " +
			"saved_search_id + type to disambiguate). Lets existing displays be adopted via the " +
			"`import` block pattern instead of being recreated as duplicates.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier. Specify id, or name (optionally combined with " +
					"saved_search_id and/or type to disambiguate).",
				Optional: true,
				Computed: true,
			},
			"name": schema.StringAttribute{
				Description: "Machine name of the search display. Specify id, or name.",
				Optional:    true,
				Computed:    true,
			},
			"saved_search_id": schema.Int64Attribute{
				Description: "ID of the saved search this display renders. Optional filter, combine " +
					"with name and/or type to disambiguate when name alone is not unique.",
				Optional: true,
				Computed: true,
			},
			"type": schema.StringAttribute{
				Description: "Display type filter (table, list, grid, tree, autocomplete, entity, " +
					"batch). Optional, combine with name and/or saved_search_id to disambiguate.",
				Optional: true,
				Computed: true,
			},
			"label":                   schema.StringAttribute{Description: "Administrative label for the display.", Computed: true},
			"settings":                schema.StringAttribute{Description: "Display settings as a JSON string.", Computed: true},
			"acl_bypass":              schema.BoolAttribute{Description: "Whether this display bypasses ACL permission checks.", Computed: true},
			"is_autocomplete_default": schema.BoolAttribute{Description: "Whether this is the default autocomplete display for its saved search's entity.", Computed: true},
		},
	}
}

func (d *SearchDisplayDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *SearchDisplayDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SearchDisplayDataSourceModel
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
	if !config.SavedSearchID.IsNull() {
		where = append(where, []any{"saved_search_id", "=", config.SavedSearchID.ValueInt64()})
	}
	if !config.Type.IsNull() {
		where = append(where, []any{"type", "=", config.Type.ValueString()})
	}

	tflog.Debug(ctx, "Reading SearchDisplay data source", map[string]any{"where": where})

	results, err := d.client.Get("SearchDisplay", where, searchDisplaySelect)
	if err != nil {
		resp.Diagnostics.AddError("Error reading SearchDisplay", err.Error())
		return
	}
	if len(results) == 0 {
		resp.Diagnostics.AddError("SearchDisplay not found", "No search_display found matching the specified criteria.")
		return
	}
	if len(results) > 1 {
		resp.Diagnostics.AddError(
			"Multiple SearchDisplays found",
			fmt.Sprintf("Found %d search displays matching the specified criteria; expected exactly one. "+
				"Narrow the filter with saved_search_id and/or type.", len(results)),
		)
		return
	}

	result := results[0]

	if id, ok := GetInt64(result, "id"); ok {
		config.ID = types.Int64Value(id)
	}
	if v, ok := GetString(result, "name"); ok {
		config.Name = types.StringValue(v)
	}
	if v, ok := GetString(result, "label"); ok && v != "" {
		config.Label = types.StringValue(v)
	} else {
		config.Label = types.StringNull()
	}
	if v, ok := GetInt64(result, "saved_search_id"); ok {
		config.SavedSearchID = types.Int64Value(v)
	} else {
		config.SavedSearchID = types.Int64Null()
	}
	if v, ok := GetString(result, "type"); ok && v != "" {
		config.Type = types.StringValue(v)
	} else {
		config.Type = types.StringNull()
	}
	if raw, ok := result["settings"]; ok && raw != nil {
		encoded, err := encodeJSONAttribute(raw)
		if err == nil && encoded != "" && encoded != "null" {
			config.Settings = types.StringValue(encoded)
		} else {
			config.Settings = types.StringNull()
		}
	} else {
		config.Settings = types.StringNull()
	}
	if v, ok := GetBool(result, "acl_bypass"); ok {
		config.ACLBypass = types.BoolValue(v)
	}
	if v, ok := GetBool(result, "is_autocomplete_default"); ok {
		config.IsAutocompleteDefault = types.BoolValue(v)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
