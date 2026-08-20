---
page_title: "civicrm_civirules_rule_action Resource - CiviCRM"
subcategory: "CiviRules"
description: |-
  Attaches an action to a CiviRules rule.
---

# civicrm_civirules_rule_action (Resource)

Attaches an action to a CiviRules rule (entity: `CiviRulesRuleAction`). One rule can have multiple actions; all active actions are executed when the rule fires. The `action_id` references a CiviRulesAction type (look up via [`civicrm_civirules_action`](../data-sources/civirules_action.md)).

## Example Usage

CiviRules stores `action_params` via `serialize()` and reads them with
`unserialize()` (see `CRM/Civirules/Action/*Form.php`). Values written
as JSON are silently accepted at write time but the action fails
(silently or with an error) at runtime because `unserialize()` cannot
decode them. Pass a PHP-serialized string.

```terraform
data "civicrm_civirules_action" "change_case_status" {
  name = "change_case_status"
}

resource "civicrm_civirules_rule_action" "close_on_complete" {
  rule_id   = civicrm_civirules_rule.my_rule.id
  action_id = data.civicrm_civirules_action.change_case_status.id
  # a:1:{s:14:"case_status_id";i:2;}
  action_params = "a:1:{s:14:\"case_status_id\";i:2;}"
  is_active = true
}
```

## Argument Reference

The following arguments are supported:

### Required

- `rule_id` (Number) ID of the [`civicrm_civirules_rule`](civirules_rule.md) this action belongs to.
- `action_id` (Number) ID of the CiviRulesAction type to execute. Look up via [`civicrm_civirules_action`](../data-sources/civirules_action.md) or the CiviRules UI.

### Optional

- `action_params` (String) PHP `serialize()`-encoded parameters passed to the action class. Structure depends on the action type (e.g. `a:1:{s:9:"status_id";s:1:"2";}` for a change-case-status action). See note in *Example Usage* above; JSON is silently accepted at write time but breaks the action at runtime.
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
