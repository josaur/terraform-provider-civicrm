---
page_title: "civicrm_civirules_rule Resource - CiviCRM"
subcategory: "CiviRules"
description: |-
  Manages a CiviRules Rule.
---

# civicrm_civirules_rule (Resource)

Manages a CiviRules Rule. A rule is the top-level object that pairs one trigger ("when") with optional conditions ("if") and one or more actions ("then"). Conditions are managed via [`civicrm_civirules_rule_condition`](civirules_rule_condition.md), actions via [`civicrm_civirules_rule_action`](civirules_rule_action.md).

Requires the [CiviRules extension](https://civicrm.org/extensions/civirules) to be installed.

## Example Usage

CiviRules stores `trigger_params` (like `condition_params` and
`action_params` on the linked condition/action resources) via `serialize()`
and reads it back with `unserialize()` (see
`CRM_Civirules_Trigger::setTriggerParams()`). Values written as JSON are
silently accepted at write time but the trigger fails to read them at
runtime because `unserialize()` cannot decode them. Pass a PHP-serialized
string — the same applies to `condition_params` on
[`civicrm_civirules_rule_condition`](civirules_rule_condition.md) and
`action_params` on [`civicrm_civirules_rule_action`](civirules_rule_action.md).

```terraform
# Look up the trigger by name so we don't hardcode its numeric ID
data "civicrm_civirules_trigger" "activity_changed" {
  name = "activity_is_changed"
}

data "civicrm_civirules_action" "change_case_status" {
  name = "change_case_status"
}

data "civicrm_civirules_condition" "activity_type" {
  name = "activity_type"
}

data "civicrm_case_status" "housing_secured" {
  name = "housing_secured"
}

# The rule itself
resource "civicrm_civirules_rule" "income_stab_closes_case" {
  name        = "housing_income_stab_to_secured"
  label       = "Housing: Income stabilization → Housing Secured"
  trigger_id  = data.civicrm_civirules_trigger.activity_changed.id
  description = "When income stabilization activity completes, set case status to Housing Secured."
  is_active   = true
}

# Condition: only for activities of type "Income and benefits stabilization"
resource "civicrm_civirules_rule_condition" "check_activity_type" {
  rule_id      = civicrm_civirules_rule.income_stab_closes_case.id
  condition_id = data.civicrm_civirules_condition.activity_type.id
  # a:1:{s:16:"activity_type_id";s:32:"Income and benefits stabilization";}
  condition_params = "a:1:{s:16:\"activity_type_id\";s:32:\"Income and benefits stabilization\";}"
  is_active = true
}

# Action: change case status
resource "civicrm_civirules_rule_action" "set_status" {
  rule_id   = civicrm_civirules_rule.income_stab_closes_case.id
  action_id = data.civicrm_civirules_action.change_case_status.id
  # a:1:{s:14:"case_status_id";i:2;}
  action_params = "a:1:{s:14:\"case_status_id\";i:2;}"
  is_active = true
}
```

## Argument Reference

The following arguments are supported:

### Required

- `name` (String) Machine name of the rule. Must be unique. Used for managed entity matching.
- `label` (String) Human-readable label shown in the CiviRules UI.
- `trigger_id` (Number) ID of the CiviRules trigger that fires this rule. Look up available triggers via the data source [`civicrm_civirules_trigger`](../data-sources/civirules_trigger.md) or the CiviRules UI.

### Optional

- `trigger_params` (String) PHP `serialize()`-encoded parameters passed to the trigger class. Structure depends on the trigger type. See note in *Example Usage* above; JSON is silently accepted at write time but breaks the trigger at runtime. Leave empty if the trigger requires no parameters.
- `description` (String) Optional description of what this rule does.
- `help_text` (String) Optional help text shown to admins in the CiviRules UI.
- `is_active` (Boolean) Whether the rule is active. Default: `true`.
- `is_debug` (Boolean) Enable debug logging for this rule. Default: `false`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (Number) The unique identifier of the rule.

## Import

CiviRules Rules can be imported using the rule ID:

```shell
terraform import civicrm_civirules_rule.example 1
```
