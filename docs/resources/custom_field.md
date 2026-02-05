---
page_title: "civicrm_custom_field Resource - CiviCRM"
subcategory: ""
description: |-
  Manages CiviCRM Custom Fields within Custom Groups.
---

# civicrm_custom_field (Resource)

Manages CiviCRM Custom Fields within Custom Groups. Custom fields allow you to extend CiviCRM entities with additional data storage.

## Example Usage

```terraform
# First, create a custom group
resource "civicrm_custom_group" "volunteer_info" {
  name    = "volunteer_info"
  title   = "Volunteer Information"
  extends = "Contact"
}

# Text field
resource "civicrm_custom_field" "skills" {
  custom_group_id = civicrm_custom_group.volunteer_info.id
  name            = "skills"
  label           = "Skills"
  data_type       = "String"
  html_type       = "Text"
  is_searchable   = true
}

# Date field
resource "civicrm_custom_field" "start_date" {
  custom_group_id  = civicrm_custom_group.volunteer_info.id
  name             = "volunteer_start_date"
  label            = "Volunteer Start Date"
  data_type        = "Date"
  html_type        = "Select Date"
  is_searchable    = true
  is_search_range  = true
  start_date_years = 10
  end_date_years   = 0
}

# Boolean checkbox
resource "civicrm_custom_field" "background_check" {
  custom_group_id = civicrm_custom_group.volunteer_info.id
  name            = "background_check_completed"
  label           = "Background Check Completed"
  data_type       = "Boolean"
  html_type       = "Radio"
  is_required     = true
}

# Multi-line text area
resource "civicrm_custom_field" "notes" {
  custom_group_id = civicrm_custom_group.volunteer_info.id
  name            = "volunteer_notes"
  label           = "Notes"
  data_type       = "Memo"
  html_type       = "TextArea"
  note_columns    = 80
  note_rows       = 6
}

# Money field
resource "civicrm_custom_field" "hourly_rate" {
  custom_group_id = civicrm_custom_group.volunteer_info.id
  name            = "hourly_rate"
  label           = "Hourly Rate"
  data_type       = "Money"
  html_type       = "Text"
  help_post       = "Enter the volunteer's hourly rate for reimbursement calculations."
}
```

## Argument Reference

The following arguments are supported:

### Required

- `custom_group_id` (Number) The ID of the custom group this field belongs to.
- `data_type` (String) The data type. Options: `String`, `Int`, `Float`, `Money`, `Memo`, `Date`, `Boolean`, `StateProvince`, `Country`, `File`, `Link`, `ContactReference`, `EntityReference`.
- `html_type` (String) The HTML input type. Options: `Text`, `TextArea`, `Select`, `Multi-Select`, `AdvMulti-Select`, `Radio`, `CheckBox`, `Select Date`, `Select State/Province`, `Select Country`, `File`, `Link`, `RichTextEditor`, `Autocomplete-Select`, `EntityRef`.
- `label` (String) The display label of the custom field.
- `name` (String) The machine name of the custom field (must be unique within the group).

### Optional

- `attributes` (String) Additional HTML attributes for the field.
- `column_name` (String) The database column name. Auto-generated if not specified.
- `date_format` (String) The date format string.
- `default_value` (String) The default value for the field.
- `end_date_years` (Number) Number of years after current date for date picker end.
- `filter` (String) Filter for entity reference fields.
- `fk_entity` (String) Foreign key entity for EntityReference fields.
- `fk_entity_on_delete` (String) Action on delete for foreign key. Options: `cascade`, `set_null`. Default: `set_null`.
- `help_post` (String) Help text displayed after the field.
- `help_pre` (String) Help text displayed before the field.
- `in_selector` (Boolean) Whether to include in selector. Default: `false`.
- `is_active` (Boolean) Whether the field is active. Default: `true`.
- `is_required` (Boolean) Whether the field is required. Default: `false`.
- `is_search_range` (Boolean) Whether to enable range search for this field. Default: `false`.
- `is_searchable` (Boolean) Whether the field is searchable. Default: `false`.
- `is_view` (Boolean) Whether the field is view-only. Default: `false`.
- `note_columns` (Number) Number of columns for note/textarea fields. Default: `60`.
- `note_rows` (Number) Number of rows for note/textarea fields. Default: `4`.
- `option_group_id` (Number) The ID of the option group for Select/Radio/CheckBox fields.
- `options_per_line` (Number) Number of options to display per line (for Radio/CheckBox).
- `serialize` (Number) Serialization method (0 for none, 1 for separator). Default: `0`.
- `start_date_years` (Number) Number of years before current date for date picker start.
- `text_length` (Number) Maximum text length for text fields. Default: `255`.
- `time_format` (Number) The time format (1 for 12-hour, 2 for 24-hour).
- `weight` (Number) The display order weight. Default: `1`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (Number) The unique identifier of the custom field.

## Import

Custom Fields can be imported using the field ID:

```shell
terraform import civicrm_custom_field.example 123
```
