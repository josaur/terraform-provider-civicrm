---
page_title: "civicrm_mail_settings Resource - CiviCRM"
subcategory: ""
description: |-
  Manages CiviCRM Mail Settings for inbound email processing.
---

# civicrm_mail_settings (Resource)

Manages CiviCRM Mail Settings for inbound email processing. Mail settings configure how CiviCRM retrieves and processes incoming emails for bounce processing and email-to-activity conversion.

## Example Usage

```terraform
# IMAP mail settings for bounce processing
resource "civicrm_mail_settings" "bounce_processing" {
  name       = "Bounce Processing"
  is_default = true
  domain     = "example.org"
  localpart  = "bounce"
  protocol   = "IMAP"
  server     = "mail.example.org"
  port       = 993
  username   = "bounce@example.org"
  password   = var.mail_password
  is_ssl     = true
  is_active  = true
}

# POP3 mail settings
resource "civicrm_mail_settings" "support_inbox" {
  name       = "Support Inbox"
  is_default = false
  domain     = "example.org"
  protocol   = "POP3"
  server     = "pop.example.org"
  port       = 995
  username   = "support@example.org"
  password   = var.support_mail_password
  is_ssl     = true
  is_active  = true

  is_contact_creation_disabled_if_no_match = true
}

# Localdir settings (for mail processed by external MTA)
resource "civicrm_mail_settings" "local_mail" {
  name      = "Local Mail Processing"
  protocol  = "Localdir"
  source    = "/var/spool/civicrm/mail"
  is_active = true
}
```

## Argument Reference

The following arguments are supported:

### Required

- `name` (String) The name of this mail setting configuration.

### Optional

- `activity_assignees` (String) The activity assignees contact handling.
- `activity_source` (String) The activity source contact handling.
- `activity_status` (String) The default activity status for email activities.
- `activity_targets` (String) The activity targets contact handling.
- `activity_type_id` (Number) The activity type ID for email activities.
- `campaign_id` (Number) The campaign ID to associate with email activities.
- `domain` (String) The email domain (e.g., `example.org`).
- `domain_id` (Number) The domain ID this mail setting belongs to.
- `is_active` (Boolean) Whether this mail setting is active. Default: `true`.
- `is_contact_creation_disabled_if_no_match` (Boolean) Whether to disable contact creation if no match is found. Default: `false`.
- `is_default` (Boolean) Whether this is the default mail setting. Default: `false`.
- `is_non_case_email_skipped` (Boolean) Whether to skip emails not associated with a case. Default: `false`.
- `is_ssl` (Boolean) Whether to use SSL/TLS for the connection. Default: `false`.
- `localpart` (String) The local part prefix for bounce processing.
- `password` (String, Sensitive) The password for mail server authentication.
- `port` (Number) The mail server port.
- `protocol` (String) The mail protocol. Options: `IMAP`, `POP3`, `Maildir`, `Localdir`.
- `return_path` (String) The return path email address.
- `server` (String) The mail server hostname.
- `source` (String) The mail source (folder path for Maildir/Localdir).
- `username` (String) The username for mail server authentication.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (Number) The unique identifier of the mail settings.

## Import

Mail Settings can be imported using the settings ID:

```shell
terraform import civicrm_mail_settings.example 123
```
