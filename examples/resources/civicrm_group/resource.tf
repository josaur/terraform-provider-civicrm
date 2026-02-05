# Create a group for volunteers
resource "civicrm_group" "volunteers" {
  name        = "volunteers"
  title       = "Volunteers"
  description = "Active volunteers in the organization"
  is_active   = true
  visibility  = "User and User Admin Only"
  group_type  = ["Access Control"]
}
