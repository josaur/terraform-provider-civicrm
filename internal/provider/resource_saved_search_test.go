package provider_test

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// checkAPIParamsIsObject asserts that SavedSearch.api_params, as returned by
// SavedSearch.get, is a JSON object (map[string]any) with select/where
// directly accessible as top-level keys — the shape CiviCRM/SearchKit
// actually reads. A round-trip through this provider's own Read path
// (checkEntityAttr) cannot distinguish this from the broken shape
// (`["{...}"]`, a one-element array containing an encoded JSON string): the
// bug was that api_params was double-JSON-encoded, so `terraform plan`
// showed no drift while SearchKit could not read the search at all. See
// https://github.com/josaur/terraform-provider-civicrm/issues/10.
func checkAPIParamsIsObject(t *testing.T, resourceAddr string, wantSelect []string) resource.TestCheckFunc {
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
		result, err := client.GetByID("SavedSearch", id, []string{"api_params"})
		if err != nil {
			t.Fatalf("GetByID(SavedSearch, %d): %s", id, err)
		}

		raw, ok := result["api_params"]
		if !ok || raw == nil {
			t.Fatalf("SavedSearch %d has no api_params", id)
		}
		obj, ok := raw.(map[string]any)
		if !ok {
			t.Fatalf("SavedSearch %d api_params is %T, not a JSON object (map[string]any) — "+
				"this is exactly the double-encoding bug from issue #10: got %#v", id, raw, raw)
		}

		selectRaw, ok := obj["select"]
		if !ok {
			t.Fatalf("SavedSearch %d api_params has no top-level 'select' key: %#v", id, obj)
		}
		selectSlice, ok := selectRaw.([]any)
		if !ok {
			t.Fatalf("SavedSearch %d api_params.select is %T, not an array: %#v", id, selectRaw, selectRaw)
		}
		if len(selectSlice) != len(wantSelect) {
			t.Fatalf("SavedSearch %d api_params.select = %v, want %v", id, selectSlice, wantSelect)
		}
		for i, want := range wantSelect {
			got, ok := selectSlice[i].(string)
			if !ok || got != want {
				t.Fatalf("SavedSearch %d api_params.select[%d] = %v, want %q", id, i, selectSlice[i], want)
			}
		}
		return nil
	}
}

func TestAccSavedSearchResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "SavedSearch", "civicrm_saved_search"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + `
resource "civicrm_saved_search" "test" {
  name        = "tf_acc_saved_search"
  label       = "TF Acceptance Test Saved Search"
  description = "tf acc saved search description"
  api_entity  = "Contact"
  api_params  = "{\"version\":4,\"select\":[\"id\"]}"
  is_template = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_saved_search.test", "id"),
					resource.TestCheckResourceAttr("civicrm_saved_search.test", "name", "tf_acc_saved_search"),
					resource.TestCheckResourceAttr("civicrm_saved_search.test", "label", "TF Acceptance Test Saved Search"),
					resource.TestCheckResourceAttr("civicrm_saved_search.test", "api_entity", "Contact"),
					resource.TestCheckResourceAttr("civicrm_saved_search.test", "is_template", "false"),
					checkEntityAttr(t, "SavedSearch", "civicrm_saved_search.test", "id", "name", "tf_acc_saved_search"),
					checkEntityAttr(t, "SavedSearch", "civicrm_saved_search.test", "id", "label", "TF Acceptance Test Saved Search"),
					checkEntityAttr(t, "SavedSearch", "civicrm_saved_search.test", "id", "description", "tf acc saved search description"),
					checkEntityAttr(t, "SavedSearch", "civicrm_saved_search.test", "id", "api_entity", "Contact"),
					checkAPIParamsIsObject(t, "civicrm_saved_search.test", []string{"id"}),
				),
			},
			// Import. api_params/form_values are jsontypes.Normalized: `plan`/`apply`
			// compare them semantically (key order, whitespace don't matter), but
			// ImportStateVerify does a plain string comparison against the
			// pre-`apply` config string, and CiviCRM API4 always returns object keys
			// alphabetically sorted regardless of the order written in config — so a
			// string compare here would spuriously fail for any config whose key
			// order isn't already alphabetical. Ignored here rather than relied on
			// staying alphabetical.
			{
				ResourceName:            "civicrm_saved_search.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"api_params"},
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_saved_search" "test" {
  name        = "tf_acc_saved_search"
  label       = "TF Acceptance Test Saved Search Updated"
  description = "tf acc saved search description updated"
  api_entity  = "Contact"
  api_params  = "{\"version\":4,\"select\":[\"id\",\"display_name\"]}"
  is_template = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_saved_search.test", "label", "TF Acceptance Test Saved Search Updated"),
					resource.TestCheckResourceAttr("civicrm_saved_search.test", "is_template", "true"),
					checkEntityAttr(t, "SavedSearch", "civicrm_saved_search.test", "id", "label", "TF Acceptance Test Saved Search Updated"),
					checkEntityAttr(t, "SavedSearch", "civicrm_saved_search.test", "id", "description", "tf acc saved search description updated"),
					checkEntityAttr(t, "SavedSearch", "civicrm_saved_search.test", "id", "is_template", "true"),
					checkAPIParamsIsObject(t, "civicrm_saved_search.test", []string{"id", "display_name"}),
				),
			},
		},
	})
}
