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
