# Create a custom field for skills
resource "civicrm_custom_field" "skills" {
  custom_group_id = civicrm_custom_group.volunteer_info.id
  name            = "skills"
  label           = "Skills"
  data_type       = "String"
  html_type       = "Text"
  is_searchable   = true
}
