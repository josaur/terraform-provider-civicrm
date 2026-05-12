package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCustomGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_custom_group" "test" {
  name       = "tf_acc_cg"
  title      = "TF Acceptance Custom Group"
  extends    = "Contact"
  table_name = "civicrm_value_tf_acc_cg"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_custom_group.test", "id"),
					resource.TestCheckResourceAttr("civicrm_custom_group.test", "name", "tf_acc_cg"),
					resource.TestCheckResourceAttr("civicrm_custom_group.test", "title", "TF Acceptance Custom Group"),
					resource.TestCheckResourceAttr("civicrm_custom_group.test", "extends", "Contact"),
					resource.TestCheckResourceAttr("civicrm_custom_group.test", "is_active", "true"),
				),
			},
			{
				ResourceName:      "civicrm_custom_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_custom_group" "test" {
  name       = "tf_acc_cg"
  title      = "TF Acceptance Custom Group Updated"
  extends    = "Contact"
  table_name = "civicrm_value_tf_acc_cg"
}`,
				Check: resource.TestCheckResourceAttr("civicrm_custom_group.test", "title", "TF Acceptance Custom Group Updated"),
			},
		},
	})
}
