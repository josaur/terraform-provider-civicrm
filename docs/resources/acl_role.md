---
page_title: "civicrm_acl_role Resource - CiviCRM"
subcategory: ""
description: |-
  Manages a CiviCRM ACL Role. ACL Roles define permission sets that can be assigned to groups.
---

# civicrm_acl_role (Resource)

Manages a CiviCRM ACL Role. ACL Roles define permission sets that can be assigned to groups of contacts through ACL Entity Roles.

## Example Usage

```terraform
# Basic ACL role
resource "civicrm_acl_role" "volunteer_manager" {
  name        = "volunteer_manager"
  label       = "Volunteer Manager"
  description = "Can view and edit volunteer contacts"
  is_active   = true
}

# Viewer role
resource "civicrm_acl_role" "data_viewer" {
  name        = "data_viewer"
  label       = "Data Viewer"
  description = "Read-only access to contact data"
  is_active   = true
}

# Combined with ACL rules
resource "civicrm_acl_role" "event_coordinator" {
  name        = "event_coordinator"
  label       = "Event Coordinator"
  description = "Manages event participants"
  is_active   = true
}

resource "civicrm_acl" "event_coordinator_edit" {
  name         = "event_coordinator_edit_participants"
  entity_id    = civicrm_acl_role.event_coordinator.id
  operation    = "Edit"
  object_table = "civicrm_group"
  object_id    = civicrm_group.event_participants.id
  is_active    = true
}
```

## Argument Reference

The following arguments are supported:

### Required

- `label` (String) The display label of the ACL role.
- `name` (String) The machine name of the ACL role (must be unique).

### Optional

- `description` (String) A description of the ACL role.
- `is_active` (Boolean) Whether the ACL role is active. Default: `true`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (Number) The unique identifier of the ACL role.

## Import

ACL Roles can be imported using the role ID:

```shell
terraform import civicrm_acl_role.example 123
```
