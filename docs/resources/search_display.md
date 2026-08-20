---
page_title: "civicrm_search_display Resource - CiviCRM"
subcategory: ""
description: |-
  Manages a CiviCRM SearchKit SearchDisplay.
---

# civicrm_search_display (Resource)

Manages a CiviCRM SearchKit SearchDisplay — the part of a search that decides which columns
are shown, what they are called, sort order, pager behavior, row links and action buttons.
`civicrm_saved_search` defines *what* to search for; `civicrm_search_display` defines how to
render it. A search without a display has no column labels and cannot be used as a screen.

This resource covers all seven SearchKit display types — they differ only in the shape of
`settings`, stored as a JSON string:

| `type` | UI label |
|---|---|
| `table` | Table |
| `list` | List |
| `grid` | Grid |
| `tree` | Tree |
| `autocomplete` | Autocomplete |
| `entity` | DB Entity |
| `batch` | Data Entry |

## Example Usage

A table display with labeled, sortable columns:

```terraform
resource "civicrm_saved_search" "individuals" {
  name       = "individuals_with_email"
  label      = "Individuals with Email"
  api_entity = "Contact"
  api_params = jsonencode({
    version = 4
    select  = ["id", "display_name", "contact_type"]
    where   = [["contact_type", "=", "Individual"]]
  })
}

resource "civicrm_search_display" "individuals_table" {
  label           = "Individuals Table"
  saved_search_id = civicrm_saved_search.individuals.id
  type            = "table"
  acl_bypass      = false

  settings = jsonencode({
    limit = 25
    sort  = [["display_name", "ASC"]]
    pager = { hide_single = true }
    columns = [
      { type = "field", key = "id", label = "ID", sortable = true },
      { type = "field", key = "display_name", label = "Name", sortable = true },
      { type = "field", key = "contact_type", label = "Kontakttyp", sortable = false },
    ]
  })
}
```

### All seven display types

Every display type accepts the same resource shape; only `settings` differs.

```terraform
resource "civicrm_search_display" "list_example" {
  label            = "Individuals List"
  saved_search_id  = civicrm_saved_search.individuals.id
  type             = "list"
  settings = jsonencode({
    title = { text = "[display_name]" }
  })
}

resource "civicrm_search_display" "grid_example" {
  label            = "Individuals Grid"
  saved_search_id  = civicrm_saved_search.individuals.id
  type             = "grid"
  settings = jsonencode({
    title = { text = "[display_name]" }
  })
}

resource "civicrm_search_display" "tree_example" {
  label            = "Individuals Tree"
  saved_search_id  = civicrm_saved_search.individuals.id
  type             = "tree"
  settings = jsonencode({
    label = { text = "[display_name]" }
  })
}

resource "civicrm_search_display" "autocomplete_example" {
  label                    = "Individuals Autocomplete"
  saved_search_id          = civicrm_saved_search.individuals.id
  type                     = "autocomplete"
  is_autocomplete_default  = false
  settings = jsonencode({
    label = { text = "[display_name]" }
  })
}

# "entity" registers a virtual DB entity named after this display, so it
# requires an explicit `name` — CiviCRM does not derive one from `label` as
# it does for every other type.
resource "civicrm_search_display" "entity_example" {
  name             = "individuals_entity"
  label            = "Individuals Entity"
  saved_search_id  = civicrm_saved_search.individuals.id
  type             = "entity"
  settings = jsonencode({
    columns = [
      { type = "field", key = "id" },
      { type = "field", key = "display_name" },
    ]
  })
}

resource "civicrm_search_display" "batch_example" {
  label            = "Individuals Batch"
  saved_search_id  = civicrm_saved_search.individuals.id
  type             = "batch"
  settings = jsonencode({
    columns = [
      { type = "field", key = "id" },
      { type = "field", key = "display_name" },
    ]
  })
}
```

## Argument Reference

### Required

- `label` (String) Administrative label for the display.
- `saved_search_id` (Number) ID of the `civicrm_saved_search` this display renders.
- `type` (String) Display type. One of: `table`, `list`, `grid`, `tree`, `autocomplete`, `entity`, `batch`.

### Optional

- `name` (String, Optional) Machine name of the search display. If omitted, CiviCRM assigns
  one derived from `label` — **except** for `type = "entity"`, where CiviCRM registers a
  virtual DB entity named after the display and therefore requires an explicit `name`.
- `settings` (String, Optional) Display settings (columns, sort, limit, pager, row links,
  action buttons, etc.) as a JSON string (e.g. from `jsonencode(...)`). The exact shape depends
  on `type` — see the examples above and the [CiviCRM SearchKit documentation](https://docs.civicrm.org/user/en/latest/searching/search-kit/).
  CiviCRM augments each `columns[]` entry with a generated `spec` sub-object server-side for
  some display types (`batch` and `entity`); this attribute strips `spec` back out when reading
  state so `plan` stays clean, while `spec` remains fully present in CiviCRM itself for
  SearchKit's own use.
- `acl_bypass` (Boolean, Optional) Whether this display bypasses ACL permission checks. Default: `false`.
- `is_autocomplete_default` (Boolean, Optional) Whether this is the default autocomplete
  display for its saved search's entity. Default: `false`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (Number) The unique identifier of the search display.

## Import

SearchDisplay resources can be imported using the numeric ID:

```shell
terraform import civicrm_search_display.example 42
```
