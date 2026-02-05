# Configure a site email address for outgoing mail
resource "civicrm_site_email_address" "default" {
  display_name = "Organization Name"
  email        = "info@example.org"
  description  = "Default sender address for all communications"
  is_default   = true
  is_active    = true
}
