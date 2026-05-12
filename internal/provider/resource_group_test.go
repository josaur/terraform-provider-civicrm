package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
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
				),
			},
			{
				ResourceName:      "civicrm_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_group" "test" {
  name  = "tf_acc_group"
  title = "TF Acceptance Test Group Updated"
}`,
				Check: resource.TestCheckResourceAttr("civicrm_group.test", "title", "TF Acceptance Test Group Updated"),
			},
		},
	})
}
