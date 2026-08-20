package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	_ resource.Resource                = &CiviRulesRuleActionResource{}
	_ resource.ResourceWithConfigure   = &CiviRulesRuleActionResource{}
	_ resource.ResourceWithImportState = &CiviRulesRuleActionResource{}
)

// CiviRulesRuleActionResource links an action type to a rule (entity: CiviRulesRuleAction).
// One rule can have multiple actions which are all executed when the rule fires.
type CiviRulesRuleActionResource struct {
	client *Client
}

type CiviRulesRuleActionResourceModel struct {
	ID           types.Int64          `tfsdk:"id"`
	RuleID       types.Int64          `tfsdk:"rule_id"`
	ActionID     types.Int64          `tfsdk:"action_id"`
	ActionParams jsontypes.Normalized `tfsdk:"action_params"`
	Delay        types.String         `tfsdk:"delay"`
	IsActive     types.Bool           `tfsdk:"is_active"`
}

func NewCiviRulesRuleActionResource() resource.Resource {
	return &CiviRulesRuleActionResource{}
}

func (r *CiviRulesRuleActionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_civirules_rule_action"
}

func (r *CiviRulesRuleActionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Attaches an action to a CiviRules rule (entity: CiviRulesRuleAction). The action_id references a CiviRulesAction type. action_params is a JSON string whose structure depends on the action class; the resource converts it to/from the PHP serialize() format CiviRules stores internally.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Unique ID of the rule-action link.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"rule_id": schema.Int64Attribute{
				Description: "ID of the civicrm_civirules_rule this action belongs to.",
				Required:    true,
			},
			"action_id": schema.Int64Attribute{
				Description: "ID of the CiviRulesAction type to execute. Look up available action IDs via the CiviRules UI or CiviRulesAction.get API.",
				Required:    true,
			},
			"action_params": schema.StringAttribute{
				Description: "Parameters passed to the action class, as a JSON string (e.g. jsonencode({status_id = 2})). CiviRules itself stores this PHP serialize()-encoded and reads it with unserialize(); this attribute handles that conversion automatically, so config only ever deals with JSON. Structure depends on the action type.",
				Optional:    true,
				CustomType:  jsontypes.NormalizedType{},
			},
			"delay": schema.StringAttribute{
				Description: "Optional delay before executing the action. Format: ISO 8601 duration string (e.g. 'P1D' = 1 day, 'PT2H' = 2 hours) or empty for immediate execution.",
				Optional:    true,
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether this action is active. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
		},
	}
}

func (r *CiviRulesRuleActionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *Client, got: %T.", req.ProviderData))
		return
	}
	r.client = client
}

func (r *CiviRulesRuleActionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CiviRulesRuleActionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating CiviRulesRuleAction", map[string]any{
		"rule_id": plan.RuleID.ValueInt64(), "action_id": plan.ActionID.ValueInt64(),
	})

	values := map[string]any{
		"rule_id":   plan.RuleID.ValueInt64(),
		"action_id": plan.ActionID.ValueInt64(),
		"is_active": plan.IsActive.ValueBool(),
	}
	if !plan.ActionParams.IsNull() && !plan.ActionParams.IsUnknown() && plan.ActionParams.ValueString() != "" {
		decoded, err := decodeJSONAttribute(plan.ActionParams.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("action_params"),
				"Invalid action_params JSON",
				"action_params must be valid JSON (e.g. produced by jsonencode(...)): "+err.Error(),
			)
			return
		}
		values["action_params"] = phpSerialize(decoded)
	}
	if !plan.Delay.IsNull() && !plan.Delay.IsUnknown() && plan.Delay.ValueString() != "" {
		values["delay"] = plan.Delay.ValueString()
	}

	result, err := r.client.Create("CiviRulesRuleAction", values)
	if err != nil {
		resp.Diagnostics.AddError("Error creating CiviRulesRuleAction", err.Error())
		return
	}

	// Re-read to get complete state (Create response is sparse)
	if createdID, ok := GetInt64(result, "id"); ok {
		if fullResult, err2 := r.client.GetByID("CiviRulesRuleAction", createdID, nil); err2 == nil {
			result = fullResult
		}
	}
	r.mapResultToState(result, &plan, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *CiviRulesRuleActionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CiviRulesRuleActionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.GetByID("CiviRulesRuleAction", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading CiviRulesRuleAction",
			"Could not read CiviRulesRuleAction ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error())
		return
	}

	r.mapResultToState(result, &state, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *CiviRulesRuleActionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CiviRulesRuleActionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var state CiviRulesRuleActionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	values := map[string]any{
		"rule_id":   plan.RuleID.ValueInt64(),
		"action_id": plan.ActionID.ValueInt64(),
		"is_active": plan.IsActive.ValueBool(),
	}
	if !plan.ActionParams.IsNull() && !plan.ActionParams.IsUnknown() && plan.ActionParams.ValueString() != "" {
		decoded, err := decodeJSONAttribute(plan.ActionParams.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("action_params"),
				"Invalid action_params JSON",
				"action_params must be valid JSON (e.g. produced by jsonencode(...)): "+err.Error(),
			)
			return
		}
		values["action_params"] = phpSerialize(decoded)
	} else {
		values["action_params"] = nil
	}
	if !plan.Delay.IsNull() && !plan.Delay.IsUnknown() && plan.Delay.ValueString() != "" {
		values["delay"] = plan.Delay.ValueString()
	} else {
		values["delay"] = nil
	}

	_, err := r.client.Update("CiviRulesRuleAction", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError("Error updating CiviRulesRuleAction",
			"Could not update CiviRulesRuleAction ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error())
		return
	}

	plan.ID = state.ID

	result, err := r.client.GetByID("CiviRulesRuleAction", state.ID.ValueInt64(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading CiviRulesRuleAction after update",
			"Could not re-read CiviRulesRuleAction ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}
	r.mapResultToState(result, &plan, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *CiviRulesRuleActionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CiviRulesRuleActionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.Delete("CiviRulesRuleAction", state.ID.ValueInt64()); err != nil {
		resp.Diagnostics.AddError("Error deleting CiviRulesRuleAction",
			"Could not delete CiviRulesRuleAction ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error())
	}
}

func (r *CiviRulesRuleActionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", "Could not parse import ID as integer: "+err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (r *CiviRulesRuleActionResource) mapResultToState(result map[string]any, model *CiviRulesRuleActionResourceModel, diags *diag.Diagnostics) {
	if id, ok := GetInt64(result, "id"); ok {
		model.ID = types.Int64Value(id)
	}
	if v, ok := GetInt64(result, "rule_id"); ok {
		model.RuleID = types.Int64Value(v)
	}
	if v, ok := GetInt64(result, "action_id"); ok {
		model.ActionID = types.Int64Value(v)
	}
	if v, ok := GetString(result, "action_params"); ok && v != "" {
		decoded, err := phpUnserialize(v)
		if err != nil {
			diags.AddAttributeError(
				path.Root("action_params"),
				"Could not decode action_params",
				"CiviRules returned action_params that is not valid PHP serialize() data: "+err.Error(),
			)
			return
		}
		encoded, err := encodeJSONAttribute(decoded)
		if err != nil {
			diags.AddAttributeError(
				path.Root("action_params"),
				"Could not encode action_params",
				err.Error(),
			)
			return
		}
		model.ActionParams = jsontypes.NewNormalizedValue(encoded)
	} else {
		model.ActionParams = jsontypes.NewNormalizedNull()
	}
	if v, ok := GetString(result, "delay"); ok && v != "" {
		model.Delay = types.StringValue(v)
	} else {
		model.Delay = types.StringNull()
	}
	if v, ok := GetBool(result, "is_active"); ok {
		model.IsActive = types.BoolValue(v)
	}
}
