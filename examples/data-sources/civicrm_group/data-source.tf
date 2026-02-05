# Look up a group by name
data "civicrm_group" "administrators" {
  name = "Administrators"
}

# Look up a group by ID
data "civicrm_group" "staff" {
  id = 5
}
