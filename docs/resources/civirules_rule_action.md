---
page_title: "civicrm_civirules_rule_action Resource - CiviCRM"
subcategory: "CiviRules"
description: |-
  Attaches an action to a CiviRules rule.
---

# civicrm_civirules_rule_action (Resource)

Attaches an action to a CiviRules rule (entity: `CiviRulesRuleAction`). One rule can have multiple actions; all active actions are executed when the rule fires. The `action_id` references a CiviRulesAction type (look up via [`civicrm_civirules_action`](../data-sources/civirules_action.md)).

## Example Usage

```terraform
data "civicrm_civirules_action" "change_case_status" {
  name = "change_case_status"
}

resource "civicrm_civirules_rule_action" "close_on_complete" {
  rule_id   = civicrm_civirules_rule.my_rule.id
  action_id = data.civicrm_civirules_action.change_case_status.id
  action_params = jsonencode({
    case_status_id = 2
  })
  is_active = true
}
```

`action_params` is a `jsonencode(...)`-formatted JSON string on the Terraform side (type
`jsontypes.Normalized`, so key order and whitespace differences don't cause drift). CiviRules
itself stores this PHP `serialize()`-encoded and reads it back with `unserialize()`; this
resource converts JSON to and from that format automatically, so config only ever deals with
JSON.

## Argument Reference

The following arguments are supported:

### Required

- `rule_id` (Number) ID of the [`civicrm_civirules_rule`](civirules_rule.md) this action belongs to.
- `action_id` (Number) ID of the CiviRulesAction type to execute. Look up via [`civicrm_civirules_action`](../data-sources/civirules_action.md) or the CiviRules UI.

### Optional

- `action_params` (String) Parameters passed to the action class, as a JSON string (e.g. `jsonencode(...)`). Structure depends on the action type (e.g. `{"status_id": "2"}` for a change-case-status action). Semantically compared — key order and whitespace don't affect drift.
- `delay` (String) Delay before executing the action. ISO 8601 duration string (e.g. `"P1D"` = 1 day, `"PT2H"` = 2 hours). Leave empty for immediate execution.
- `is_active` (Boolean) Whether this action is active. Default: `true`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (Number) The unique identifier of the rule-action link.

## Import

CiviRules Rule Actions can be imported using the record ID:

```shell
terraform import civicrm_civirules_rule_action.example 1
```
