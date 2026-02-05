---
page_title: "civicrm_acl_role Data Source - CiviCRM"
subcategory: ""
description: |-
  Fetches a CiviCRM ACL Role by ID or name.
---

# civicrm_acl_role (Data Source)

Fetches a CiviCRM ACL Role by ID or name. Use this data source to look up existing ACL roles to reference in your configuration.

## Example Usage

```terraform
# Look up an ACL role by name
data "civicrm_acl_role" "editor" {
  name = "editor"
}

# Look up an ACL role by ID
data "civicrm_acl_role" "viewer" {
  id = 2
}

# Use the data source to create an ACL rule
resource "civicrm_acl" "editor_access" {
  name         = "editor_access_volunteers"
  entity_id    = data.civicrm_acl_role.editor.id
  operation    = "Edit"
  object_table = "civicrm_group"
  object_id    = civicrm_group.volunteers.id
  is_active    = true
}

# Output role information
output "editor_role_label" {
  value = data.civicrm_acl_role.editor.label
}
```

## Argument Reference

The following arguments are supported. At least one of `id` or `name` must be specified.

- `id` (Number, Optional) The unique identifier of the ACL role.
- `name` (String, Optional) The machine name of the ACL role.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

- `description` (String) A description of the ACL role.
- `is_active` (Boolean) Whether the ACL role is active.
- `label` (String) The display label of the ACL role.
- `value` (String) The value of the ACL role (used internally by CiviCRM).
- `weight` (Number) The sort weight of the ACL role.
