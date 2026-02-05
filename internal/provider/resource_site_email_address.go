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
	_ resource.Resource                = &SiteEmailAddressResource{}
	_ resource.ResourceWithConfigure   = &SiteEmailAddressResource{}
	_ resource.ResourceWithImportState = &SiteEmailAddressResource{}
)

// SiteEmailAddressResource manages site email addresses in CiviCRM.
type SiteEmailAddressResource struct {
	client *Client
}

type SiteEmailAddressResourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	DisplayName types.String `tfsdk:"display_name"`
	Email       types.String `tfsdk:"email"`
	Description types.String `tfsdk:"description"`
	IsActive    types.Bool   `tfsdk:"is_active"`
	IsDefault   types.Bool   `tfsdk:"is_default"`
	DomainID    types.Int64  `tfsdk:"domain_id"`
}

func NewSiteEmailAddressResource() resource.Resource {
	return &SiteEmailAddressResource{}
}

func (r *SiteEmailAddressResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_email_address"
}

func (r *SiteEmailAddressResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages CiviCRM Site Email Addresses used as sender addresses for outgoing emails.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier of the site email address.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"display_name": schema.StringAttribute{
				Description: "The display name shown as the sender name (e.g., 'CiviCRM Support').",
				Required:    true,
			},
			"email": schema.StringAttribute{
				Description: "The email address used as the sender address.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of this email address configuration.",
				Optional:    true,
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether this email address is active. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"is_default": schema.BoolAttribute{
				Description: "Whether this is the default email address. Default: false.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"domain_id": schema.Int64Attribute{
				Description: "The domain ID this email address belongs to.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func (r *SiteEmailAddressResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SiteEmailAddressResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SiteEmailAddressResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating site email address", map[string]any{
		"display_name": plan.DisplayName.ValueString(),
		"email":        plan.Email.ValueString(),
	})

	// Build values for API call
	values := map[string]any{
		"display_name": plan.DisplayName.ValueString(),
		"email":        plan.Email.ValueString(),
		"is_active":    plan.IsActive.ValueBool(),
		"is_default":   plan.IsDefault.ValueBool(),
	}

	if !plan.Description.IsNull() {
		values["description"] = plan.Description.ValueString()
	}

	if !plan.DomainID.IsNull() {
		values["domain_id"] = plan.DomainID.ValueInt64()
	}

	// Call API
	result, err := r.client.Create("SiteEmailAddress", values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating site email address",
			"Could not create site email address, unexpected error: "+err.Error(),
		)
		return
	}

	// Update state with response
	if id, ok := GetInt64(result, "id"); ok {
		plan.ID = types.Int64Value(id)
	}

	if displayName, ok := GetString(result, "display_name"); ok {
		plan.DisplayName = types.StringValue(displayName)
	}

	if email, ok := GetString(result, "email"); ok {
		plan.Email = types.StringValue(email)
	}

	if description, ok := GetString(result, "description"); ok && description != "" {
		plan.Description = types.StringValue(description)
	} else {
		plan.Description = types.StringNull()
	}

	if isActive, ok := GetBool(result, "is_active"); ok {
		plan.IsActive = types.BoolValue(isActive)
	}

	if isDefault, ok := GetBool(result, "is_default"); ok {
		plan.IsDefault = types.BoolValue(isDefault)
	}

	if domainID, ok := GetInt64(result, "domain_id"); ok {
		plan.DomainID = types.Int64Value(domainID)
	}

	tflog.Debug(ctx, "Created site email address", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SiteEmailAddressResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SiteEmailAddressResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading site email address", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	result, err := r.client.GetByID("SiteEmailAddress", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading site email address",
			"Could not read site email address ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Update state
	if displayName, ok := GetString(result, "display_name"); ok {
		state.DisplayName = types.StringValue(displayName)
	}

	if email, ok := GetString(result, "email"); ok {
		state.Email = types.StringValue(email)
	}

	if description, ok := GetString(result, "description"); ok && description != "" {
		state.Description = types.StringValue(description)
	} else {
		state.Description = types.StringNull()
	}

	if isActive, ok := GetBool(result, "is_active"); ok {
		state.IsActive = types.BoolValue(isActive)
	}

	if isDefault, ok := GetBool(result, "is_default"); ok {
		state.IsDefault = types.BoolValue(isDefault)
	}

	if domainID, ok := GetInt64(result, "domain_id"); ok {
		state.DomainID = types.Int64Value(domainID)
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *SiteEmailAddressResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SiteEmailAddressResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state SiteEmailAddressResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating site email address", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	// Build values for API call
	values := map[string]any{
		"display_name": plan.DisplayName.ValueString(),
		"email":        plan.Email.ValueString(),
		"is_active":    plan.IsActive.ValueBool(),
		"is_default":   plan.IsDefault.ValueBool(),
	}

	if !plan.Description.IsNull() {
		values["description"] = plan.Description.ValueString()
	} else {
		values["description"] = nil
	}

	if !plan.DomainID.IsNull() {
		values["domain_id"] = plan.DomainID.ValueInt64()
	}

	// Call API
	result, err := r.client.Update("SiteEmailAddress", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating site email address",
			"Could not update site email address ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	// Update state
	plan.ID = state.ID

	if displayName, ok := GetString(result, "display_name"); ok {
		plan.DisplayName = types.StringValue(displayName)
	}

	if email, ok := GetString(result, "email"); ok {
		plan.Email = types.StringValue(email)
	}

	if description, ok := GetString(result, "description"); ok && description != "" {
		plan.Description = types.StringValue(description)
	} else {
		plan.Description = types.StringNull()
	}

	if isActive, ok := GetBool(result, "is_active"); ok {
		plan.IsActive = types.BoolValue(isActive)
	}

	if isDefault, ok := GetBool(result, "is_default"); ok {
		plan.IsDefault = types.BoolValue(isDefault)
	}

	if domainID, ok := GetInt64(result, "domain_id"); ok {
		plan.DomainID = types.Int64Value(domainID)
	}

	tflog.Debug(ctx, "Updated site email address", map[string]any{
		"id": plan.ID.ValueInt64(),
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SiteEmailAddressResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SiteEmailAddressResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting site email address", map[string]any{
		"id": state.ID.ValueInt64(),
	})

	err := r.client.Delete("SiteEmailAddress", state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting site email address",
			"Could not delete site email address ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "Deleted site email address", map[string]any{
		"id": state.ID.ValueInt64(),
	})
}

func (r *SiteEmailAddressResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
