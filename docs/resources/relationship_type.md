---
page_title: "civicrm_relationship_type Resource - CiviCRM"
subcategory: ""
description: |-
  Manages CiviCRM Relationship Types that define how contacts can be related to each other.
---

# civicrm_relationship_type (Resource)

Manages CiviCRM Relationship Types that define how contacts can be related to each other. Relationship types are bidirectional with different labels for each direction (A to B and B to A).

## Example Usage

```terraform
# Parent/Child relationship
resource "civicrm_relationship_type" "parent_child" {
  name_a_b       = "Child of"
  label_a_b      = "Child of"
  name_b_a       = "Parent of"
  label_b_a      = "Parent of"
  description    = "Parent-child family relationship"
  contact_type_a = "Individual"
  contact_type_b = "Individual"
  is_active      = true
}

# Employee/Employer relationship
resource "civicrm_relationship_type" "employment" {
  name_a_b       = "Employee of"
  label_a_b      = "Employee of"
  name_b_a       = "Employer of"
  label_b_a      = "Employer of"
  description    = "Employment relationship"
  contact_type_a = "Individual"
  contact_type_b = "Organization"
  is_active      = true
}

# Volunteer relationship
resource "civicrm_relationship_type" "volunteer" {
  name_a_b       = "Volunteer for"
  label_a_b      = "Volunteer for"
  name_b_a       = "Has volunteer"
  label_b_a      = "Has volunteer"
  description    = "Volunteer relationship with organization"
  contact_type_a = "Individual"
  contact_type_b = "Organization"
  is_active      = true
}

# Mentor relationship (same type on both sides)
resource "civicrm_relationship_type" "mentor" {
  name_a_b       = "Mentee of"
  label_a_b      = "Mentee of"
  name_b_a       = "Mentor of"
  label_b_a      = "Mentor of"
  description    = "Mentorship relationship"
  contact_type_a = "Individual"
  contact_type_b = "Individual"
  is_active      = true
}

# Partner organizations relationship
resource "civicrm_relationship_type" "partner" {
  name_a_b       = "Partner of"
  label_a_b      = "Partner of"
  name_b_a       = "Partner of"
  label_b_a      = "Partner of"
  description    = "Partnership between organizations"
  contact_type_a = "Organization"
  contact_type_b = "Organization"
  is_active      = true
}
```

## Argument Reference

The following arguments are supported:

### Required

- `label_a_b` (String) The display label from A to B perspective.
- `label_b_a` (String) The display label from B to A perspective.
- `name_a_b` (String) The relationship name from A to B perspective (e.g., `Child of`).
- `name_b_a` (String) The relationship name from B to A perspective (e.g., `Parent of`).

### Optional

- `contact_sub_type_a` (String) The contact subtype for side A.
- `contact_sub_type_b` (String) The contact subtype for side B.
- `contact_type_a` (String) The contact type for side A. Options: `Individual`, `Organization`, `Household`. Leave empty for any type.
- `contact_type_b` (String) The contact type for side B. Options: `Individual`, `Organization`, `Household`. Leave empty for any type.
- `description` (String) A description of the relationship type.
- `is_active` (Boolean) Whether the relationship type is active. Default: `true`.
- `is_reserved` (Boolean) Whether this is a reserved system relationship type. Default: `false`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (Number) The unique identifier of the relationship type.

## Understanding A/B Relationships

CiviCRM relationships are bidirectional. When you create a relationship between two contacts:

- Contact A → Contact B uses `name_a_b` / `label_a_b`
- Contact B → Contact A uses `name_b_a` / `label_b_a`

For example, with a Parent/Child relationship:
- If John is contact A and Jane is contact B
- John is "Child of" Jane (A to B)
- Jane is "Parent of" John (B to A)

## Import

Relationship Types can be imported using the type ID:

```shell
terraform import civicrm_relationship_type.example 123
```
