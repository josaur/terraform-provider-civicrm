---
page_title: "civicrm_saved_search Data Source - CiviCRM"
subcategory: ""
description: |-
  Fetches a CiviCRM SavedSearch by ID or name.
---

# civicrm_saved_search (Data Source)

Fetches a CiviCRM SavedSearch by ID or name. At least one of `id` or `name` must be specified.

## Example Usage

```terraform
# Look up by name
data "civicrm_saved_search" "example" {
  name = "example"
}

# Look up by ID
data "civicrm_saved_search" "by_id" {
  id = 1
}
```

## Argument Reference

- `id` (Number, Optional) The unique identifier.
- `name` (String, Optional) The name of the saved_search.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

- `label` (String) Administrative label for search.
- `form_values` (String) Submitted form values for this search.
- `mapping_id` (Number) Foreign key to civicrm_mapping used for saved search-builder searches..
- `search_custom_id` (Number) Foreign key to civicrm_option value table used for saved custom searches..
- `api_entity` (String) Entity name for API based search.
- `api_params` (String) Parameters for API based search.
- `created_id` (Number) FK to contact table..
- `modified_id` (Number) FK to contact table..
- `expires_date` (String) Optional date after which the search is not needed.
- `created_date` (String) When the search was created..
- `modified_date` (String) When the search was last modified..
- `description` (String) Saved Search Description.
- `is_template` (Boolean) Search templates are used as a starting point for building new searches.
