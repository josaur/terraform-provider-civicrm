---
page_title: "civicrm_group Data Source - CiviCRM"
subcategory: ""
description: |-
  Fetches a CiviCRM Group by ID or name.
---

# civicrm_group (Data Source)

Fetches a CiviCRM Group by ID or name. Use this data source to look up existing groups to reference in your configuration.

## Example Usage

```terraform
# Look up a group by name
data "civicrm_group" "administrators" {
  name = "Administrators"
}

# Look up a group by ID
data "civicrm_group" "staff" {
  id = 5
}

# Use the data source to assign ACL permissions
resource "civicrm_acl_entity_role" "admin_privileges" {
  acl_role_id  = civicrm_acl_role.admin.id
  entity_table = "civicrm_group"
  entity_id    = data.civicrm_group.administrators.id
  is_active    = true
}

# Output group information
output "admin_group_title" {
  value = data.civicrm_group.administrators.title
}
```

## Argument Reference

The following arguments are supported. At least one of `id` or `name` must be specified.

- `id` (Number, Optional) The unique identifier of the group.
- `name` (String, Optional) The machine name of the group.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

- `description` (String) A description of the group.
- `is_active` (Boolean) Whether the group is active.
- `title` (String) The display title of the group.
- `visibility` (String) The visibility of the group.
