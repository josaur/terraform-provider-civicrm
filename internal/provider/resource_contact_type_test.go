package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccContactTypeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// parent_id 1 = Individual (built-in CiviCRM contact type)
				Config: providerConfig() + `
resource "civicrm_contact_type" "test" {
  name      = "tf_acc_contact_type"
  label     = "TF Acceptance Contact Type"
  parent_id = 1
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_contact_type.test", "id"),
					resource.TestCheckResourceAttr("civicrm_contact_type.test", "name", "tf_acc_contact_type"),
					resource.TestCheckResourceAttr("civicrm_contact_type.test", "label", "TF Acceptance Contact Type"),
					resource.TestCheckResourceAttr("civicrm_contact_type.test", "is_active", "true"),
				),
			},
			{
				ResourceName:      "civicrm_contact_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_contact_type" "test" {
  name      = "tf_acc_contact_type"
  label     = "TF Acceptance Contact Type Updated"
  parent_id = 1
}`,
				Check: resource.TestCheckResourceAttr("civicrm_contact_type.test", "label", "TF Acceptance Contact Type Updated"),
			},
		},
	})
}
