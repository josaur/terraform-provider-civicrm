package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &NavigationResource{}
	_ resource.ResourceWithConfigure   = &NavigationResource{}
	_ resource.ResourceWithImportState = &NavigationResource{}
)

// navigationPermissionOperators are the values Navigation.permission_operator
// accepts, from Navigation.getFields(loadOptions).
var navigationPermissionOperators = []string{"AND", "OR"}

// navigationSelect is passed to every GetByID/Get call for this resource, so
// state always reflects every field this resource manages, independent of
// which fields a bare select=nil call happens to return by default.
var navigationSelect = []string{
	"id", "domain_id", "label", "name", "url", "icon", "permission",
	"permission_operator", "parent_id", "is_active", "has_separator", "weight",
}

type NavigationResource struct {
	client *Client
}

type NavigationResourceModel struct {
	ID                 types.Int64  `tfsdk:"id"`
	DomainID           types.Int64  `tfsdk:"domain_id"`
	Label              types.String `tfsdk:"label"`
	Name               types.String `tfsdk:"name"`
	URL                types.String `tfsdk:"url"`
	Icon               types.String `tfsdk:"icon"`
	Permission         types.List   `tfsdk:"permission"`
	PermissionOperator types.String `tfsdk:"permission_operator"`
	ParentID           types.Int64  `tfsdk:"parent_id"`
	IsActive           types.Bool   `tfsdk:"is_active"`
	HasSeparator       types.Int64  `tfsdk:"has_separator"`
	Weight             types.Int64  `tfsdk:"weight"`
}

func NewNavigationResource() resource.Resource {
	return &NavigationResource{}
}

func (r *NavigationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_navigation"
}

func (r *NavigationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a CiviCRM Navigation menu entry. Navigation entries carry a " +
			"`permission` list and `permission_operator`, making the menu the place where the " +
			"interface is tailored per user group: each group sees only the entries its " +
			"permissions allow. Menu nesting is expressed via the self-referencing `parent_id`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "The unique identifier of the navigation entry.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"domain_id": schema.Int64Attribute{
				Description: "FK to the domain this entry belongs to. Defaults to the current domain.",
				Optional:    true,
				Computed:    true,
			},
			"label": schema.StringAttribute{
				Description: "Menu title shown in the UI.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Internal machine name of the menu entry.",
				Optional:    true,
			},
			"url": schema.StringAttribute{
				Description: "Target URL, for custom links. Leave unset for entries that only " +
					"group children (e.g. a top-level menu with no link of its own).",
				Optional: true,
			},
			"icon": schema.StringAttribute{
				Description: "CSS class of the icon shown next to the label.",
				Optional:    true,
			},
			"permission": schema.ListAttribute{
				Description: "Permissions required to see this menu entry, as a list of strings " +
					"(e.g. [\"access CiviCRM\"]). CiviCRM stores this comma-separated internally " +
					"(SERIALIZE_COMMA) but the API accepts and returns it as a list; pass a list here, " +
					"not a comma-joined string.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"permission_operator": schema.StringAttribute{
				Description: "How multiple permission entries combine: \"AND\" or \"OR\". Only has an " +
					"effect with two or more permissions; leave unset for zero or one permission.",
				Optional: true,
				Validators: []validator.String{
					stringOneOf(navigationPermissionOperators...),
				},
			},
			"parent_id": schema.Int64Attribute{
				Description: "FK to another civicrm_navigation entry this one is nested under. Use " +
					"the civicrm_navigation data source to resolve an existing top-level entry (e.g. " +
					"\"Search\", \"Administer\") by name rather than hardcoding its id.",
				Optional: true,
			},
			"is_active": schema.BoolAttribute{
				Description: "Whether the entry is active. Default: true.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"has_separator": schema.Int64Attribute{
				Description: "Separator line around this entry: 0 = none, 1 = after, 2 = before. Default: 0.",
				Optional:    true,
				Computed:    true,
				Validators: []validator.Int64{
					int64OneOf(0, 1, 2),
				},
			},
			"weight": schema.Int64Attribute{
				Description: "Ordering of the entry among its siblings. If omitted, CiviCRM assigns " +
					"the next free weight for the given parent.",
				Optional: true,
				Computed: true,
			},
		},
	}
}

func (r *NavigationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *NavigationResource) valuesFromPlan(ctx context.Context, plan *NavigationResourceModel, forUpdate bool) (map[string]any, error) {
	values := map[string]any{
		"is_active": plan.IsActive.ValueBool(),
	}

	setOrNull := func(key string, isNull bool, val any) {
		if !isNull {
			values[key] = val
		} else if forUpdate {
			values[key] = nil
		}
	}

	// setIfPresent only sets key when the plan has a known, non-null value —
	// never explicitly nulls it out. domain_id, has_separator and weight are
	// Optional+Computed fields CiviCRM defaults/recomputes server-side
	// (domain_id to the current domain, has_separator/weight to 0/next-free);
	// explicitly nulling domain_id on update sends `domain_id = 0`, which
	// violates the FK constraint to civicrm_domain.
	setIfPresent := func(key string, isAbsent bool, val any) {
		if !isAbsent {
			values[key] = val
		}
	}

	setIfPresent("domain_id", plan.DomainID.IsNull() || plan.DomainID.IsUnknown(), plan.DomainID.ValueInt64())
	setOrNull("label", plan.Label.IsNull(), plan.Label.ValueString())
	setOrNull("name", plan.Name.IsNull(), plan.Name.ValueString())
	setOrNull("url", plan.URL.IsNull(), plan.URL.ValueString())
	setOrNull("icon", plan.Icon.IsNull(), plan.Icon.ValueString())
	setOrNull("permission_operator", plan.PermissionOperator.IsNull(), plan.PermissionOperator.ValueString())
	setOrNull("parent_id", plan.ParentID.IsNull(), plan.ParentID.ValueInt64())
	setIfPresent("has_separator", plan.HasSeparator.IsNull() || plan.HasSeparator.IsUnknown(), plan.HasSeparator.ValueInt64())
	setIfPresent("weight", plan.Weight.IsNull() || plan.Weight.IsUnknown(), plan.Weight.ValueInt64())

	if !plan.Permission.IsNull() && !plan.Permission.IsUnknown() {
		var permissions []string
		diags := plan.Permission.ElementsAs(ctx, &permissions, false)
		if diags.HasError() {
			return nil, fmt.Errorf("could not read permission list")
		}
		values["permission"] = permissions
	} else if forUpdate {
		values["permission"] = nil
	}

	return values, nil
}

func (r *NavigationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NavigationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating Navigation", map[string]any{"label": plan.Label.ValueString()})

	values, err := r.valuesFromPlan(ctx, &plan, false)
	if err != nil {
		resp.Diagnostics.AddAttributeError(path.Root("permission"), "Invalid permission", err.Error())
		return
	}

	result, err := r.client.Create("Navigation", values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Navigation",
			"Could not create Navigation: "+err.Error(),
		)
		return
	}

	if createdID, ok := GetInt64(result, "id"); ok {
		if fullResult, err2 := r.client.GetByID("Navigation", createdID, navigationSelect); err2 == nil {
			result = fullResult
		}
	}

	r.mapResultToState(ctx, result, &plan, &resp.Diagnostics)

	tflog.Debug(ctx, "Created Navigation", map[string]any{"id": plan.ID.ValueInt64()})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *NavigationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NavigationResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading Navigation", map[string]any{"id": state.ID.ValueInt64()})

	result, err := r.client.GetByID("Navigation", state.ID.ValueInt64(), navigationSelect)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Navigation",
			"Could not read Navigation ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	r.mapResultToState(ctx, result, &state, &resp.Diagnostics)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *NavigationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NavigationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var state NavigationResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating Navigation", map[string]any{"id": state.ID.ValueInt64()})

	values, err := r.valuesFromPlan(ctx, &plan, true)
	if err != nil {
		resp.Diagnostics.AddAttributeError(path.Root("permission"), "Invalid permission", err.Error())
		return
	}

	_, err = r.client.Update("Navigation", state.ID.ValueInt64(), values)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Navigation",
			"Could not update Navigation ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	plan.ID = state.ID

	result, err := r.client.GetByID("Navigation", state.ID.ValueInt64(), navigationSelect)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Navigation after update",
			"Could not re-read Navigation ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}
	r.mapResultToState(ctx, result, &plan, &resp.Diagnostics)

	tflog.Debug(ctx, "Updated Navigation", map[string]any{"id": plan.ID.ValueInt64()})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *NavigationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NavigationResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting Navigation", map[string]any{"id": state.ID.ValueInt64()})

	err := r.client.Delete("Navigation", state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Navigation",
			"Could not delete Navigation ID "+strconv.FormatInt(state.ID.ValueInt64(), 10)+": "+err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "Deleted Navigation", map[string]any{"id": state.ID.ValueInt64()})
}

func (r *NavigationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (r *NavigationResource) mapResultToState(ctx context.Context, result map[string]any, model *NavigationResourceModel, diags *diag.Diagnostics) {
	if id, ok := GetInt64(result, "id"); ok {
		model.ID = types.Int64Value(id)
	}

	if v, ok := GetInt64(result, "domain_id"); ok {
		model.DomainID = types.Int64Value(v)
	} else {
		model.DomainID = types.Int64Null()
	}

	if v, ok := GetString(result, "label"); ok && v != "" {
		model.Label = types.StringValue(v)
	} else {
		model.Label = types.StringNull()
	}

	if v, ok := GetString(result, "name"); ok && v != "" {
		model.Name = types.StringValue(v)
	} else {
		model.Name = types.StringNull()
	}

	if v, ok := GetString(result, "url"); ok && v != "" {
		model.URL = types.StringValue(v)
	} else {
		model.URL = types.StringNull()
	}

	if v, ok := GetString(result, "icon"); ok && v != "" {
		model.Icon = types.StringValue(v)
	} else {
		model.Icon = types.StringNull()
	}

	if permRaw, ok := result["permission"]; ok && permRaw != nil {
		if permSlice, ok := permRaw.([]any); ok {
			values := make([]string, 0, len(permSlice))
			for _, v := range permSlice {
				if s, ok := v.(string); ok {
					values = append(values, s)
				}
			}
			if len(values) > 0 {
				permList, d := types.ListValueFrom(ctx, types.StringType, values)
				diags.Append(d...)
				model.Permission = permList
			} else {
				model.Permission = types.ListNull(types.StringType)
			}
		} else {
			model.Permission = types.ListNull(types.StringType)
		}
	} else {
		model.Permission = types.ListNull(types.StringType)
	}

	if v, ok := GetString(result, "permission_operator"); ok && v != "" {
		model.PermissionOperator = types.StringValue(v)
	} else {
		model.PermissionOperator = types.StringNull()
	}

	if v, ok := GetInt64(result, "parent_id"); ok {
		model.ParentID = types.Int64Value(v)
	} else {
		model.ParentID = types.Int64Null()
	}

	if v, ok := GetBool(result, "is_active"); ok {
		model.IsActive = types.BoolValue(v)
	}

	if v, ok := GetInt64(result, "has_separator"); ok {
		model.HasSeparator = types.Int64Value(v)
	} else {
		model.HasSeparator = types.Int64Value(0)
	}

	if v, ok := GetInt64(result, "weight"); ok {
		model.Weight = types.Int64Value(v)
	} else {
		model.Weight = types.Int64Null()
	}
}
