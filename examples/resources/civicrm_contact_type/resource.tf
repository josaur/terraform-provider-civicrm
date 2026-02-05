# Create a custom contact type for volunteers
resource "civicrm_contact_type" "volunteer" {
  name        = "Volunteer"
  label       = "Volunteer"
  description = "Volunteer contacts"
  parent_id   = 1 # Individual
  icon        = "fa-hand-paper"
  is_active   = true
}
