package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPriceFieldResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_price_set" "field_test" {
  name    = "tf_acc_pf_price_set"
  title   = "TF Acceptance Price Set for Field"
  extends = "CiviMember"
}

resource "civicrm_price_field" "test" {
  price_set_id = civicrm_price_set.field_test.id
  name         = "tf_acc_price_field"
  label        = "TF Acceptance Price Field"
  html_type    = "Radio"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_price_field.test", "id"),
					resource.TestCheckResourceAttr("civicrm_price_field.test", "name", "tf_acc_price_field"),
					resource.TestCheckResourceAttr("civicrm_price_field.test", "label", "TF Acceptance Price Field"),
					resource.TestCheckResourceAttr("civicrm_price_field.test", "html_type", "Radio"),
					resource.TestCheckResourceAttr("civicrm_price_field.test", "is_active", "true"),
				),
			},
			{
				ResourceName:      "civicrm_price_field.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_price_set" "field_test" {
  name    = "tf_acc_pf_price_set"
  title   = "TF Acceptance Price Set for Field"
  extends = "CiviMember"
}

resource "civicrm_price_field" "test" {
  price_set_id = civicrm_price_set.field_test.id
  name         = "tf_acc_price_field"
  label        = "TF Acceptance Price Field Updated"
  html_type    = "Radio"
}`,
				Check: resource.TestCheckResourceAttr("civicrm_price_field.test", "label", "TF Acceptance Price Field Updated"),
			},
		},
	})
}
