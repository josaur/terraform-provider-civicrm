---
page_title: "civicrm_activity_type Data Source - CiviCRM"
subcategory: "Activities"
description: |-
  Fetches a CiviCRM Activity Type by ID or name.
---

# civicrm_activity_type (Data Source)

Fetches a CiviCRM Activity Type (an `OptionValue` in the `activity_type` option group) by ID or name. Useful to look up built-in activity types or reference the `value` field in CiviRules `action_params`.

## Example Usage

```terraform
data "civicrm_activity_type" "phone_call" {
  name = "Phone Call"
}

data "civicrm_activity_type" "meeting" {
  name = "Meeting"
}

# Use value in a CiviRules action
resource "civicrm_civirules_rule_action" "log_call" {
  rule_id   = civicrm_civirules_rule.my_rule.id
  action_id = data.civicrm_civirules_action.create_activity.id
  action_params = jsonencode({
    activity_type_id = data.civicrm_activity_type.phone_call.value
  })
  is_active = true
}
```

## Argument Reference

At least one of `id` or `name` must be specified.

- `id` (Number, Optional) Unique ID of the OptionValue record.
- `name` (String, Optional) Machine name of the activity type.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

- `label` (String) Display label shown in the UI.
- `description` (String) Description of the activity type.
- `is_active` (Boolean) Whether the activity type is active.
- `is_reserved` (Boolean) Whether the activity type is reserved by the system.
- `weight` (Number) Sort weight.
- `value` (String) Internal numeric value used by CiviCRM. Reference this in CiviRules `action_params`.
- `color` (String) Hex color code for calendar/timeline display.
- `icon` (String) CSS icon class.
