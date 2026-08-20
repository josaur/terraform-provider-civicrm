---
page_title: "civicrm_search_display Data Source - CiviCRM"
subcategory: ""
description: |-
  Fetches a CiviCRM SearchKit SearchDisplay by ID or name.
---

# civicrm_search_display (Data Source)

Fetches a CiviCRM SearchKit SearchDisplay by `id`, or by `name` (optionally combined with
`saved_search_id` and/or `type` to disambiguate). Lets existing displays — including
CiviCRM's own built-in ones — be adopted via the `import` block pattern documented for
[`civicrm_message_template`](../resources/message_template.md#adopting-civicrms-built-in-workflow-templates-via-import),
instead of being recreated as duplicates.

## Example Usage

```terraform
data "civicrm_search_display" "example" {
  name = "Individuals_Table"
}
```

## Argument Reference

- `id` (Number, Optional) The unique identifier. Specify `id`, or `name`.
- `name` (String, Optional) Machine name of the search display. Specify `id`, or `name`.
- `saved_search_id` (Number, Optional) Filter by the saved search this display renders.
  Combine with `name` and/or `type` to disambiguate if `name` alone matches more than one
  display.
- `type` (String, Optional) Filter by display type (`table`, `list`, `grid`, `tree`,
  `autocomplete`, `entity`, `batch`). Combine with `name` and/or `saved_search_id` to
  disambiguate.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

- `label` (String) Administrative label for the display.
- `settings` (String) Display settings as a JSON string.
- `acl_bypass` (Boolean) Whether this display bypasses ACL permission checks.
- `is_autocomplete_default` (Boolean) Whether this is the default autocomplete display for
  its saved search's entity.
