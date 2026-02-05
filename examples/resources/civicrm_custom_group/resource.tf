# Create a custom field group for volunteer information
resource "civicrm_custom_group" "volunteer_info" {
  name    = "volunteer_info"
  title   = "Volunteer Information"
  extends = "Contact"
  style   = "Inline"
}
