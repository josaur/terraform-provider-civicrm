---
page_title: "civicrm_contact_type Resource - CiviCRM"
subcategory: ""
description: |-
  Manages CiviCRM Contact Types and subtypes.
---

# civicrm_contact_type (Resource)

Manages CiviCRM Contact Types and subtypes. CiviCRM has three built-in contact types (Individual, Organization, Household), and you can create custom subtypes that inherit from these.

## Example Usage

```terraform
# Create a subtype of Individual
resource "civicrm_contact_type" "volunteer" {
  name        = "Volunteer"
  label       = "Volunteer"
  description = "Volunteer contacts"
  parent_id   = 1  # Individual
  icon        = "fa-hand-paper"
  is_active   = true
}

# Create a subtype of Organization
resource "civicrm_contact_type" "corporate_sponsor" {
  name        = "Corporate_Sponsor"
  label       = "Corporate Sponsor"
  description = "Corporate sponsoring organizations"
  parent_id   = 3  # Organization
  icon        = "fa-building"
  is_active   = true
}

# Create a subtype of Individual for staff
resource "civicrm_contact_type" "staff" {
  name        = "Staff"
  label       = "Staff Member"
  description = "Organization staff members"
  parent_id   = 1  # Individual
  icon        = "fa-id-badge"
  is_active   = true
}

# Create a subtype with image URL
resource "civicrm_contact_type" "donor" {
  name        = "Donor"
  label       = "Donor"
  description = "Financial donors"
  parent_id   = 1  # Individual
  image_url   = "https://example.org/images/donor-icon.png"
  is_active   = true
}
```

## Argument Reference

The following arguments are supported:

### Required

- `label` (String) The display label of the contact type.
- `name` (String) The machine name of the contact type (must be unique).

### Optional

- `description` (String) A description of the contact type.
- `icon` (String) FontAwesome icon class (e.g., `fa-user`, `fa-building`).
- `image_url` (String) URL to an image for this contact type.
- `is_active` (Boolean) Whether the contact type is active. Default: `true`.
- `is_reserved` (Boolean) Whether this is a reserved system contact type. Default: `false`.
- `parent_id` (Number) The parent contact type ID. Use `1` for Individual subtypes, `2` for Household subtypes, `3` for Organization subtypes.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (Number) The unique identifier of the contact type.

## Parent Type Reference

When creating subtypes, use these parent IDs:

| Parent Type   | ID |
|---------------|:--:|
| Individual    | 1  |
| Household     | 2  |
| Organization  | 3  |

## Import

Contact Types can be imported using the type ID:

```shell
terraform import civicrm_contact_type.example 123
```
