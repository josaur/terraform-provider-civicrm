package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "Group", "civicrm_group"),
		Steps: []resource.TestStep{
			// Create: verify Terraform state and DB agree.
			{
				Config: providerConfig() + `
resource "civicrm_group" "test" {
  name  = "tf_acc_group"
  title = "TF Acceptance Test Group"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_group.test", "id"),
					resource.TestCheckResourceAttr("civicrm_group.test", "name", "tf_acc_group"),
					resource.TestCheckResourceAttr("civicrm_group.test", "title", "TF Acceptance Test Group"),
					resource.TestCheckResourceAttr("civicrm_group.test", "is_active", "true"),
					// DB verification
					checkEntityAttr(t, "Group", "civicrm_group.test", "id", "title", "TF Acceptance Test Group"),
					checkEntityAttr(t, "Group", "civicrm_group.test", "id", "name", "tf_acc_group"),
					checkEntityAttr(t, "Group", "civicrm_group.test", "id", "is_active", "true"),
				),
			},
			// Import: Terraform state must survive a round-trip through the API.
			{
				ResourceName:      "civicrm_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: title change is reflected in the DB.
			{
				Config: providerConfig() + `
resource "civicrm_group" "test" {
  name  = "tf_acc_group"
  title = "TF Acceptance Test Group Updated"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_group.test", "title", "TF Acceptance Test Group Updated"),
					// DB verification after update
					checkEntityAttr(t, "Group", "civicrm_group.test", "id", "title", "TF Acceptance Test Group Updated"),
				),
			},
			// Deactivate: is_active=false is persisted in the DB.
			{
				Config: providerConfig() + `
resource "civicrm_group" "test" {
  name      = "tf_acc_group"
  title     = "TF Acceptance Test Group Updated"
  is_active = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_group.test", "is_active", "false"),
					checkEntityAttr(t, "Group", "civicrm_group.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
