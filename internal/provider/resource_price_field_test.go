package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const tfAccPriceSetConfig = `
resource "civicrm_price_set" "field_test" {
  name    = "tf_acc_pf_price_set"
  title   = "TF Acceptance Price Set for Field"
  extends = "CiviMember"
}
`

func TestAccPriceFieldResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "PriceField", "civicrm_price_field"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + tfAccPriceSetConfig + `
resource "civicrm_price_field" "test" {
  price_set_id = civicrm_price_set.field_test.id
  name         = "tf_acc_price_field"
  label        = "TF Acceptance Price Field"
  html_type    = "Radio"
  is_active    = true
  is_required  = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_price_field.test", "id"),
					resource.TestCheckResourceAttr("civicrm_price_field.test", "name", "tf_acc_price_field"),
					resource.TestCheckResourceAttr("civicrm_price_field.test", "label", "TF Acceptance Price Field"),
					resource.TestCheckResourceAttr("civicrm_price_field.test", "html_type", "Radio"),
					resource.TestCheckResourceAttr("civicrm_price_field.test", "is_active", "true"),
					resource.TestCheckResourceAttr("civicrm_price_field.test", "is_required", "true"),
					checkEntityAttr(t, "PriceField", "civicrm_price_field.test", "id", "name", "tf_acc_price_field"),
					checkEntityAttr(t, "PriceField", "civicrm_price_field.test", "id", "label", "TF Acceptance Price Field"),
					checkEntityAttr(t, "PriceField", "civicrm_price_field.test", "id", "html_type", "Radio"),
					checkEntityAttr(t, "PriceField", "civicrm_price_field.test", "id", "is_active", "true"),
					checkEntityAttr(t, "PriceField", "civicrm_price_field.test", "id", "is_required", "true"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_price_field.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + tfAccPriceSetConfig + `
resource "civicrm_price_field" "test" {
  price_set_id = civicrm_price_set.field_test.id
  name         = "tf_acc_price_field"
  label        = "TF Acceptance Price Field Updated"
  html_type    = "Radio"
  is_active    = true
  is_required  = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_price_field.test", "label", "TF Acceptance Price Field Updated"),
					resource.TestCheckResourceAttr("civicrm_price_field.test", "is_required", "false"),
					checkEntityAttr(t, "PriceField", "civicrm_price_field.test", "id", "label", "TF Acceptance Price Field Updated"),
					checkEntityAttr(t, "PriceField", "civicrm_price_field.test", "id", "is_active", "true"),
					checkEntityAttr(t, "PriceField", "civicrm_price_field.test", "id", "is_required", "false"),
				),
			},
			// Deactivate: is_active=false must be persisted in DB.
			{
				Config: providerConfig() + tfAccPriceSetConfig + `
resource "civicrm_price_field" "test" {
  price_set_id = civicrm_price_set.field_test.id
  name         = "tf_acc_price_field"
  label        = "TF Acceptance Price Field Updated"
  html_type    = "Radio"
  is_active    = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_price_field.test", "is_active", "false"),
					checkEntityAttr(t, "PriceField", "civicrm_price_field.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
