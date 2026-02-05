---
page_title: "civicrm_tag Resource - CiviCRM"
subcategory: ""
description: |-
  Manages CiviCRM Tags for categorizing contacts and other entities.
---

# civicrm_tag (Resource)

Manages CiviCRM Tags for categorizing contacts and other entities. Tags can be organized hierarchically and assigned colors for visual distinction.

## Example Usage

```terraform
# Simple tag
resource "civicrm_tag" "volunteer" {
  name  = "Volunteer"
  label = "Volunteer"
  color = "#28a745"
}

# Tag for specific entity types
resource "civicrm_tag" "priority_donor" {
  name        = "priority_donor"
  label       = "Priority Donor"
  description = "High-value donors requiring special attention"
  color       = "#ffc107"
  used_for    = ["civicrm_contact"]
}

# Tagset (container for other tags)
resource "civicrm_tag" "skills" {
  name        = "skills"
  label       = "Skills"
  description = "Skills tagset for volunteers"
  is_tagset   = true
}

# Child tag under a tagset
resource "civicrm_tag" "skill_accounting" {
  name        = "skill_accounting"
  label       = "Accounting"
  parent_id   = civicrm_tag.skills.id
  color       = "#17a2b8"
}

# Reserved system tag
resource "civicrm_tag" "system_import" {
  name        = "system_import"
  label       = "Imported Contact"
  is_reserved = true
  used_for    = ["civicrm_contact"]
}
```

## Argument Reference

The following arguments are supported:

### Required

- `name` (String) The machine name of the tag (must be unique, no spaces).

### Optional

- `color` (String) The color for the tag in hex format (e.g., `#ff0000`).
- `description` (String) A description of the tag.
- `is_reserved` (Boolean) Whether this is a reserved system tag. Default: `false`.
- `is_selectable` (Boolean) Whether this tag can be selected. Default: `true`.
- `is_tagset` (Boolean) Whether this is a tagset (container for other tags). Default: `false`.
- `label` (String) The display label of the tag. Defaults to the `name` if not specified.
- `parent_id` (Number) The parent tag ID for hierarchical tags.
- `used_for` (List of String) Entity types this tag can be used for (e.g., `civicrm_contact`, `civicrm_activity`).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (Number) The unique identifier of the tag.

## Import

Tags can be imported using the tag ID:

```shell
terraform import civicrm_tag.example 123
```
