package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTagsetResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_tagset" "test" {
  name     = "tf_acc_tagset"
  label    = "TF Acceptance Test Tagset"
  used_for = ["civicrm_contact"]
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_tagset.test", "id"),
					resource.TestCheckResourceAttr("civicrm_tagset.test", "name", "tf_acc_tagset"),
					resource.TestCheckResourceAttr("civicrm_tagset.test", "label", "TF Acceptance Test Tagset"),
				),
			},
			{
				ResourceName:      "civicrm_tagset.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_tagset" "test" {
  name     = "tf_acc_tagset"
  label    = "TF Acceptance Test Tagset Updated"
  used_for = ["civicrm_contact"]
}`,
				Check: resource.TestCheckResourceAttr("civicrm_tagset.test", "label", "TF Acceptance Test Tagset Updated"),
			},
		},
	})
}
