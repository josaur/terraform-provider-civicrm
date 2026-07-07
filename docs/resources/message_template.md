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
- `workflow_name` (String) For workflow templates: the machine name of the workflow this template overrides (e.g. `contribution_online_receipt`). Leave empty for user-defined templates.

## Attributes Reference

- `id` (Number) Unique ID of the message template.

## Import

Message Templates can be imported using the template ID:

```shell
terraform import civicrm_message_template.example 789
```
