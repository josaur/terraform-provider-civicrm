package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPriceFieldValueResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
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

resource "civicrm_price_field_value" "test" {
  price_field_id = civicrm_price_field.pfv_test.id
  name           = "tf_acc_price_field_value"
  label          = "TF Acceptance Price Field Value"
  amount         = 10.00
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_price_field_value.test", "id"),
					resource.TestCheckResourceAttr("civicrm_price_field_value.test", "name", "tf_acc_price_field_value"),
					resource.TestCheckResourceAttr("civicrm_price_field_value.test", "label", "TF Acceptance Price Field Value"),
					resource.TestCheckResourceAttr("civicrm_price_field_value.test", "amount", "10"),
				),
			},
			{
				ResourceName:      "civicrm_price_field_value.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
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

resource "civicrm_price_field_value" "test" {
  price_field_id = civicrm_price_field.pfv_test.id
  name           = "tf_acc_price_field_value"
  label          = "TF Acceptance Price Field Value Updated"
  amount         = 20.00
}`,
				Check: resource.TestCheckResourceAttr("civicrm_price_field_value.test", "amount", "20"),
			},
		},
	})
}
