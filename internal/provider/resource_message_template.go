package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &MessageTemplateResource{}
	_ resource.ResourceWithConfigure   = &MessageTemplateResource{}
	_ resource.ResourceWithImportState = &MessageTemplateResource{}
)

// MessageTemplateResource manages CiviCRM MessageTemplates. MessageTemplates
// hold the HTML/text bodies used by CiviMail, transactional mail
// (Email.send), and workflow notifications. Use for both user-defined
// templates and for overriding the default workflow templates.
type MessageTemplateResource struct {
	client *Client
}

type MessageTemplateResourceModel struct {
	ID         types.Int64  `tfsdk:"id"`
	MsgTitle   types.String `tfsdk:"msg_title"`
	MsgSubject types.String `tfsdk:"msg_subject"`
	MsgHTML    types.String `tfsdk:"msg_html"`
	MsgText    types.String `tfsdk:"msg_text"`
	IsActive   types.Bool   `tfsdk:"is_active"`
	IsReserved types.Bool   `tfsdk:"is_reserved"`
	IsDefault  types.Bool   `tfsdk:"is_default"`
	WorkflowName types.String `tfsdk:"workflow_name"`
}

func NewMessageTemplateResource() resource.Resource {
	return &MessageTemplateResource{}
}

func (r *MessageTemplateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_message_template"
}

func (r *MessageTemplateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a CiviCRM MessageTemplate. MessageTemplates hold the HTML/text bodies used by CiviMail, transactional mail (Email.send), and workflow notifications.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Unique ID of the message template.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"msg_title": schema.StringAttribute{
				Description: "Internal title of the template (used to look it up by name).",
				Required:    true,
			},
			"msg_subject": schema.StringAttribute{
				Description: "Subject line for outgoing mails.",
				Optional:    true,
			},
			"msg_html": schema.StringAttribute{
				Description: "HTML body of the message.",
				Optional:    true,
			},
			"msg_text": schema.StringAttribute{
				Description: "Plain-text body of the message.",
				Optional:    true,
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether the template is active. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"is_reserved": schema.BoolAttribute{
				Description: "Whether the template is reserved (system-managed). Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"is_default": schema.BoolAttribute{
				Description: "For workflow templates: whether this is the default version. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"workflow_name": schema.StringAttribute{
				Description: "For workflow templates: the machine name of the workflow this template overrides (e.g. `contribution_online_receipt`). Leave empty for user-defined templates.",
				Optional:    true,
			},
		},
	}
}

func (r *MessageTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *Client, got: %T.", req.ProviderData),
		)
		return
	}
	r.client = client
}

func (r *MessageTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan MessageTemplateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating MessageTemplate", map[string]any{"msg_title": plan.MsgTitle.ValueString()})

	values := map[string]any{
		"msg_title":   plan.MsgTitle.ValueString(),
		"is_active":   plan.IsActive.ValueBool(),
		"is_reserved": plan.IsReserved.ValueBool(),
		"is_default":  plan.IsDefault.ValueBool(),
	}
	if !plan.MsgSubject.IsNull() && !plan.MsgSubject.IsUnknown() {
		values["msg_subject"] = plan.MsgSubject.ValueString()
	}
	if !plan.MsgHTML.IsNull() && !plan.MsgHTML.IsUnknown() {
		values["msg_html"] = plan.MsgHTML.ValueString()
	}
	if !plan.MsgText.IsNull() && !plan.MsgText.IsUnknown() {
		values["msg_text"] = plan.MsgText.ValueString()
	}
	if !plan.WorkflowName.IsNull() && !plan.WorkflowName.IsUnknown() && plan.WorkflowName.ValueString() != "" {
		values["workflow_name"] = plan.WorkflowName.ValueString()
	}

	result, err := r.client.Create("MessageTemplate", values)
	if err != nil {
		resp.Diagnostics.AddError("Error creating MessageTemplate", err.Error())
		return
	}

	if createdID, ok := GetInt64(result, "id"); ok {
		if fullResult, err2 := r.client.GetByID("MessageTemplate", createdID, nil); err2 == nil {
			result = fullResult
		}
	}
	r.mapResultToState(result, &plan)
	tflog.Debug(ctx, "Created MessageTemplate", map[string]any{"id": plan.ID.ValueInt64()})
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *MessageTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state MessageTemplateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.GetByID("MessageTemplate", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading MessageTemplate",
			"Could not read MessageTemplate ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	r.mapResultToState(result, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *MessageTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan MessageTemplateResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var state MessageTemplateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	values := map[string]any{
		"msg_title":   plan.MsgTitle.ValueString(),
		"is_active":   plan.IsActive.ValueBool(),
		"is_reserved": plan.IsReserved.ValueBool(),
		"is_default":  plan.IsDefault.ValueBool(),
	}
	if !plan.MsgSubject.IsNull() && !plan.MsgSubject.IsUnknown() {
		values["msg_subject"] = plan.MsgSubject.ValueString()
	} else {
		values["msg_subject"] = nil
	}
	if !plan.MsgHTML.IsNull() && !plan.MsgHTML.IsUnknown() {
		values["msg_html"] = plan.MsgHTML.ValueString()
	} else {
		values["msg_html"] = nil
	}
	if !plan.MsgText.IsNull() && !plan.MsgText.IsUnknown() {
		values["msg_text"] = plan.MsgText.ValueString()
	} else {
		values["msg_text"] = nil
	}
	if !plan.WorkflowName.IsNull() && !plan.WorkflowName.IsUnknown() && plan.WorkflowName.ValueString() != "" {
		values["workflow_name"] = plan.WorkflowName.ValueString()
	} else {
		values["workflow_name"] = nil
	}

	_, err := r.client.Update("MessageTemplate", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating MessageTemplate",
			"Could not update MessageTemplate ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	plan.ID = state.ID

	result, err := r.client.GetByID("MessageTemplate", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading MessageTemplate after update",
			"Could not re-read MessageTemplate ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}
	r.mapResultToState(result, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *MessageTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state MessageTemplateResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.Delete("MessageTemplate", state.ID.ValueInt64()); err != nil {
		resp.Diagnostics.AddError(
			"Error deleting MessageTemplate",
			"Could not delete MessageTemplate ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
	}
}

func (r *MessageTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", "Could not parse import ID as integer: "+err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (r *MessageTemplateResource) mapResultToState(result map[string]any, model *MessageTemplateResourceModel) {
	if id, ok := GetInt64(result, "id"); ok {
		model.ID = types.Int64Value(id)
	}
	if title, ok := GetString(result, "msg_title"); ok {
		model.MsgTitle = types.StringValue(title)
	}
	if subj, ok := GetString(result, "msg_subject"); ok && subj != "" {
		model.MsgSubject = types.StringValue(subj)
	} else {
		model.MsgSubject = types.StringNull()
	}
	if html, ok := GetString(result, "msg_html"); ok && html != "" {
		model.MsgHTML = types.StringValue(html)
	} else {
		model.MsgHTML = types.StringNull()
	}
	if text, ok := GetString(result, "msg_text"); ok && text != "" {
		model.MsgText = types.StringValue(text)
	} else {
		model.MsgText = types.StringNull()
	}
	if active, ok := GetBool(result, "is_active"); ok {
		model.IsActive = types.BoolValue(active)
	}
	if reserved, ok := GetBool(result, "is_reserved"); ok {
		model.IsReserved = types.BoolValue(reserved)
	}
	if def, ok := GetBool(result, "is_default"); ok {
		model.IsDefault = types.BoolValue(def)
	}
	if wf, ok := GetString(result, "workflow_name"); ok && wf != "" {
		model.WorkflowName = types.StringValue(wf)
	} else {
		model.WorkflowName = types.StringNull()
	}
}
