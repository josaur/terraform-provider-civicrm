---
page_title: "civicrm_navigation Resource - CiviCRM"
subcategory: ""
description: |-
  Manages a CiviCRM Navigation menu entry.
---

# civicrm_navigation (Resource)

Manages a CiviCRM Navigation menu entry. Navigation entries carry a `permission` list and a
`permission_operator`, which makes the menu the place where the interface is tailored per user
group: each group sees the entries its permissions allow, and nothing else. Entries can be
nested under one another via the self-referencing `parent_id`.

## Example Usage

```terraform
# Resolve an existing top-level menu entry to nest under
data "civicrm_navigation" "search" {
  name = "Search"
}

resource "civicrm_navigation" "open_invoices" {
  label      = "Open Invoices"
  name       = "open_invoices"
  url        = "civicrm/search#/display/open_invoices/open_invoices_table"
  parent_id  = data.civicrm_navigation.search.id
  permission = ["access CiviContribute"]
  weight     = 10
  is_active  = true
}

# Same search, different audience: a second entry with a different
# permission, so each group gets the entries relevant to it.
resource "civicrm_navigation" "invoices_with_sponsor" {
  label               = "Invoices with Sponsor"
  name                = "invoices_with_sponsor"
  url                 = "civicrm/search#/display/invoices_with_sponsor/invoices_with_sponsor_table"
  parent_id           = data.civicrm_navigation.search.id
  permission          = ["access CiviContribute", "administer CiviCRM data"]
  permission_operator = "AND"
  weight              = 11
}
```

`permission` is a list of permission name strings (e.g. `["access CiviCRM"]`). CiviCRM stores it
internally as a comma-joined string, but API v4 accepts and returns it as a list — pass a list
here, not a comma-joined string or `jsonencode(...)`. `permission_operator` (`"AND"` or `"OR"`)
only has an effect once two or more permissions are set.

## Argument Reference

### Optional

- `domain_id` (Number, Optional) FK to the domain this entry belongs to. Defaults to the current domain.
- `label` (String, Optional) Menu title shown in the UI.
- `name` (String, Optional) Internal machine name of the menu entry.
- `url` (String, Optional) Target URL, for custom links. Leave unset for entries that only group children.
- `icon` (String, Optional) CSS class of the icon shown next to the label.
- `permission` (List of String, Optional) Permissions required to see this menu entry.
- `permission_operator` (String, Optional) How multiple permission entries combine: `AND` or `OR`.
- `parent_id` (Number, Optional) FK to another `civicrm_navigation` entry this one is nested under.
- `is_active` (Boolean, Optional) Whether the entry is active. Default: `true`.
- `has_separator` (Number, Optional) Separator line around this entry: `0` = none, `1` = after, `2` = before. Default: `0`.
- `weight` (Number, Optional) Ordering of the entry among its siblings. If omitted, CiviCRM assigns the next free weight for the given parent.

## Attributes Reference

- `id` (Number) Unique ID of the navigation entry.

## Import

Navigation resources can be imported using the numeric ID:

```shell
terraform import civicrm_navigation.example 42
```

Existing entries — including CiviCRM's own built-in top-level items (`Search`, `Contributions`,
`Administer`, …) — can be resolved by name via the
[`civicrm_navigation` data source](../data-sources/navigation.md) and adopted with an `import`
block, following the same pattern documented for
[`civicrm_message_template`](message_template.md#adopting-civicrms-built-in-workflow-templates-via-import),
so a custom child entry can be nested under a standard parent without managing the whole tree.

```terraform
data "civicrm_navigation" "search" {
  name = "Search"
}

import {
  to = civicrm_navigation.search
  id = data.civicrm_navigation.search.id
}

resource "civicrm_navigation" "search" {
  label = data.civicrm_navigation.search.label
  name  = data.civicrm_navigation.search.name
}
```
