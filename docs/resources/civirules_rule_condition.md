---
page_title: "civicrm_civirules_rule_condition Resource - CiviCRM"
subcategory: "CiviRules"
description: |-
  Attaches a condition to a CiviRules rule.
---

# civicrm_civirules_rule_condition (Resource)

Attaches a condition to a CiviRules rule (entity: `CiviRulesRuleCondition`). One rule can have multiple conditions; by default all conditions must pass (AND logic) for the rule to execute. The `condition_id` references a CiviRulesCondition type (look up via [`civicrm_civirules_condition`](../data-sources/civirules_condition.md)).

## Example Usage

```terraform
data "civicrm_civirules_condition" "activity_type" {
  name = "activity_type"
}

# Only proceed if the activity type is "Follow up"
resource "civicrm_civirules_rule_condition" "check_type" {
  rule_id      = civicrm_civirules_rule.my_rule.id
  condition_id = data.civicrm_civirules_condition.activity_type.id
  condition_params = jsonencode({
    activity_type_id = "Follow up"
  })
  is_active = true
  negate    = false
}
```

`condition_params` is a `jsonencode(...)`-formatted JSON string on the Terraform side (type
`jsontypes.Normalized`, so key order and whitespace differences don't cause drift). CiviRules
itself stores this PHP `serialize()`-encoded and reads it back with `unserialize()`; this
resource converts JSON to and from that format automatically, so config only ever deals with
JSON.

## Argument Reference

The following arguments are supported:

### Required

- `rule_id` (Number) ID of the [`civicrm_civirules_rule`](civirules_rule.md) this condition belongs to.
- `condition_id` (Number) ID of the CiviRulesCondition type to apply. Look up via [`civicrm_civirules_condition`](../data-sources/civirules_condition.md) or the CiviRules UI.

### Optional

- `condition_params` (String) Parameters passed to the condition class, as a JSON string (e.g. `jsonencode(...)`). Structure depends on the condition type (e.g. `{"case_type_id": "3"}` for a case-type condition). Semantically compared — key order and whitespace don't affect drift.
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
