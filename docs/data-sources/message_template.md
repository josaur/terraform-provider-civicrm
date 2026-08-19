---
page_title: "civicrm_message_template Data Source - CiviCRM"
subcategory: ""
description: |-
  Fetches a CiviCRM MessageTemplate by ID or by workflow_name.
---

# civicrm_message_template (Data Source)

Fetches a CiviCRM MessageTemplate by `id`, or by `workflow_name` for one of CiviCRM's
built-in workflow templates (receipts, invoices, membership notifications, etc.). Specify
either `id`, or `workflow_name`.

Looking up by `workflow_name` is the mechanism used to resolve a built-in template's numeric
`id` at plan/apply time without hardcoding it — see
[the `civicrm_message_template` resource docs](../resources/message_template.md#adopting-civicrms-built-in-workflow-templates-via-import)
for the full import pattern.

## Example Usage

```terraform
# By id:
data "civicrm_message_template" "example" {
  id = 1
}

# By workflow_name (defaults to the editable, non-reserved copy):
data "civicrm_message_template" "invoice" {
  workflow_name = "contribution_invoice_receipt"
}
```

## Argument Reference

- `id` (Number, Optional) The unique identifier. Specify either `id`, or `workflow_name`.
- `workflow_name` (String, Optional) Look up a built-in workflow template by its
  `workflow_name` instead of by `id`. Specify either `id`, or `workflow_name`.
- `is_reserved` (Boolean, Optional) When looking up by `workflow_name`, whether to match the
  reserved (CiviCRM's unmodifiable original, `true`) or non-reserved (the editable copy
  actually used when sending mail, `false`) row. Defaults to `false`. Ignored when looking up
  by `id`.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

- `msg_title` (String) Descriptive title of message.
- `msg_subject` (String) Subject for email message..
- `msg_text` (String) Text formatted message.
- `msg_html` (String) HTML formatted message.
- `is_active` (Boolean) Is Active.
- `workflow_id` (Number) a pseudo-FK to civicrm_option_value.
- `is_default` (Boolean) is this the default message template for the workflow referenced by workflow_id?.
- `is_sms` (Boolean) Is this message template used for sms?.
- `pdf_format_id` (Number) a pseudo-FK to civicrm_option_value containing PDF Page Format..
