---
page_title: "civicrm_message_template Resource - CiviCRM"
subcategory: ""
description: |-
  Manages a CiviCRM MessageTemplate.
---

# civicrm_message_template (Resource)

Manages a CiviCRM MessageTemplate. MessageTemplates hold the HTML and text bodies used by CiviMail, transactional mail (`Email.send` with `template_id`), and workflow notifications. Use for user-defined templates or to override the default workflow templates.

## Example Usage

```terraform
# User-defined transactional template
resource "civicrm_message_template" "invitation" {
  msg_title   = "event_invitation"
  msg_subject = "You're invited to {event.title}"
  msg_html    = file("${path.module}/templates/event_invitation.html")
  msg_text    = file("${path.module}/templates/event_invitation.txt")
}

# Workflow override (replaces the built-in contribution receipt)
resource "civicrm_message_template" "contribution_receipt_override" {
  msg_title     = "Contribution Online Receipt (custom)"
  workflow_name = "contribution_online_receipt"
  msg_subject   = "Receipt for your contribution"
  msg_html      = file("${path.module}/templates/receipt.html")
  msg_text      = file("${path.module}/templates/receipt.txt")
  is_default    = true
}
```

## Argument Reference

### Required

- `msg_title` (String) Internal title of the template (used to look it up by name).

### Optional

- `msg_subject` (String) Subject line for outgoing mails.
- `msg_html` (String) HTML body of the message.
- `msg_text` (String) Plain-text body of the message.
- `is_active` (Boolean) Whether the template is active. Default: `true`.
- `is_reserved` (Boolean) Whether the template is reserved (system-managed). Default: `false`.
- `is_default` (Boolean) For workflow templates: whether this is the default version. Default: `false`.
- `is_sms` (Boolean) Whether this message template is used for SMS. Default: `false`.
- `workflow_name` (String) For workflow templates: the machine name of the workflow this template overrides (e.g. `contribution_online_receipt`). Leave empty for user-defined templates.
- `workflow_id` (Number) Pseudo-FK to `civicrm_option_value` for the workflow this template belongs to. Usually left computed; set `workflow_name` instead.
- `pdf_format_id` (Number) Pseudo-FK to `civicrm_option_value` for the PDF page format used when rendering this template (e.g. invoices).

## Attributes Reference

- `id` (Number) Unique ID of the message template.

## Import

Message Templates can be imported using the template ID:

```shell
terraform import civicrm_message_template.example 789
```

### Adopting CiviCRM's built-in workflow templates via import

CiviCRM ships a set of built-in workflow templates out of the box — receipts, invoices,
membership notifications, password resets, etc. These already exist in every install and
**must be imported, not created** — `create` always inserts a new row, so applying a
`civicrm_message_template` config with, say, `workflow_name = "contribution_invoice_receipt"`
and no prior import produces a second, unused "Contributions - Invoice" template rather than
managing the one CiviCRM actually sends.

Each workflow (identified by `workflow_name`) has **two** rows in CiviCRM:

- `is_reserved = true` — CiviCRM's unmodifiable original, used as a fallback/reset target.
  Do not manage this one with Terraform.
- `is_reserved = false`, `is_default = true` — the editable copy that is actually used when
  sending mail. This is the one to import.

#### Import a single template, with the id resolved at plan time

Use the [`civicrm_message_template` data source](../data-sources/message_template.md) to look
up the current numeric id by `workflow_name` — no id needs to be hardcoded, so the same config
works unmodified across CiviCRM instances where the underlying id differs. Verified end-to-end
against a local `docker-compose` CiviCRM instance (`terraform init` → `apply` → `plan` showing
no drift):

```terraform
data "civicrm_message_template" "invoice_lookup" {
  workflow_name = "contribution_invoice_receipt"
}

import {
  to = civicrm_message_template.invoice
  id = data.civicrm_message_template.invoice_lookup.id
}

resource "civicrm_message_template" "invoice" {
  msg_title     = data.civicrm_message_template.invoice_lookup.msg_title
  workflow_name = "contribution_invoice_receipt"
  msg_subject   = data.civicrm_message_template.invoice_lookup.msg_subject
  msg_html      = data.civicrm_message_template.invoice_lookup.msg_html
  msg_text      = data.civicrm_message_template.invoice_lookup.msg_text
}
```

`terraform apply` imports the existing template into state (no duplicate is created), seeded
with its current content via the data source so the first plan shows no unintended diff. From
then on it's a normal resource: edit `msg_html`/`msg_subject`/etc. in the config as usual.

#### Import all standard templates at once

The same pattern generalizes to every built-in workflow with `for_each`. `workflow_name`
values are stable CiviCRM core constants (not instance-specific data), so the list itself can
be hardcoded even though no numeric id ever is. Verified against the same local instance with
2 of the 32 workflows (`contribution_invoice_receipt`, `contribution_offline_receipt`) —
`for_each` over the full set follows the identical mechanism:

```terraform
locals {
  # CiviCRM's built-in workflow template names (core constants, stable across instances).
  standard_message_template_workflows = [
    "case_activity",
    "contribution_dupalert",
    "contribution_invoice_receipt",
    "contribution_offline_receipt",
    "contribution_online_receipt",
    "contribution_recurring_billing",
    "contribution_recurring_cancelled",
    "contribution_recurring_edit",
    "contribution_recurring_notify",
    "event_offline_receipt",
    "event_online_receipt",
    "friend",
    "membership_autorenew_billing",
    "membership_autorenew_cancelled",
    "membership_offline_receipt",
    "membership_online_receipt",
    "participant_cancelled",
    "participant_confirm",
    "participant_expired",
    "participant_transferred",
    "password_reset",
    "payment_or_refund_notification",
    "pcp_notify",
    "pcp_owner_notify",
    "pcp_status_change",
    "pcp_supporter_notify",
    "petition_confirmation_needed",
    "petition_sign",
    "pledge_acknowledge",
    "pledge_reminder",
    "test_preview",
    "uf_notify",
  ]
}

data "civicrm_message_template" "standard" {
  for_each      = toset(local.standard_message_template_workflows)
  workflow_name = each.value
}

import {
  for_each = data.civicrm_message_template.standard
  to       = civicrm_message_template.standard[each.key]
  id       = each.value.id
}

resource "civicrm_message_template" "standard" {
  for_each      = data.civicrm_message_template.standard
  msg_title     = each.value.msg_title
  workflow_name = each.key
  msg_subject   = each.value.msg_subject
  msg_html      = each.value.msg_html
  msg_text      = each.value.msg_text
}
```

`workflow_name` values are core CiviCRM constants that may not exist on every instance
(depending on installed components — e.g. `event_offline_receipt` needs CiviEvent, `pledge_*`
needs CiviPledge). Trim the list to the components actually enabled on your instance, or the
data source lookup fails for a workflow with no matching template.

`terraform apply` imports all matching templates in one pass, each into its own
`civicrm_message_template.standard["<workflow_name>"]` state entry — no ids hardcoded anywhere,
and no duplicates created for templates that already exist.
