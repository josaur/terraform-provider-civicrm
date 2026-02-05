---
page_title: "civicrm_acl Data Source - CiviCRM"
subcategory: ""
description: |-
  Fetches a CiviCRM ACL rule by ID or name.
---

# civicrm_acl (Data Source)

Fetches a CiviCRM ACL rule by ID or name. Use this data source to look up existing ACL rules to reference in your configuration.

## Example Usage

```terraform
# Look up an ACL rule by name
data "civicrm_acl" "existing_rule" {
  name = "admin_edit_all"
}

# Look up an ACL rule by ID
data "civicrm_acl" "specific_rule" {
  id = 10
}

# Output ACL rule details
output "acl_operation" {
  value = data.civicrm_acl.existing_rule.operation
}

output "acl_object_id" {
  value = data.civicrm_acl.existing_rule.object_id
}
```

## Argument Reference

The following arguments are supported. At least one of `id` or `name` must be specified.

- `id` (Number, Optional) The unique identifier of the ACL rule.
- `name` (String, Optional) The name of the ACL rule.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

- `deny` (Boolean) Whether this ACL denies rather than allows access.
- `entity_id` (Number) The ID of the ACL role this rule belongs to.
- `entity_table` (String) The entity table that owns this ACL.
- `is_active` (Boolean) Whether the ACL rule is active.
- `object_id` (Number) The ID of the specific object being permissioned.
- `object_table` (String) The type of object being permissioned.
- `operation` (String) The operation this ACL grants.
- `priority` (Number) The priority of the ACL rule.
