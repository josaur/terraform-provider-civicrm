package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const tfAccPriceSetAndFieldConfig = `
resource "civicrm_price_set" "pfv_test" {
  name    = "tf_acc_pfv_price_set"
  title   = "TF Acceptance Price Set for Field Value"
  extends = "CiviMember"
}

resource "civicrm_price_field" "pfv_test" {
  price_set_id = civicrm_price_set.pfv_test.id
  name         = "tf_acc_pfv_price_field"
  label        = "TF Acceptance Price Field for Value"
  html_type    = "Radio"
}
`

func TestAccPriceFieldValueResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "PriceFieldValue", "civicrm_price_field_value"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + tfAccPriceSetAndFieldConfig + `
resource "civicrm_price_field_value" "test" {
  price_field_id = civicrm_price_field.pfv_test.id
  name           = "tf_acc_price_field_value"
  label          = "TF Acceptance Price Field Value"
  amount         = 10.00
  is_active      = true
  is_default     = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_price_field_value.test", "id"),
					resource.TestCheckResourceAttr("civicrm_price_field_value.test", "name", "tf_acc_price_field_value"),
					resource.TestCheckResourceAttr("civicrm_price_field_value.test", "label", "TF Acceptance Price Field Value"),
					resource.TestCheckResourceAttr("civicrm_price_field_value.test", "amount", "10"),
					resource.TestCheckResourceAttr("civicrm_price_field_value.test", "is_active", "true"),
					resource.TestCheckResourceAttr("civicrm_price_field_value.test", "is_default", "false"),
					checkEntityAttr(t, "PriceFieldValue", "civicrm_price_field_value.test", "id", "name", "tf_acc_price_field_value"),
					checkEntityAttr(t, "PriceFieldValue", "civicrm_price_field_value.test", "id", "label", "TF Acceptance Price Field Value"),
					checkEntityAttr(t, "PriceFieldValue", "civicrm_price_field_value.test", "id", "is_active", "true"),
					checkEntityAttr(t, "PriceFieldValue", "civicrm_price_field_value.test", "id", "is_default", "false"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_price_field_value.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + tfAccPriceSetAndFieldConfig + `
resource "civicrm_price_field_value" "test" {
  price_field_id = civicrm_price_field.pfv_test.id
  name           = "tf_acc_price_field_value"
  label          = "TF Acceptance Price Field Value Updated"
  amount         = 20.00
  is_active      = true
  is_default     = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_price_field_value.test", "label", "TF Acceptance Price Field Value Updated"),
					resource.TestCheckResourceAttr("civicrm_price_field_value.test", "amount", "20"),
					checkEntityAttr(t, "PriceFieldValue", "civicrm_price_field_value.test", "id", "label", "TF Acceptance Price Field Value Updated"),
					checkEntityAttr(t, "PriceFieldValue", "civicrm_price_field_value.test", "id", "is_active", "true"),
					checkEntityAttr(t, "PriceFieldValue", "civicrm_price_field_value.test", "id", "is_default", "false"),
				),
			},
			// Deactivate: is_active=false must be persisted in DB.
			{
				Config: providerConfig() + tfAccPriceSetAndFieldConfig + `
resource "civicrm_price_field_value" "test" {
  price_field_id = civicrm_price_field.pfv_test.id
  name           = "tf_acc_price_field_value"
  label          = "TF Acceptance Price Field Value Updated"
  amount         = 20.00
  is_active      = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_price_field_value.test", "is_active", "false"),
					checkEntityAttr(t, "PriceFieldValue", "civicrm_price_field_value.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
