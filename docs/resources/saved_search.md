---
page_title: "civicrm_saved_search Resource - CiviCRM"
subcategory: ""
description: |-
  Manages a CiviCRM SavedSearch.
---

# civicrm_saved_search (Resource)

Manages a CiviCRM SavedSearch. A saved search defines *what* to search for (entity, filters,
joins); it has no columns, labels, or sort order of its own — that's the job of one or more
[`civicrm_search_display`](search_display.md) resources pointing at it.

## Example Usage

```terraform
resource "civicrm_saved_search" "example" {
  name = "example_saved_search"
}
```

### SearchKit search with a filter and a join

`api_params` holds the SearchKit/API4 query definition as a JSON string: `select`, `where`
(filters), and `join`. This example filters `Contact` to active Individuals and left-joins
`Email` to pull each contact's email address(es):

```terraform
resource "civicrm_saved_search" "individuals_with_email" {
  name        = "individuals_with_email"
  label       = "Individuals with Email"
  api_entity  = "Contact"
  api_params = jsonencode({
    version = 4
    select = [
      "id",
      "display_name",
      "email_Email_01.email",
    ]
    join = [
      ["Email AS email_Email_01", "LEFT", ["id", "=", "email_Email_01.contact_id"]],
    ]
    where = [
      ["contact_type", "=", "Individual"],
      ["is_deleted", "=", false],
    ]
  })
}
```

Updating `where`/`join`/`select` (e.g. changing the `contact_type` filter, or adding another
join) is a normal in-place `update`.

`api_params` and `form_values` are `jsonencode(...)`-formatted JSON strings on the Terraform
side (type `jsontypes.Normalized`, so key order and whitespace differences don't cause drift).
CiviCRM's API4 expects the decoded structure for both, so this resource decodes the JSON string
before sending it to `create`/`update`, and re-encodes CiviCRM's response back into a JSON
string for state.

## Argument Reference

### Required

- `name` (String) Unique name of saved search.

### Optional

- `label` (String, Optional) Administrative label for search.
- `form_values` (String, Optional) Submitted form values for this search, as a JSON string (e.g. `jsonencode(...)`). Semantically compared — key order and whitespace don't affect drift.
- `mapping_id` (Number, Optional) Foreign key to civicrm_mapping used for saved search-builder searches..
- `search_custom_id` (Number, Optional) Foreign key to civicrm_option value table used for saved custom searches..
- `api_entity` (String, Optional) Entity name for API based search.
- `api_params` (String, Optional) Parameters for API based search (`select`/`where`/`join`/etc.), as a JSON string (e.g. `jsonencode(...)`). Semantically compared — key order and whitespace don't affect drift.
- `created_id` (Number, Optional) FK to contact table..
- `modified_id` (Number, Optional) FK to contact table..
- `expires_date` (String, Optional) Optional date after which the search is not needed.
- `created_date` (String, Optional) When the search was created..
- `modified_date` (String, Optional) When the search was last modified..
- `description` (String, Optional) Saved Search Description.
- `is_template` (Boolean, Optional) Search templates are used as a starting point for building new searches.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` (Number) The unique identifier of the saved_search.

## Import

SavedSearch resources can be imported using the numeric ID:

```shell
terraform import civicrm_saved_search.example 42
```
