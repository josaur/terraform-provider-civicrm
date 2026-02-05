# Create a relationship type for volunteers
resource "civicrm_relationship_type" "volunteer" {
  name_a_b       = "Volunteer for"
  label_a_b      = "Volunteer for"
  name_b_a       = "Has volunteer"
  label_b_a      = "Has volunteer"
  description    = "Volunteer relationship with organization"
  contact_type_a = "Individual"
  contact_type_b = "Organization"
  is_active      = true
}
