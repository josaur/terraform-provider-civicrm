---
page_title: "civicrm_site_email_address Resource - CiviCRM"
subcategory: ""
description: |-
  Manages CiviCRM Site Email Addresses used as sender addresses for outgoing emails.
---

# civicrm_site_email_address (Resource)

Manages CiviCRM Site Email Addresses used as sender addresses for outgoing emails. These addresses appear in the "From" field when sending emails from CiviCRM.

## Example Usage

```terraform
# Default organization email
resource "civicrm_site_email_address" "default" {
  display_name = "Organization Name"
  email        = "info@example.org"
  description  = "Default sender address for all communications"
  is_default   = true
  is_active    = true
}

# Newsletter sender
resource "civicrm_site_email_address" "newsletter" {
  display_name = "Organization Newsletter"
  email        = "newsletter@example.org"
  description  = "Sender address for newsletter mailings"
  is_active    = true
}

# Support team email
resource "civicrm_site_email_address" "support" {
  display_name = "Support Team"
  email        = "support@example.org"
  description  = "Support team communications"
  is_active    = true
}

# Donations department
resource "civicrm_site_email_address" "donations" {
  display_name = "Donations Department"
  email        = "donations@example.org"
  description  = "Donation receipts and thank you messages"
  is_active    = true
}
```

## Argument Reference

The following arguments are supported:

### Required

- `display_name` (String) The display name shown as the sender name (e.g., `CiviCRM Support`).
- `email` (String) The email address used as the sender address.

### Optional

- `description` (String) A description of this email address configuration.
- `domain_id` (Number) The domain ID this email address belongs to.
- `is_active` (Boolean) Whether this email address is active. Default: `true`.
- `is_default` (Boolean) Whether this is the default email address. Default: `false`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (Number) The unique identifier of the site email address.

## Import

Site Email Addresses can be imported using the address ID:

```shell
terraform import civicrm_site_email_address.example 123
```
