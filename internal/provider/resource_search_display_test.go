package provider_test

import (
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// checkSearchDisplaySettingsIsObject asserts that SearchDisplay.settings, as
// returned by SearchDisplay.get, is a JSON object (map[string]any) with the
// configured columns directly accessible — the shape CiviCRM/SearchKit
// actually reads, and the same class of check as checkAPIParamsIsObject in
// resource_saved_search_test.go for the analogous bug on
// SavedSearch.api_params (issue #10). A round-trip through this provider's
// own Read path cannot distinguish this from a double-JSON-encoded value.
func checkSearchDisplaySettingsIsObject(t *testing.T, resourceAddr string, wantColumnKeys []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		t.Helper()
		rs, ok := s.RootModule().Resources[resourceAddr]
		if !ok {
			t.Fatalf("resource %q not found in state", resourceAddr)
		}
		id, err := strconv.ParseInt(rs.Primary.Attributes["id"], 10, 64)
		if err != nil {
			t.Fatalf("resource %q has no valid id: %s", resourceAddr, err)
		}

		client := newTestClient(t)
		result, err := client.GetByID("SearchDisplay", id, []string{"settings"})
		if err != nil {
			t.Fatalf("GetByID(SearchDisplay, %d): %s", id, err)
		}

		raw, ok := result["settings"]
		if !ok || raw == nil {
			t.Fatalf("SearchDisplay %d has no settings", id)
		}
		obj, ok := raw.(map[string]any)
		if !ok {
			t.Fatalf("SearchDisplay %d settings is %T, not a JSON object (map[string]any): got %#v", id, raw, raw)
		}

		columnsRaw, ok := obj["columns"]
		if !ok {
			t.Fatalf("SearchDisplay %d settings has no top-level 'columns' key: %#v", id, obj)
		}
		columns, ok := columnsRaw.([]any)
		if !ok {
			t.Fatalf("SearchDisplay %d settings.columns is %T, not an array: %#v", id, columnsRaw, columnsRaw)
		}
		if len(columns) != len(wantColumnKeys) {
			t.Fatalf("SearchDisplay %d settings.columns has %d entries, want %d: %#v", id, len(columns), len(wantColumnKeys), columns)
		}
		for i, wantKey := range wantColumnKeys {
			col, ok := columns[i].(map[string]any)
			if !ok {
				t.Fatalf("SearchDisplay %d settings.columns[%d] is %T, not an object: %#v", id, i, columns[i], columns[i])
			}
			gotKey, _ := col["key"].(string)
			if gotKey != wantKey {
				t.Fatalf("SearchDisplay %d settings.columns[%d].key = %q, want %q", id, i, gotKey, wantKey)
			}
		}
		return nil
	}
}

func TestAccSearchDisplayResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "SearchDisplay", "civicrm_search_display"),
		Steps: []resource.TestStep{
			// Create: table display with labeled columns, the concrete case from issue #11.
			{
				Config: providerConfig() + `
resource "civicrm_saved_search" "test" {
  name        = "tf_acc_search_display_ss"
  label       = "TF Acceptance SearchDisplay SavedSearch"
  api_entity  = "Contact"
  api_params  = jsonencode({
    version = 4
    select  = ["id", "display_name", "contact_type"]
    where   = [["contact_type", "=", "Individual"]]
  })
}

resource "civicrm_search_display" "test" {
  label            = "TF Acceptance Test Search Display"
  saved_search_id  = civicrm_saved_search.test.id
  type             = "table"
  acl_bypass       = false
  settings = jsonencode({
    limit   = 25
    columns = [
      { type = "field", key = "id", label = "ID", sortable = true },
      { type = "field", key = "display_name", label = "Name", sortable = true },
    ]
  })
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_search_display.test", "id"),
					resource.TestCheckResourceAttr("civicrm_search_display.test", "label", "TF Acceptance Test Search Display"),
					resource.TestCheckResourceAttr("civicrm_search_display.test", "type", "table"),
					resource.TestCheckResourceAttr("civicrm_search_display.test", "acl_bypass", "false"),
					checkEntityAttr(t, "SearchDisplay", "civicrm_search_display.test", "id", "label", "TF Acceptance Test Search Display"),
					checkEntityAttr(t, "SearchDisplay", "civicrm_search_display.test", "id", "type", "table"),
					checkSearchDisplaySettingsIsObject(t, "civicrm_search_display.test", []string{"id", "display_name"}),
				),
			},
			// Import
			{
				ResourceName:            "civicrm_search_display.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"settings"},
			},
			// Update: change a column label, verify it's persisted (not just round-tripped).
			{
				Config: providerConfig() + `
resource "civicrm_saved_search" "test" {
  name        = "tf_acc_search_display_ss"
  label       = "TF Acceptance SearchDisplay SavedSearch"
  api_entity  = "Contact"
  api_params  = jsonencode({
    version = 4
    select  = ["id", "display_name", "contact_type"]
    where   = [["contact_type", "=", "Individual"]]
  })
}

resource "civicrm_search_display" "test" {
  label            = "TF Acceptance Test Search Display Updated"
  saved_search_id  = civicrm_saved_search.test.id
  type             = "table"
  settings = jsonencode({
    limit   = 25
    columns = [
      { type = "field", key = "id", label = "ID Updated", sortable = true },
      { type = "field", key = "display_name", label = "Name", sortable = true },
    ]
  })
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_search_display.test", "label", "TF Acceptance Test Search Display Updated"),
					checkEntityAttr(t, "SearchDisplay", "civicrm_search_display.test", "id", "label", "TF Acceptance Test Search Display Updated"),
					checkSearchDisplaySettingsIsObject(t, "civicrm_search_display.test", []string{"id", "display_name"}),
				),
			},
		},
	})
}

// TestAccSearchDisplayResource_AllTypes verifies that every SearchKit display
// type accepted by the resource's type validator can actually be created in
// CiviCRM, without a plan-time inconsistency after apply. "entity" requires
// an explicit name (CiviCRM registers a virtual DB entity named after the
// display); the resource surfaces that as a clear validation error instead
// of CiviCRM's internal TypeError if name is left unset — covered by
// TestAccSearchDisplayResource_EntityTypeRequiresName below. "batch" and
// "entity" specifically exercise the settings.columns[].spec stripping (see
// stripColumnSpecs in resource_search_display.go): CiviCRM auto-populates a
// spec sub-object on saved column definitions for these two types (verified
// against a live instance — table/list/grid/tree/autocomplete do not get
// this treatment), which would otherwise show as permanent drift.
func TestAccSearchDisplayResource_AllTypes(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "SearchDisplay", "civicrm_search_display"),
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_saved_search" "test" {
  name        = "tf_acc_sd_alltypes_ss"
  label       = "TF Acceptance SearchDisplay AllTypes SavedSearch"
  api_entity  = "Contact"
  api_params  = jsonencode({
    version = 4
    select  = ["id", "display_name"]
    where   = [["contact_type", "=", "Individual"]]
  })
}

resource "civicrm_search_display" "table" {
  label            = "tf_acc_sd_table"
  saved_search_id  = civicrm_saved_search.test.id
  type             = "table"
  settings = jsonencode({
    columns = [{ type = "field", key = "id", label = "ID" }]
  })
}

resource "civicrm_search_display" "list" {
  label            = "tf_acc_sd_list"
  saved_search_id  = civicrm_saved_search.test.id
  type             = "list"
  settings = jsonencode({
    title = { text = "[display_name]" }
  })
}

resource "civicrm_search_display" "grid" {
  label            = "tf_acc_sd_grid"
  saved_search_id  = civicrm_saved_search.test.id
  type             = "grid"
  settings = jsonencode({
    title = { text = "[display_name]" }
  })
}

resource "civicrm_search_display" "tree" {
  label            = "tf_acc_sd_tree"
  saved_search_id  = civicrm_saved_search.test.id
  type             = "tree"
  settings = jsonencode({
    label = { text = "[display_name]" }
  })
}

resource "civicrm_search_display" "autocomplete" {
  label            = "tf_acc_sd_autocomplete"
  saved_search_id  = civicrm_saved_search.test.id
  type             = "autocomplete"
  settings = jsonencode({
    label = { text = "[display_name]" }
  })
}

resource "civicrm_search_display" "entity" {
  name             = "tf_acc_sd_entity"
  label            = "tf_acc_sd_entity"
  saved_search_id  = civicrm_saved_search.test.id
  type             = "entity"
  settings = jsonencode({
    columns = [{ type = "field", key = "id" }]
  })
}

resource "civicrm_search_display" "batch" {
  label            = "tf_acc_sd_batch"
  saved_search_id  = civicrm_saved_search.test.id
  type             = "batch"
  settings = jsonencode({
    columns = [{ type = "field", key = "id" }]
  })
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_search_display.table", "type", "table"),
					resource.TestCheckResourceAttr("civicrm_search_display.list", "type", "list"),
					resource.TestCheckResourceAttr("civicrm_search_display.grid", "type", "grid"),
					resource.TestCheckResourceAttr("civicrm_search_display.tree", "type", "tree"),
					resource.TestCheckResourceAttr("civicrm_search_display.autocomplete", "type", "autocomplete"),
					resource.TestCheckResourceAttr("civicrm_search_display.entity", "type", "entity"),
					resource.TestCheckResourceAttr("civicrm_search_display.batch", "type", "batch"),
					checkSearchDisplaySettingsIsObject(t, "civicrm_search_display.batch", []string{"id"}),
					checkSearchDisplaySettingsIsObject(t, "civicrm_search_display.entity", []string{"id"}),
				),
			},
		},
	})
}

func TestAccSearchDisplayResource_EntityTypeRequiresName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_saved_search" "test" {
  name        = "tf_acc_sd_entity_noname_ss"
  api_entity  = "Contact"
  api_params  = jsonencode({ version = 4, select = ["id"] })
}

resource "civicrm_search_display" "test" {
  label            = "tf_acc_sd_entity_noname"
  saved_search_id  = civicrm_saved_search.test.id
  type             = "entity"
  settings         = jsonencode({ columns = [{ type = "field", key = "id" }] })
}`,
				ExpectError: regexp.MustCompile(`name is required for type "entity"`),
			},
		},
	})
}
