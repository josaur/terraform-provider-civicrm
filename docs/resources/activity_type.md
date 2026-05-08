---
page_title: "civicrm_activity_type Resource - CiviCRM"
subcategory: "Activities"
description: |-
  Manages a CiviCRM Activity Type.
---

# civicrm_activity_type (Resource)

Manages a CiviCRM Activity Type. Activity types are `OptionValue` records in the `activity_type` option group and define the kinds of activities that can be recorded against contacts (e.g. phone calls, meetings, assessments).

The `value` field is the internal numeric identifier CiviCRM uses to reference the activity type in API queries and CiviRules `action_params`.

## Example Usage

```terraform
resource "civicrm_activity_type" "beratungsgespraech" {
  name        = "beratungsgespraech"
  label       = "Beratungsgespräch"
  description = "Persönliches Beratungsgespräch mit dem Klienten"
  color       = "#3498db"
  is_active   = true
  weight      = 20
}

resource "civicrm_activity_type" "hausbesuch" {
  name   = "hausbesuch"
  label  = "Hausbesuch"
  color  = "#27ae60"
  icon   = "crm-i fa-home"
  weight = 30
}

# Use the value in a CiviRules action
resource "civicrm_civirules_rule_action" "create_followup" {
  rule_id   = civicrm_civirules_rule.my_rule.id
  action_id = data.civicrm_civirules_action.create_activity.id
  action_params = jsonencode({
    activity_type_id = civicrm_activity_type.beratungsgespraech.value
  })
  is_active = true
}
```

## Argument Reference

The following arguments are supported:

### Required

- `name` (String) Machine name of the activity type (e.g. `"beratungsgespraech"`). Must be unique within the `activity_type` option group.
- `label` (String) Display label shown in the UI (e.g. `"Beratungsgespräch"`).

### Optional

- `description` (String) Description of the activity type.
- `is_active` (Boolean) Whether the activity type is active. Default: `true`.
- `is_reserved` (Boolean) Whether the activity type is reserved (protected from deletion by CiviCRM). Default: `false`.
- `weight` (Number) Sort weight. Controls display order in dropdowns.
- `value` (String) Internal numeric value used by CiviCRM to identify this activity type. Auto-generated if not set.
- `color` (String) Hex color code for calendar and timeline display (e.g. `"#3498db"`).
- `icon` (String) CSS icon class (e.g. `"crm-i fa-phone"`).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (Number) The unique identifier of the OptionValue record.
- `value` (String) The internal value assigned by CiviCRM (computed if not specified). Reference this in CiviRules `action_params`.
- `weight` (Number) The sort weight (computed if not specified).

## Import

Activity Types can be imported using the OptionValue ID:

```shell
terraform import civicrm_activity_type.example 42
```
