# Assign the ACL role to a group
resource "civicrm_acl_entity_role" "team_leaders_as_managers" {
  acl_role_id  = civicrm_acl_role.volunteer_manager.id
  entity_table = "civicrm_group"
  entity_id    = civicrm_group.team_leaders.id
  is_active    = true
}
