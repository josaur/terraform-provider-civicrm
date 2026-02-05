# Look up an ACL role by name
data "civicrm_acl_role" "editor" {
  name = "editor"
}

# Look up an ACL role by ID
data "civicrm_acl_role" "viewer" {
  id = 2
}
