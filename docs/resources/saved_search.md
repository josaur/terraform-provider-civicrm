---
page_title: "civicrm_saved_search Resource - CiviCRM"
subcategory: ""
description: |-
  Manages a CiviCRM SavedSearch.
---

# civicrm_saved_search (Resource)

Manages a CiviCRM SavedSearch.

## Example Usage

```terraform
resource "civicrm_saved_search" "example" {
  name = "example_saved_search"
}
```

### SearchKit search with a filter and a join

`api_params` holds the SearchKit/API4 query definition as a JSON string: `select`, `where`
(filters), and `join`. This example filters `Contact` to active Individuals and left-joins
`Email` to pull each contact's email address(es) — verified end-to-end against a local
`docker-compose` CiviCRM instance: applied via `terraform apply`, the resulting `api_params`
was extracted from CiviCRM and re-run directly through `Contact.get` with the same
`select`/`where`/`join`, confirming the filter and join both take effect (results correctly
limited to Individuals, correctly joined to `Email` — contacts with multiple emails yield one
result row per email, the expected join behavior), then removed again via `terraform destroy`.

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
join) is a normal in-place `update` — verified by changing the filter from `"Individual"` to
`"Organization"` and confirming the new filter is actually applied (not just round-tripped):
`SavedSearch.get` in a `cv ev` PHP context showed `api_params['where']` containing the updated
value, directly accessible as CiviCRM/SearchKit itself would read it.

`api_params` and `form_values` are `jsonencode(...)`-formatted JSON strings on the Terraform
side (type `jsontypes.Normalized`, so key order and whitespace differences don't cause drift),
but CiviCRM's API4 expects the *decoded* structure for both (`api_params` is `SERIALIZE_JSON`,
`form_values` is `SERIALIZE_PHP`, and API4 performs its own encoding either way) — the resource
decodes the JSON string before sending it to `create`/`update`, and re-encodes CiviCRM's
response back into a JSON string for state. An earlier version of this resource passed the
already-encoded string straight through, which caused CiviCRM to store a JSON array containing
a JSON *string* rather than the JSON *object* SearchKit expects, silently breaking the search:
`terraform plan` showed no drift (the read path decoded the same extra layer the write path
added), but `select`/`where` were not accessible from `api_params` in CiviCRM/SearchKit code,
since they weren't top-level keys. See
[GitHub issue #10](https://github.com/josaur/terraform-provider-civicrm/issues/10). Verified
fixed against the local `docker-compose` instance: `api_params` now stores as a plain JSON
object (confirmed via `SavedSearch.get` in a `cv ev` PHP context, with `select`/`where`
directly accessible as top-level keys, matching a search built through the CiviCRM UI), and
re-applying a config with a different `jsonencode(...)` key order than the stored value
produces no plan diff.

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
