---
page_title: "civicrm_civirules_rule_condition Resource - CiviCRM"
subcategory: "CiviRules"
description: |-
  Attaches a condition to a CiviRules rule.
---

# civicrm_civirules_rule_condition (Resource)

Attaches a condition to a CiviRules rule (entity: `CiviRulesRuleCondition`). One rule can have multiple conditions; by default all conditions must pass (AND logic) for the rule to execute. The `condition_id` references a CiviRulesCondition type (look up via [`civicrm_civirules_condition`](../data-sources/civirules_condition.md)).

## Example Usage

CiviRules stores `condition_params` via `serialize()` and reads them with
`unserialize()` (see `CRM/Civirules/Condition/*Form.php`). Values written
as JSON make the condition unreadable in the UI and skip the condition's
runtime check silently. Pass a PHP-serialized string.

```terraform
data "civicrm_civirules_condition" "activity_type" {
  name = "activity_type"
}

# Only proceed if the activity type is "Follow up"
resource "civicrm_civirules_rule_condition" "check_type" {
  rule_id      = civicrm_civirules_rule.my_rule.id
  condition_id = data.civicrm_civirules_condition.activity_type.id
  # a:1:{s:16:"activity_type_id";s:9:"Follow up";}
  condition_params = "a:1:{s:16:\"activity_type_id\";s:9:\"Follow up\";}"
  is_active = true
  negate    = false
}
```

## Argument Reference

The following arguments are supported:

### Required

- `rule_id` (Number) ID of the [`civicrm_civirules_rule`](civirules_rule.md) this condition belongs to.
- `condition_id` (Number) ID of the CiviRulesCondition type to apply. Look up via [`civicrm_civirules_condition`](../data-sources/civirules_condition.md) or the CiviRules UI.

### Optional

- `condition_params` (String) PHP `serialize()`-encoded parameters passed to the condition class. Structure depends on the condition type (e.g. `a:1:{s:12:"case_type_id";s:1:"3";}` for a case-type condition). See note in *Example Usage* above; JSON is silently accepted at write time but breaks the condition at runtime.
- `is_active` (Boolean) Whether this condition is active. Default: `true`.
- `negate` (Boolean) When `true`, the condition logic is inverted (NOT). Default: `false`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (Number) The unique identifier of the rule-condition link.

## Import

CiviRules Rule Conditions can be imported using the record ID:

```shell
terraform import civicrm_civirules_rule_condition.example 1
```
