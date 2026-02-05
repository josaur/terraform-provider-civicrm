---
page_title: "civicrm_custom_group Resource - CiviCRM"
subcategory: ""
description: |-
  Manages CiviCRM Custom Field Groups. Custom groups organize custom fields that extend CiviCRM entities.
---

# civicrm_custom_group (Resource)

Manages CiviCRM Custom Field Groups. Custom groups organize custom fields that extend CiviCRM entities like Contacts, Activities, Contributions, and more.

## Example Usage

```terraform
# Basic custom group for all contacts
resource "civicrm_custom_group" "volunteer_info" {
  name    = "volunteer_info"
  title   = "Volunteer Information"
  extends = "Contact"
  style   = "Inline"
}

# Custom group for organizations only
resource "civicrm_custom_group" "company_details" {
  name    = "company_details"
  title   = "Company Details"
  extends = "Organization"
  style   = "Tab"
}

# Custom group with help text
resource "civicrm_custom_group" "emergency_contact" {
  name      = "emergency_contact"
  title     = "Emergency Contact"
  extends   = "Individual"
  style     = "Inline"
  help_pre  = "Please provide emergency contact information."
  help_post = "This information will only be used in case of emergency."
}

# Multi-record custom group
resource "civicrm_custom_group" "employment_history" {
  name         = "employment_history"
  title        = "Employment History"
  extends      = "Individual"
  style        = "Tab with table"
  is_multiple  = true
  max_multiple = 10
}

# Custom group for activities
resource "civicrm_custom_group" "activity_details" {
  name             = "activity_details"
  title            = "Activity Details"
  extends          = "Activity"
  collapse_display = true
}
```

## Argument Reference

The following arguments are supported:

### Required

- `extends` (String) The entity type this custom group extends. Examples: `Contact`, `Individual`, `Organization`, `Household`, `Activity`, `Contribution`, `Membership`, `Participant`, `Event`, `Relationship`.
- `name` (String) The machine name of the custom group (must be unique).
- `title` (String) The display title of the custom group.

### Optional

- `collapse_adv_display` (Boolean) Whether to collapse in advanced search display. Default: `true`.
- `collapse_display` (Boolean) Whether to collapse the group display by default. Default: `false`.
- `extends_entity_column_id` (Number) For extending specific subtypes, the column ID.
- `extends_entity_column_value` (List of String) For extending specific subtypes, the allowed values.
- `help_post` (String) Help text displayed after the custom fields.
- `help_pre` (String) Help text displayed before the custom fields.
- `icon` (String) The icon for the custom group (CSS class name).
- `is_active` (Boolean) Whether the custom group is active. Default: `true`.
- `is_multiple` (Boolean) Whether multiple records can be stored per entity. Default: `false`.
- `is_public` (Boolean) Whether this group is visible on public forms. Default: `true`.
- `is_reserved` (Boolean) Whether this is a reserved system group. Default: `false`.
- `max_multiple` (Number) Maximum number of multiple records (if `is_multiple` is `true`).
- `min_multiple` (Number) Minimum number of multiple records (if `is_multiple` is `true`).
- `style` (String) The display style. Options: `Inline`, `Tab`, `Tab with table`. Default: `Inline`.
- `table_name` (String) The database table name for storing custom field values. Auto-generated if not specified.
- `weight` (Number) The display order weight. Default: `1`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (Number) The unique identifier of the custom group.

## Import

Custom Groups can be imported using the group ID:

```shell
terraform import civicrm_custom_group.example 123
```
