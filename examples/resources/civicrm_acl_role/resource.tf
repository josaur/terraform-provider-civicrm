# Create an ACL role for volunteer managers
resource "civicrm_acl_role" "volunteer_manager" {
  name        = "volunteer_manager"
  label       = "Volunteer Manager"
  description = "Can view and edit volunteer contacts"
  is_active   = true
}
