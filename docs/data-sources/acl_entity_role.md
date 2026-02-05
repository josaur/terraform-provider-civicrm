---
page_title: "civicrm_acl_entity_role Data Source - CiviCRM"
subcategory: ""
description: |-
  Fetches a CiviCRM ACL Entity Role assignment by ID or by role and entity combination.
---

# civicrm_acl_entity_role (Data Source)

Fetches a CiviCRM ACL Entity Role assignment by ID or by role and entity combination. Use this data source to look up existing ACL entity role assignments.

## Example Usage

```terraform
# Look up by ID
data "civicrm_acl_entity_role" "by_id" {
  id = 5
}

# Look up by role and entity combination
data "civicrm_acl_entity_role" "by_combination" {
  acl_role_id  = 10
  entity_id    = 20
  entity_table = "civicrm_group"
}

# Output assignment details
output "assignment_is_active" {
  value = data.civicrm_acl_entity_role.by_id.is_active
}

# Check if a specific group has a role assigned
data "civicrm_acl_entity_role" "admin_assignment" {
  acl_role_id = data.civicrm_acl_role.admin.id
  entity_id   = data.civicrm_group.administrators.id
}
```

## Argument Reference

The following arguments are supported. Either `id` or the combination of `acl_role_id` and `entity_id` must be specified.

- `acl_role_id` (Number, Optional) The ID of the ACL role. Use with `entity_id` to look up by combination.
- `entity_id` (Number, Optional) The ID of the entity. Use with `acl_role_id` to look up by combination.
- `entity_table` (String, Optional) The table containing the entity. Usually `civicrm_group`.
- `id` (Number, Optional) The unique identifier of the ACL entity role assignment.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

- `is_active` (Boolean) Whether this role assignment is active.
