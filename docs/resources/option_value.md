---
page_title: "civicrm_option_value Resource - CiviCRM"
subcategory: ""
description: |-
  Manages a single OptionValue inside an OptionGroup.
---

# civicrm_option_value (Resource)

Manages a single OptionValue inside an OptionGroup. OptionValues back the choices of Radio/Select/Multi-Select custom fields (via `option_group_id` on [`civicrm_custom_field`](custom_field.md)) as well as many built-in CiviCRM enumerations.

For the `case_status` option group there is a dedicated [`civicrm_case_status`](case_status.md) resource with a fixed group.

## Example Usage

```terraform
resource "civicrm_option_group" "priority" {
  name  = "task_priority"
  title = "Task priority"
}

resource "civicrm_option_value" "priority_high" {
  option_group_id = civicrm_option_group.priority.id
  name            = "high"
  label           = "High"
  value           = "high"
  weight          = 10
  color           = "#dc3545"
}

resource "civicrm_option_value" "priority_low" {
  option_group_id = civicrm_option_group.priority.id
  name            = "low"
  label           = "Low"
  value           = "low"
  weight          = 20
  is_default      = true
}
```

## Argument Reference

### Required

- `option_group_id` (Number) ID of the parent OptionGroup.
- `label` (String) Display label shown to users.
- `value` (String) Stored value (what the custom field records). Must be unique within the group.

### Optional

- `name` (String) Machine name. Defaults to `value` if not set.
- `description` (String) Optional description.
- `weight` (Number) Sort weight. Controls display order.
- `is_active` (Boolean) Whether the option value is active. Default: `true`.
- `is_reserved` (Boolean) Whether the option value is reserved (protected from deletion). Default: `false`.
- `is_default` (Boolean) Whether this is the default option in the group. Default: `false`.
- `grouping` (String) Optional category bucket used by some CiviCRM contexts (e.g. `case_status` uses `Opened`/`Closed`).
- `color` (String) Optional color in hex format (e.g. `#ff0000`).
- `icon` (String) Optional icon (CSS class name).

## Attributes Reference

- `id` (Number) Unique ID of the option value.

## Import

Option Values can be imported using the value ID:

```shell
terraform import civicrm_option_value.example 456
```
