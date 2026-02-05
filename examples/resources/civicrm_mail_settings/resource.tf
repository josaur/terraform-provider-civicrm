# Configure IMAP mail settings for bounce processing
resource "civicrm_mail_settings" "bounce_processing" {
  name       = "Bounce Processing"
  is_default = true
  domain     = "example.org"
  protocol   = "IMAP"
  server     = "mail.example.org"
  port       = 993
  username   = "bounce@example.org"
  password   = var.mail_password
  is_ssl     = true
  is_active  = true
}
