package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPriceSetResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_price_set" "test" {
  name    = "tf_acc_price_set"
  title   = "TF Acceptance Price Set"
  extends = "CiviMember"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_price_set.test", "id"),
					resource.TestCheckResourceAttr("civicrm_price_set.test", "name", "tf_acc_price_set"),
					resource.TestCheckResourceAttr("civicrm_price_set.test", "title", "TF Acceptance Price Set"),
					resource.TestCheckResourceAttr("civicrm_price_set.test", "is_active", "true"),
				),
			},
			{
				ResourceName:      "civicrm_price_set.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_price_set" "test" {
  name    = "tf_acc_price_set"
  title   = "TF Acceptance Price Set Updated"
  extends = "CiviMember"
}`,
				Check: resource.TestCheckResourceAttr("civicrm_price_set.test", "title", "TF Acceptance Price Set Updated"),
			},
		},
	})
}
