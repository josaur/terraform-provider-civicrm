package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPriceSetResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "PriceSet", "civicrm_price_set"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + `
resource "civicrm_price_set" "test" {
  name      = "tf_acc_price_set"
  title     = "TF Acceptance Price Set"
  extends   = "CiviMember"
  is_active = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_price_set.test", "id"),
					resource.TestCheckResourceAttr("civicrm_price_set.test", "name", "tf_acc_price_set"),
					resource.TestCheckResourceAttr("civicrm_price_set.test", "title", "TF Acceptance Price Set"),
					resource.TestCheckResourceAttr("civicrm_price_set.test", "is_active", "true"),
					checkEntityAttr(t, "PriceSet", "civicrm_price_set.test", "id", "name", "tf_acc_price_set"),
					checkEntityAttr(t, "PriceSet", "civicrm_price_set.test", "id", "title", "TF Acceptance Price Set"),
					checkEntityAttr(t, "PriceSet", "civicrm_price_set.test", "id", "is_active", "true"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_price_set.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_price_set" "test" {
  name      = "tf_acc_price_set"
  title     = "TF Acceptance Price Set Updated"
  extends   = "CiviMember"
  is_active = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_price_set.test", "title", "TF Acceptance Price Set Updated"),
					checkEntityAttr(t, "PriceSet", "civicrm_price_set.test", "id", "title", "TF Acceptance Price Set Updated"),
					checkEntityAttr(t, "PriceSet", "civicrm_price_set.test", "id", "is_active", "true"),
				),
			},
			// Deactivate: is_active=false must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_price_set" "test" {
  name      = "tf_acc_price_set"
  title     = "TF Acceptance Price Set Updated"
  extends   = "CiviMember"
  is_active = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_price_set.test", "is_active", "false"),
					checkEntityAttr(t, "PriceSet", "civicrm_price_set.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
