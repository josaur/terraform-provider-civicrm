package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMembershipTypeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_membership_type" "test" {
  name                = "tf_acc_membership_type"
  member_of_contact_id = 1
  financial_type_id   = 2
  duration_unit       = "year"
  duration_interval   = 1
  period_type         = "rolling"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_membership_type.test", "id"),
					resource.TestCheckResourceAttr("civicrm_membership_type.test", "name", "tf_acc_membership_type"),
					resource.TestCheckResourceAttr("civicrm_membership_type.test", "duration_unit", "year"),
					resource.TestCheckResourceAttr("civicrm_membership_type.test", "period_type", "rolling"),
					resource.TestCheckResourceAttr("civicrm_membership_type.test", "is_active", "true"),
				),
			},
			{
				ResourceName:      "civicrm_membership_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_membership_type" "test" {
  name                = "tf_acc_membership_type_updated"
  member_of_contact_id = 1
  financial_type_id   = 2
  duration_unit       = "year"
  duration_interval   = 2
  period_type         = "rolling"
}`,
				Check: resource.TestCheckResourceAttr("civicrm_membership_type.test", "duration_interval", "2"),
			},
		},
	})
}
