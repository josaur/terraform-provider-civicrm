package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &MessageTemplateDataSource{}
var _ datasource.DataSourceWithConfigure = &MessageTemplateDataSource{}

type MessageTemplateDataSource struct {
	client *Client
}

type MessageTemplateDataSourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	MsgTitle     types.String `tfsdk:"msg_title"`
	MsgSubject   types.String `tfsdk:"msg_subject"`
	MsgText      types.String `tfsdk:"msg_text"`
	MsgHTML      types.String `tfsdk:"msg_html"`
	IsActive     types.Bool   `tfsdk:"is_active"`
	WorkflowID   types.Int64  `tfsdk:"workflow_id"`
	WorkflowName types.String `tfsdk:"workflow_name"`
	IsDefault    types.Bool   `tfsdk:"is_default"`
	IsReserved   types.Bool   `tfsdk:"is_reserved"`
	IsSms        types.Bool   `tfsdk:"is_sms"`
	PdfFormatID  types.Int64  `tfsdk:"pdf_format_id"`
}

func NewMessageTemplateDataSource() datasource.DataSource {
	return &MessageTemplateDataSource{}
}

func (d *MessageTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_message_template"
}

func (d *MessageTemplateDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches a CiviCRM MessageTemplate by ID, or by workflow_name for CiviCRM's built-in workflow templates (e.g. the Invoice template).",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier. Specify either id, or workflow_name.",
				Optional:    true,
				Computed:    true,
			},
			"workflow_name": schema.StringAttribute{
				Description: "Message Template Workflow Name. As input, looks up a built-in workflow template by " +
					"its workflow_name (e.g. \"contribution_invoice_receipt\") instead of by id — combined with " +
					"is_reserved (default false) to resolve the editable copy of a built-in template without " +
					"hardcoding its numeric id, useful for `import` blocks whose id is derived at plan/apply time " +
					"rather than known in advance. Specify either id, or workflow_name.",
				Optional: true,
				Computed: true,
			},
			"msg_title":   schema.StringAttribute{Description: "Descriptive title of message.", Computed: true},
			"msg_subject": schema.StringAttribute{Description: "Subject for email message..", Computed: true},
			"msg_text":    schema.StringAttribute{Description: "Text formatted message.", Computed: true},
			"msg_html":    schema.StringAttribute{Description: "HTML formatted message.", Computed: true},
			"is_active":   schema.BoolAttribute{Description: "Is Active.", Computed: true},
			"workflow_id": schema.Int64Attribute{Description: "a pseudo-FK to civicrm_option_value.", Computed: true},
			"is_default":  schema.BoolAttribute{Description: "is this the default message template for the workflow referenced by workflow_id?.", Computed: true},
			"is_reserved": schema.BoolAttribute{
				Description: "is this the reserved message template which we ship for the workflow referenced by " +
					"workflow_id? As input when looking up by workflow_name, selects the reserved (true) or " +
					"non-reserved/editable (false, the default) row. Ignored when looking up by id.",
				Optional: true,
				Computed: true,
			},
			"is_sms":        schema.BoolAttribute{Description: "Is this message template used for sms?.", Computed: true},
			"pdf_format_id": schema.Int64Attribute{Description: "a pseudo-FK to civicrm_option_value containing PDF Page Format..", Computed: true},
		},
	}
}

func (d *MessageTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *MessageTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config MessageTemplateDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.ID.IsNull() && config.WorkflowName.IsNull() {
		resp.Diagnostics.AddError("Missing Filter", "One of 'id' or 'workflow_name' must be specified.")
		return
	}

	var where [][]any
	if !config.ID.IsNull() {
		where = [][]any{{"id", "=", config.ID.ValueInt64()}}
	} else {
		isReserved := false
		if !config.IsReserved.IsNull() {
			isReserved = config.IsReserved.ValueBool()
		}
		where = [][]any{
			{"workflow_name", "=", config.WorkflowName.ValueString()},
			{"is_reserved", "=", isReserved},
		}
	}

	tflog.Debug(ctx, "Reading MessageTemplate data source", map[string]any{"where": where})

	results, err := d.client.Get("MessageTemplate", where, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading MessageTemplate", err.Error())
		return
	}
	if len(results) == 0 {
		resp.Diagnostics.AddError("MessageTemplate not found", "No message_template found matching the specified criteria.")
		return
	}

	result := results[0]

	if id, ok := GetInt64(result, "id"); ok {
		config.ID = types.Int64Value(id)
	}
	if v, ok := GetString(result, "msg_title"); ok && v != "" {
		config.MsgTitle = types.StringValue(v)
	} else {
		config.MsgTitle = types.StringNull()
	}
	if v, ok := GetString(result, "msg_subject"); ok && v != "" {
		config.MsgSubject = types.StringValue(v)
	} else {
		config.MsgSubject = types.StringNull()
	}
	if v, ok := GetString(result, "msg_text"); ok && v != "" {
		config.MsgText = types.StringValue(v)
	} else {
		config.MsgText = types.StringNull()
	}
	if v, ok := GetString(result, "msg_html"); ok && v != "" {
		config.MsgHTML = types.StringValue(v)
	} else {
		config.MsgHTML = types.StringNull()
	}
	if v, ok := GetBool(result, "is_active"); ok {
		config.IsActive = types.BoolValue(v)
	}
	if v, ok := GetInt64(result, "workflow_id"); ok {
		config.WorkflowID = types.Int64Value(v)
	} else {
		config.WorkflowID = types.Int64Null()
	}
	if v, ok := GetString(result, "workflow_name"); ok && v != "" {
		config.WorkflowName = types.StringValue(v)
	} else {
		config.WorkflowName = types.StringNull()
	}
	if v, ok := GetBool(result, "is_default"); ok {
		config.IsDefault = types.BoolValue(v)
	}
	if v, ok := GetBool(result, "is_reserved"); ok {
		config.IsReserved = types.BoolValue(v)
	}
	if v, ok := GetBool(result, "is_sms"); ok {
		config.IsSms = types.BoolValue(v)
	}
	if v, ok := GetInt64(result, "pdf_format_id"); ok {
		config.PdfFormatID = types.Int64Value(v)
	} else {
		config.PdfFormatID = types.Int64Null()
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
