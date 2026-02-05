---
page_title: "civicrm_group Resource - CiviCRM"
subcategory: ""
description: |-
  Manages a CiviCRM Group. Groups are collections of contacts that can be used for ACL role assignments.
---

# civicrm_group (Resource)

Manages a CiviCRM Group. Groups are collections of contacts that can be used for ACL role assignments, mailing lists, and organizing contacts.

## Example Usage

```terraform
# Basic group
resource "civicrm_group" "volunteers" {
  name        = "volunteers"
  title       = "Volunteers"
  description = "Active volunteers in the organization"
  is_active   = true
  visibility  = "User and User Admin Only"
}

# Group with types
resource "civicrm_group" "newsletter_subscribers" {
  name        = "newsletter_subscribers"
  title       = "Newsletter Subscribers"
  description = "Contacts subscribed to the newsletter"
  is_active   = true
  visibility  = "Public Pages"
  group_type  = ["Mailing List"]
}

# Access Control group
resource "civicrm_group" "admin_staff" {
  name        = "admin_staff"
  title       = "Administrative Staff"
  description = "Staff with administrative access"
  is_active   = true
  group_type  = ["Access Control"]
}

# Nested group with parent
resource "civicrm_group" "west_region_volunteers" {
  name        = "west_region_volunteers"
  title       = "West Region Volunteers"
  description = "Volunteers in the western region"
  is_active   = true
  parents     = [civicrm_group.volunteers.id]
}
```

## Argument Reference

The following arguments are supported:

### Required

- `name` (String) The machine name of the group (must be unique).
- `title` (String) The display title of the group.

### Optional

- `description` (String) A description of the group.
- `frontend_description` (String) The public description of the group shown on frontend pages.
- `frontend_title` (String) The public title of the group shown on frontend pages.
- `group_type` (List of String) The types of the group. Valid values: `Access Control`, `Mailing List`.
- `is_active` (Boolean) Whether the group is active. Default: `true`.
- `is_hidden` (Boolean) Whether the group is hidden from the user interface. Default: `false`.
- `is_reserved` (Boolean) Whether the group is reserved (system group). Default: `false`.
- `parents` (List of Number) List of parent group IDs for nested groups.
- `visibility` (String) The visibility of the group. Options: `User and User Admin Only`, `Public Pages`. Default: `User and User Admin Only`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (Number) The unique identifier of the group.

## Import

Groups can be imported using the group ID:

```shell
terraform import civicrm_group.example 123
```
