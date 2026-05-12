package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTagResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_tag" "test" {
  name     = "tf_acc_tag"
  label    = "TF Acceptance Test Tag"
  used_for = ["civicrm_contact"]
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_tag.test", "id"),
					resource.TestCheckResourceAttr("civicrm_tag.test", "name", "tf_acc_tag"),
					resource.TestCheckResourceAttr("civicrm_tag.test", "label", "TF Acceptance Test Tag"),
				),
			},
			{
				ResourceName:      "civicrm_tag.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_tag" "test" {
  name     = "tf_acc_tag"
  label    = "TF Acceptance Test Tag Updated"
  used_for = ["civicrm_contact"]
}`,
				Check: resource.TestCheckResourceAttr("civicrm_tag.test", "label", "TF Acceptance Test Tag Updated"),
			},
		},
	})
}
