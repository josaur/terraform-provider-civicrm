package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

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
				),
			},
			// Import
			{
				ResourceName:      "civicrm_saved_search.test",
				ImportState:       true,
				ImportStateVerify: true,
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
				),
			},
		},
	})
}
