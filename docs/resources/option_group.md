---
page_title: "civicrm_option_group Resource - CiviCRM"
subcategory: ""
description: |-
  Manages a CiviCRM OptionGroup.
---

# civicrm_option_group (Resource)

Manages a CiviCRM OptionGroup. An OptionGroup is a named container for OptionValues used to back Radio/Select/Multi-Select custom fields as well as many built-in CiviCRM enumerations (e.g. `gender`, `activity_type`). Pair with [`civicrm_option_value`](option_value.md) to declare the members.

## Example Usage

```terraform
resource "civicrm_option_group" "newsletter_consent" {
  name        = "newsletter_consent"
  title       = "Newsletter consent"
  description = "Values for the newsletter_consent custom field."
  data_type   = "String"
}

resource "civicrm_option_value" "consent_yes" {
  option_group_id = civicrm_option_group.newsletter_consent.id
  name            = "yes"
  label           = "Yes"
  value           = "yes"
  weight          = 10
}

resource "civicrm_option_value" "consent_no" {
  option_group_id = civicrm_option_group.newsletter_consent.id
  name            = "no"
  label           = "No"
  value           = "no"
  weight          = 20
}
```

## Argument Reference

### Required

- `name` (String) Machine name (must be unique, no spaces).
- `title` (String) Display title shown in the admin UI.

### Optional

- `description` (String) Optional description.
- `data_type` (String) Data type of the values. Common values: `String`, `Integer`. Default: `String`.
- `is_active` (Boolean) Whether the option group is active. Default: `true`.
- `is_reserved` (Boolean) Whether the option group is reserved (protected from deletion). Default: `false`.

## Attributes Reference

- `id` (Number) Unique ID of the option group.

## Import

Option Groups can be imported using the group ID:

```shell
terraform import civicrm_option_group.example 123
```
