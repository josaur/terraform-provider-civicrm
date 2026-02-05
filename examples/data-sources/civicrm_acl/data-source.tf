# Look up an ACL rule by name
data "civicrm_acl" "existing_rule" {
  name = "admin_edit_all"
}

# Look up an ACL rule by ID
data "civicrm_acl" "specific_rule" {
  id = 10
}
