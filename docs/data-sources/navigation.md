---
page_title: "civicrm_navigation Data Source - CiviCRM"
subcategory: ""
description: |-
  Fetches a CiviCRM Navigation menu entry by ID or name.
---

# civicrm_navigation (Data Source)

Fetches a CiviCRM Navigation menu entry by `id` or by `name`. Lets an existing entry —
including CiviCRM's own built-in top-level items (e.g. `Search`, `Administer`) — be resolved by
name, for example as a `parent_id` for a new `civicrm_navigation` resource, or adopted via the
`import` block pattern documented for
[`civicrm_message_template`](../resources/message_template.md#adopting-civicrms-built-in-workflow-templates-via-import)
instead of being recreated as a duplicate.

## Example Usage

```terraform
data "civicrm_navigation" "search" {
  name = "Search"
}
```

## Argument Reference

- `id` (Number, Optional) The unique identifier. Specify `id`, or `name`.
- `name` (String, Optional) Internal machine name of the menu entry. Specify `id`, or `name`.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

- `label` (String) Menu title shown in the UI.
- `url` (String) Target URL, for custom links.
- `icon` (String) CSS class of the icon shown next to the label.
- `permission` (List of String) Permissions required to see this menu entry.
- `permission_operator` (String) How multiple permission entries combine: `AND` or `OR`.
- `parent_id` (Number) FK to the parent `civicrm_navigation` entry, if nested.
- `is_active` (Boolean) Whether the entry is active.
- `has_separator` (Number) Separator line around this entry: `0` = none, `1` = after, `2` = before.
- `weight` (Number) Ordering of the entry among its siblings.
