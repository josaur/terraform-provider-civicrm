# Development Notes

Internal notes for contributors. Not published to the Terraform Registry.

## CiviCRM API v4 Behavior

### Sparse Create/Update responses

CiviCRM API v4 returns only a minimal set of fields from `create` and `update` operations — typically just `id`, the fields you explicitly changed, `custom`, and `check_permissions`. It does **not** return computed defaults (e.g. `weight`, `domain_id`, `visibility_id`, `options_per_line`).

**Consequence for the provider:** After every `create` and `update`, the provider must do a separate `GetByID` call to load the full state. All resource implementations follow this pattern:

```go
// Create
result, err := r.client.Create("Entity", values)
// ...error check...
if createdID, ok := GetInt64(result, "id"); ok {
    if fullResult, err2 := r.client.GetByID("Entity", createdID, nil); err2 == nil {
        result = fullResult
    }
}
r.mapResultToState(result, &plan)

// Update
_, err := r.client.Update("Entity", state.ID.ValueInt64(), values)
// ...error check...
result, err := r.client.GetByID("Entity", state.ID.ValueInt64(), nil)
// ...error check...
r.mapResultToState(result, &plan)
```

If you skip this and use the sparse response directly, Terraform will report "provider produced inconsistent result" because computed fields remain Unknown after apply.

### `extends` is returned as an array

For `PriceSet` (and similar entities), CiviCRM returns `extends` as a JSON array: `["CiviMember"]`. The `GetString` helper cannot parse arrays. The `mapResultToState` for PriceSet handles this with a type switch:

```go
if raw, ok := result["extends"]; ok && raw != nil {
    switch v := raw.(type) {
    case string:
        model.Extends = types.StringValue(v)
    case []any:
        if len(v) > 0 {
            if s, ok := v[0].(string); ok {
                model.Extends = types.StringValue(s)
            }
        }
    }
}
```

If you add a new resource with an `extends`-like field that CiviCRM stores as an array, apply the same pattern.

### The string `"null"` vs JSON `null`

CiviCRM occasionally serializes a PHP `null` as the JSON string `"null"` rather than JSON `null`. The `GetString` helper in `client.go` treats the string `"null"` as a missing value and returns `("", false)` so it maps to `types.StringNull()` in state.

### `extends_entity_column_value: ["null"]`

`CustomGroup.extends_entity_column_value` is returned as `["null"]` (a list containing the string "null") when no subtypes are selected. The `mapResultToState` filters out the string `"null"` from this list.

### Zero-value defaults for Optional fields

CiviCRM fills in `0` for unset integer/float columns. For fields that are `Optional` only (not `Computed`) in the schema — such as `min_amount`, `minimum_fee`, `extends_entity_column_id`, `min_multiple`, `max_multiple`, `non_deductible_amount` — a returned value of `0` must be mapped to `null`, not `Int64Value(0)` / `Float64Value(0)`. Otherwise Terraform diffs will show phantom changes.

### `domain_id` must not be sent as `0`

`PriceSet`, `MailSettings`, and `SiteEmailAddress` all have a `domain_id` FK that references `civicrm_domain.id`. The only valid value in a standard install is `1`. Sending `0` causes a FK constraint violation.

`domain_id` is `Optional + Computed` in the schema. During planning, Terraform marks unset computed fields as `Unknown` (not `Null`). A guard of `if !plan.DomainID.IsNull()` is **not sufficient** — you must also check `!plan.DomainID.IsUnknown()`.

This applies to all `Optional + Computed` fields: always guard with both `IsNull()` and `IsUnknown()` before reading `ValueXxx()`.

## Terraform Plugin Framework Patterns

### Computed fields and Unknown values

When a field is `Optional: true, Computed: true` and the user does not set it, Terraform plans it as `Unknown` ("known after apply"). Calling `ValueInt64()` on an Unknown value returns `0`, not a meaningful default. Always guard:

```go
if !plan.Foo.IsNull() && !plan.Foo.IsUnknown() {
    values["foo"] = plan.Foo.ValueInt64()
}
```

### `UseStateForUnknown` for stable computed fields

Fields that CiviCRM assigns once and never changes on its own (e.g. `table_name` on `CustomGroup`, `id` on all resources) should use `UseStateForUnknown()` plan modifier. This prevents Terraform from marking them as "known after apply" on every update:

```go
PlanModifiers: []planmodifier.String{
    stringplanmodifier.UseStateForUnknown(),
},
```

### ImportState: fields not returned consistently

Some fields are not included in read/get responses or behave differently during import:
- `ACLRole.value` — sometimes returned as the string `"null"`, now handled in `GetString`
- `CiviRulesRuleCondition.negate` — not returned when `false`; add to `ImportStateVerifyIgnore`

When adding a new resource, run `ImportState` and compare carefully. Use `ImportStateVerifyIgnore` only as a last resort when the field genuinely cannot be read back, not as a workaround for provider bugs.

## Running Acceptance Tests

```bash
export CIVICRM_URL="http://your-instance"
export CIVICRM_API_KEY="your-key"
make testacc

# Single test:
TF_ACC=1 go test -v ./internal/provider/... -run TestAccPriceSetResource -timeout 10m
```

Tests require `TF_ACC=1` and a live CiviCRM instance. Without `TF_ACC=1`, `go test ./...` runs normally and skips all acceptance tests.

### Orphaned test data

If a test run is interrupted before the destroy step, CiviCRM may retain test data. All test resources use the prefix `tf_acc_` for easy identification. `CustomGroup` also uses a predictable `table_name` (`civicrm_value_tf_acc_cg`) because CiviCRM's delete logic fails if `table_name` is null.

To clean up manually:
```bash
curl -X POST "$CIVICRM_URL/civicrm/ajax/api4/ENTITY/get" \
  -H "X-Civi-Auth: Bearer $CIVICRM_API_KEY" \
  --data-urlencode 'params={"where":[["name","LIKE","tf_acc%"]],"select":["id","name"]}'
```

### CiviRules tests

`TestAccCiviRulesRuleConditionResource` hardcodes `condition_id = 1`. If your test instance has different condition IDs (e.g. CiviRules is not installed or has different data), adjust the ID in the test config.

## Adding a New Resource

1. Create `internal/provider/resource_<name>.go`
2. Follow the read-after-write pattern for both `Create` and `Update` (see above)
3. In `mapResultToState`, handle sparse/null returns defensively:
   - Use `GetString`, `GetInt64`, `GetBool` helpers
   - Always set the model field to a null value in the `else` branch
   - Treat zero values as null for `Optional`-only numeric fields
   - Treat `"null"` string as null (handled by `GetString`)
4. For all `Optional + Computed` fields, guard with `IsNull() && IsUnknown()` before reading values
5. Add `internal/provider/resource_<name>_test.go` with Create+Check, ImportState, Update+Check steps
6. Register the resource in `provider.go`
7. Run `make testacc` (or the single-test variant) against a live instance
