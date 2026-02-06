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
	_ resource.Resource                = &MailSettingsResource{}
	_ resource.ResourceWithConfigure   = &MailSettingsResource{}
	_ resource.ResourceWithImportState = &MailSettingsResource{}
)

// MailSettingsResource manages mail settings in CiviCRM.
type MailSettingsResource struct {
	client *Client
}

type MailSettingsResourceModel struct {
	ID                                 types.Int64  `tfsdk:"id"`
	DomainID                           types.Int64  `tfsdk:"domain_id"`
	Name                               types.String `tfsdk:"name"`
	IsDefault                          types.Bool   `tfsdk:"is_default"`
	Domain                             types.String `tfsdk:"domain"`
	Localpart                          types.String `tfsdk:"localpart"`
	ReturnPath                         types.String `tfsdk:"return_path"`
	Protocol                           types.String `tfsdk:"protocol"`
	Server                             types.String `tfsdk:"server"`
	Port                               types.Int64  `tfsdk:"port"`
	Username                           types.String `tfsdk:"username"`
	Password                           types.String `tfsdk:"password"`
	IsSSL                              types.Bool   `tfsdk:"is_ssl"`
	Source                             types.String `tfsdk:"source"`
	ActivityStatus                     types.String `tfsdk:"activity_status"`
	IsNonCaseEmailSkipped              types.Bool   `tfsdk:"is_non_case_email_skipped"`
	IsContactCreationDisabledIfNoMatch types.Bool   `tfsdk:"is_contact_creation_disabled_if_no_match"`
	IsActive                           types.Bool   `tfsdk:"is_active"`
	ActivityTypeID                     types.Int64  `tfsdk:"activity_type_id"`
	CampaignID                         types.Int64  `tfsdk:"campaign_id"`
	ActivitySource                     types.String `tfsdk:"activity_source"`
	ActivityTargets                    types.String `tfsdk:"activity_targets"`
	ActivityAssignees                  types.String `tfsdk:"activity_assignees"`
}

func NewMailSettingsResource() resource.Resource {
	return &MailSettingsResource{}
}

func (r *MailSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mail_settings"
}

func (r *MailSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages CiviCRM Mail Settings for inbound email processing.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier of the mail settings.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"domain_id": schema.Int64Attribute{
				Description: "The domain ID this mail setting belongs to.",
				Optional:    true,
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of this mail setting configuration.",
				Required:    true,
			},
			"is_default": schema.BoolAttribute{
				Description: "Whether this is the default mail setting. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"domain": schema.StringAttribute{
				Description: "The email domain (e.g., 'example.org').",
				Optional:    true,
			},
			"localpart": schema.StringAttribute{
				Description: "The local part prefix for bounce processing.",
				Optional:    true,
			},
			"return_path": schema.StringAttribute{
				Description: "The return path email address.",
				Optional:    true,
			},
			"protocol": schema.StringAttribute{
				Description: "The mail protocol (e.g., 'IMAP', 'POP3', 'Maildir', 'Localdir').",
				Optional:    true,
			},
			"server": schema.StringAttribute{
				Description: "The mail server hostname.",
				Optional:    true,
			},
			"port": schema.Int64Attribute{
				Description: "The mail server port.",
				Optional:    true,
			},
			"username": schema.StringAttribute{
				Description: "The username for mail server authentication.",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "The password for mail server authentication.",
				Optional:    true,
				Sensitive:   true,
			},
			"is_ssl": schema.BoolAttribute{
				Description: "Whether to use SSL/TLS for the connection. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"source": schema.StringAttribute{
				Description: "The mail source (folder path for Maildir/Localdir).",
				Optional:    true,
			},
			"activity_status": schema.StringAttribute{
				Description: "The default activity status for email activities.",
				Optional:    true,
			},
			"is_non_case_email_skipped": schema.BoolAttribute{
				Description: "Whether to skip emails not associated with a case. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"is_contact_creation_disabled_if_no_match": schema.BoolAttribute{
				Description: "Whether to disable contact creation if no match is found. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether this mail setting is active. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"activity_type_id": schema.Int64Attribute{
				Description: "The activity type ID for email activities.",
				Optional:    true,
			},
			"campaign_id": schema.Int64Attribute{
				Description: "The campaign ID to associate with email activities.",
				Optional:    true,
			},
			"activity_source": schema.StringAttribute{
				Description: "The activity source contact handling.",
				Optional:    true,
			},
			"activity_targets": schema.StringAttribute{
				Description: "The activity targets contact handling.",
				Optional:    true,
			},
			"activity_assignees": schema.StringAttribute{
				Description: "The activity assignees contact handling.",
				Optional:    true,
			},
		},
	}
}

func (r *MailSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MailSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan MailSettingsResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating mail settings", map[string]any{
		"name": plan.Name.ValueString(),
	})

	// Build values for API call
	values := map[string]any{
		"name":                      plan.Name.ValueString(),
		"is_default":                plan.IsDefault.ValueBool(),
		"is_ssl":                    plan.IsSSL.ValueBool(),
		"is_non_case_email_skipped": plan.IsNonCaseEmailSkipped.ValueBool(),
		"is_contact_creation_disabled_if_no_match": plan.IsContactCreationDisabledIfNoMatch.ValueBool(),
		"is_active": plan.IsActive.ValueBool(),
	}

	if !plan.DomainID.IsNull() {
		values["domain_id"] = plan.DomainID.ValueInt64()
	}

	if !plan.Domain.IsNull() {
		values["domain"] = plan.Domain.ValueString()
	}

	if !plan.Localpart.IsNull() {
		values["localpart"] = plan.Localpart.ValueString()
	}

	if !plan.ReturnPath.IsNull() {
		values["return_path"] = plan.ReturnPath.ValueString()
	}

	if !plan.Protocol.IsNull() {
		values["protocol"] = plan.Protocol.ValueString()
	}

	if !plan.Server.IsNull() {
		values["server"] = plan.Server.ValueString()
	}

	if !plan.Port.IsNull() {
		values["port"] = plan.Port.ValueInt64()
	}

	if !plan.Username.IsNull() {
		values["username"] = plan.Username.ValueString()
	}

	if !plan.Password.IsNull() {
		values["password"] = plan.Password.ValueString()
	}

	if !plan.Source.IsNull() {
		values["source"] = plan.Source.ValueString()
	}

	if !plan.ActivityStatus.IsNull() {
		values["activity_status"] = plan.ActivityStatus.ValueString()
	}

	if !plan.ActivityTypeID.IsNull() {
		values["activity_type_id"] = plan.ActivityTypeID.ValueInt64()
	}

	if !plan.CampaignID.IsNull() {
		values["campaign_id"] = plan.CampaignID.ValueInt64()
	}

	if !plan.ActivitySource.IsNull() {
		values["activity_source"] = plan.ActivitySource.ValueString()
	}

	if !plan.ActivityTargets.IsNull() {
		values["activity_targets"] = plan.ActivityTargets.ValueString()
	}

	if !plan.ActivityAssignees.IsNull() {
		values["activity_assignees"] = plan.ActivityAssignees.ValueString()
	}

	// Call API
	result, err := r.client.Create("MailSettings", values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating mail settings",
			"Could not create mail settings, unexpected error: "+err.Error(),
		)
		return
	}

	// Update state with response
	r.mapResponseToModel(result, &plan)

	tflog.Debug(ctx, "Created mail settings", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *MailSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state MailSettingsResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading mail settings", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	result, err := r.client.GetByID("MailSettings", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading mail settings",
			"Could not read mail settings ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Update state
	r.mapResponseToModel(result, &state)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *MailSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan MailSettingsResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state MailSettingsResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating mail settings", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	// Build values for API call
	values := map[string]any{
		"name":                      plan.Name.ValueString(),
		"is_default":                plan.IsDefault.ValueBool(),
		"is_ssl":                    plan.IsSSL.ValueBool(),
		"is_non_case_email_skipped": plan.IsNonCaseEmailSkipped.ValueBool(),
		"is_contact_creation_disabled_if_no_match": plan.IsContactCreationDisabledIfNoMatch.ValueBool(),
		"is_active": plan.IsActive.ValueBool(),
	}

	if !plan.DomainID.IsNull() {
		values["domain_id"] = plan.DomainID.ValueInt64()
	}

	if !plan.Domain.IsNull() {
		values["domain"] = plan.Domain.ValueString()
	} else {
		values["domain"] = nil
	}

	if !plan.Localpart.IsNull() {
		values["localpart"] = plan.Localpart.ValueString()
	} else {
		values["localpart"] = nil
	}

	if !plan.ReturnPath.IsNull() {
		values["return_path"] = plan.ReturnPath.ValueString()
	} else {
		values["return_path"] = nil
	}

	if !plan.Protocol.IsNull() {
		values["protocol"] = plan.Protocol.ValueString()
	} else {
		values["protocol"] = nil
	}

	if !plan.Server.IsNull() {
		values["server"] = plan.Server.ValueString()
	} else {
		values["server"] = nil
	}

	if !plan.Port.IsNull() {
		values["port"] = plan.Port.ValueInt64()
	} else {
		values["port"] = nil
	}

	if !plan.Username.IsNull() {
		values["username"] = plan.Username.ValueString()
	} else {
		values["username"] = nil
	}

	if !plan.Password.IsNull() {
		values["password"] = plan.Password.ValueString()
	} else {
		values["password"] = nil
	}

	if !plan.Source.IsNull() {
		values["source"] = plan.Source.ValueString()
	} else {
		values["source"] = nil
	}

	if !plan.ActivityStatus.IsNull() {
		values["activity_status"] = plan.ActivityStatus.ValueString()
	} else {
		values["activity_status"] = nil
	}

	if !plan.ActivityTypeID.IsNull() {
		values["activity_type_id"] = plan.ActivityTypeID.ValueInt64()
	} else {
		values["activity_type_id"] = nil
	}

	if !plan.CampaignID.IsNull() {
		values["campaign_id"] = plan.CampaignID.ValueInt64()
	} else {
		values["campaign_id"] = nil
	}

	if !plan.ActivitySource.IsNull() {
		values["activity_source"] = plan.ActivitySource.ValueString()
	} else {
		values["activity_source"] = nil
	}

	if !plan.ActivityTargets.IsNull() {
		values["activity_targets"] = plan.ActivityTargets.ValueString()
	} else {
		values["activity_targets"] = nil
	}

	if !plan.ActivityAssignees.IsNull() {
		values["activity_assignees"] = plan.ActivityAssignees.ValueString()
	} else {
		values["activity_assignees"] = nil
	}

	// Call API
	result, err := r.client.Update("MailSettings", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating mail settings",
			"Could not update mail settings ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Update state
	plan.ID = state.ID
	r.mapResponseToModel(result, &plan)

	tflog.Debug(ctx, "Updated mail settings", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *MailSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state MailSettingsResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting mail settings", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	err := r.client.Delete("MailSettings", state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting mail settings",
			"Could not delete mail settings ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "Deleted mail settings", map[string]any{
		"id": state.ID.ValueInt64(),
	})
}

func (r *MailSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

// mapResponseToModel maps API response to the model
func (r *MailSettingsResource) mapResponseToModel(result map[string]any, model *MailSettingsResourceModel) {
	if id, ok := GetInt64(result, "id"); ok {
		model.ID = types.Int64Value(id)
	}

	if domainID, ok := GetInt64(result, "domain_id"); ok {
		model.DomainID = types.Int64Value(domainID)
	}

	if name, ok := GetString(result, "name"); ok {
		model.Name = types.StringValue(name)
	}

	if isDefault, ok := GetBool(result, "is_default"); ok {
		model.IsDefault = types.BoolValue(isDefault)
	}

	if domain, ok := GetString(result, "domain"); ok && domain != "" {
		model.Domain = types.StringValue(domain)
	} else {
		model.Domain = types.StringNull()
	}

	if localpart, ok := GetString(result, "localpart"); ok && localpart != "" {
		model.Localpart = types.StringValue(localpart)
	} else {
		model.Localpart = types.StringNull()
	}

	if returnPath, ok := GetString(result, "return_path"); ok && returnPath != "" {
		model.ReturnPath = types.StringValue(returnPath)
	} else {
		model.ReturnPath = types.StringNull()
	}

	if protocol, ok := GetString(result, "protocol"); ok && protocol != "" {
		model.Protocol = types.StringValue(protocol)
	} else {
		model.Protocol = types.StringNull()
	}

	if server, ok := GetString(result, "server"); ok && server != "" {
		model.Server = types.StringValue(server)
	} else {
		model.Server = types.StringNull()
	}

	if port, ok := GetInt64(result, "port"); ok {
		model.Port = types.Int64Value(port)
	} else {
		model.Port = types.Int64Null()
	}

	if username, ok := GetString(result, "username"); ok && username != "" {
		model.Username = types.StringValue(username)
	} else {
		model.Username = types.StringNull()
	}

	// Don't read password back from API for security reasons
	// Keep the planned value

	if isSSL, ok := GetBool(result, "is_ssl"); ok {
		model.IsSSL = types.BoolValue(isSSL)
	}

	if source, ok := GetString(result, "source"); ok && source != "" {
		model.Source = types.StringValue(source)
	} else {
		model.Source = types.StringNull()
	}

	if activityStatus, ok := GetString(result, "activity_status"); ok && activityStatus != "" {
		model.ActivityStatus = types.StringValue(activityStatus)
	} else {
		model.ActivityStatus = types.StringNull()
	}

	if isNonCaseEmailSkipped, ok := GetBool(result, "is_non_case_email_skipped"); ok {
		model.IsNonCaseEmailSkipped = types.BoolValue(isNonCaseEmailSkipped)
	}

	if isContactCreationDisabled, ok := GetBool(result, "is_contact_creation_disabled_if_no_match"); ok {
		model.IsContactCreationDisabledIfNoMatch = types.BoolValue(isContactCreationDisabled)
	}

	if isActive, ok := GetBool(result, "is_active"); ok {
		model.IsActive = types.BoolValue(isActive)
	}

	if activityTypeID, ok := GetInt64(result, "activity_type_id"); ok {
		model.ActivityTypeID = types.Int64Value(activityTypeID)
	} else {
		model.ActivityTypeID = types.Int64Null()
	}

	if campaignID, ok := GetInt64(result, "campaign_id"); ok {
		model.CampaignID = types.Int64Value(campaignID)
	} else {
		model.CampaignID = types.Int64Null()
	}

	if activitySource, ok := GetString(result, "activity_source"); ok && activitySource != "" {
		model.ActivitySource = types.StringValue(activitySource)
	} else {
		model.ActivitySource = types.StringNull()
	}

	if activityTargets, ok := GetString(result, "activity_targets"); ok && activityTargets != "" {
		model.ActivityTargets = types.StringValue(activityTargets)
	} else {
		model.ActivityTargets = types.StringNull()
	}

	if activityAssignees, ok := GetString(result, "activity_assignees"); ok && activityAssignees != "" {
		model.ActivityAssignees = types.StringValue(activityAssignees)
	} else {
		model.ActivityAssignees = types.StringNull()
	}
}
