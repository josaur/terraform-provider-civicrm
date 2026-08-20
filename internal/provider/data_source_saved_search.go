package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &SavedSearchDataSource{}
var _ datasource.DataSourceWithConfigure = &SavedSearchDataSource{}

type SavedSearchDataSource struct {
	client *Client
}

type SavedSearchDataSourceModel struct {
	ID             types.Int64  `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Label          types.String `tfsdk:"label"`
	FormValues     types.String `tfsdk:"form_values"`
	MappingID      types.Int64  `tfsdk:"mapping_id"`
	SearchCustomID types.Int64  `tfsdk:"search_custom_id"`
	APIEntity      types.String `tfsdk:"api_entity"`
	APIParams      types.String `tfsdk:"api_params"`
	CreatedID      types.Int64  `tfsdk:"created_id"`
	ModifiedID     types.Int64  `tfsdk:"modified_id"`
	ExpiresDate    types.String `tfsdk:"expires_date"`
	CreatedDate    types.String `tfsdk:"created_date"`
	ModifiedDate   types.String `tfsdk:"modified_date"`
	Description    types.String `tfsdk:"description"`
	IsTemplate     types.Bool   `tfsdk:"is_template"`
}

func NewSavedSearchDataSource() datasource.DataSource {
	return &SavedSearchDataSource{}
}

func (d *SavedSearchDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_saved_search"
}

func (d *SavedSearchDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches a CiviCRM SavedSearch by ID or name.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier. Specify either id or name.",
				Optional:    true,
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the saved_search. Specify either id or name.",
				Optional:    true,
				Computed:    true,
			},
			"label":            schema.StringAttribute{Description: "Administrative label for search.", Computed: true},
			"form_values":      schema.StringAttribute{Description: "Submitted form values for this search.", Computed: true},
			"mapping_id":       schema.Int64Attribute{Description: "Foreign key to civicrm_mapping used for saved search-builder searches..", Computed: true},
			"search_custom_id": schema.Int64Attribute{Description: "Foreign key to civicrm_option value table used for saved custom searches..", Computed: true},
			"api_entity":       schema.StringAttribute{Description: "Entity name for API based search.", Computed: true},
			"api_params":       schema.StringAttribute{Description: "Parameters for API based search.", Computed: true},
			"created_id":       schema.Int64Attribute{Description: "FK to contact table..", Computed: true},
			"modified_id":      schema.Int64Attribute{Description: "FK to contact table..", Computed: true},
			"expires_date":     schema.StringAttribute{Description: "Optional date after which the search is not needed.", Computed: true},
			"created_date":     schema.StringAttribute{Description: "When the search was created..", Computed: true},
			"modified_date":    schema.StringAttribute{Description: "When the search was last modified..", Computed: true},
			"description":      schema.StringAttribute{Description: "Saved Search Description.", Computed: true},
			"is_template":      schema.BoolAttribute{Description: "Search templates are used as a starting point for building new searches.", Computed: true},
		},
	}
}

func (d *SavedSearchDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *SavedSearchDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SavedSearchDataSourceModel
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

	tflog.Debug(ctx, "Reading SavedSearch data source", map[string]any{"where": where})

	results, err := d.client.Get("SavedSearch", where, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading SavedSearch", err.Error())
		return
	}
	if len(results) == 0 {
		resp.Diagnostics.AddError("SavedSearch not found", "No saved_search found matching the specified criteria.")
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
	if raw, ok := result["form_values"]; ok && raw != nil {
		encoded, err := encodeJSONAttribute(raw)
		if err == nil && encoded != "" && encoded != "null" {
			config.FormValues = types.StringValue(encoded)
		} else {
			config.FormValues = types.StringNull()
		}
	} else {
		config.FormValues = types.StringNull()
	}
	if v, ok := GetInt64(result, "mapping_id"); ok {
		config.MappingID = types.Int64Value(v)
	} else {
		config.MappingID = types.Int64Null()
	}
	if v, ok := GetInt64(result, "search_custom_id"); ok {
		config.SearchCustomID = types.Int64Value(v)
	} else {
		config.SearchCustomID = types.Int64Null()
	}
	if v, ok := GetString(result, "api_entity"); ok && v != "" {
		config.APIEntity = types.StringValue(v)
	} else {
		config.APIEntity = types.StringNull()
	}
	if raw, ok := result["api_params"]; ok && raw != nil {
		encoded, err := encodeJSONAttribute(raw)
		if err == nil && encoded != "" && encoded != "null" {
			config.APIParams = types.StringValue(encoded)
		} else {
			config.APIParams = types.StringNull()
		}
	} else {
		config.APIParams = types.StringNull()
	}
	if v, ok := GetInt64(result, "created_id"); ok {
		config.CreatedID = types.Int64Value(v)
	} else {
		config.CreatedID = types.Int64Null()
	}
	if v, ok := GetInt64(result, "modified_id"); ok {
		config.ModifiedID = types.Int64Value(v)
	} else {
		config.ModifiedID = types.Int64Null()
	}
	if v, ok := GetString(result, "expires_date"); ok && v != "" {
		config.ExpiresDate = types.StringValue(v)
	} else {
		config.ExpiresDate = types.StringNull()
	}
	if v, ok := GetString(result, "created_date"); ok {
		config.CreatedDate = types.StringValue(v)
	}
	if v, ok := GetString(result, "modified_date"); ok {
		config.ModifiedDate = types.StringValue(v)
	}
	if v, ok := GetString(result, "description"); ok && v != "" {
		config.Description = types.StringValue(v)
	} else {
		config.Description = types.StringNull()
	}
	if v, ok := GetBool(result, "is_template"); ok {
		config.IsTemplate = types.BoolValue(v)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
