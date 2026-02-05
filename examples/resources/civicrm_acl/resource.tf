# Grant edit access to a group
resource "civicrm_acl" "managers_edit_volunteers" {
  name         = "managers_edit_volunteers"
  entity_id    = civicrm_acl_role.volunteer_manager.id
  operation    = "Edit"
  object_table = "civicrm_group"
  object_id    = civicrm_group.volunteers.id
  is_active    = true
}
