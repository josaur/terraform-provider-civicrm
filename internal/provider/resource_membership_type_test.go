package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMembershipTypeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "MembershipType", "civicrm_membership_type"),
		Steps: []resource.TestStep{
			// Create: verify state and DB.
			{
				Config: providerConfig() + `
resource "civicrm_membership_type" "test" {
  name                 = "tf_acc_membership_type"
  member_of_contact_id = 1
  financial_type_id    = 2
  duration_unit        = "year"
  duration_interval    = 1
  period_type          = "rolling"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_membership_type.test", "id"),
					resource.TestCheckResourceAttr("civicrm_membership_type.test", "name", "tf_acc_membership_type"),
					resource.TestCheckResourceAttr("civicrm_membership_type.test", "duration_unit", "year"),
					resource.TestCheckResourceAttr("civicrm_membership_type.test", "period_type", "rolling"),
					resource.TestCheckResourceAttr("civicrm_membership_type.test", "is_active", "true"),
					// DB verification
					checkEntityAttr(t, "MembershipType", "civicrm_membership_type.test", "id", "name", "tf_acc_membership_type"),
					checkEntityAttr(t, "MembershipType", "civicrm_membership_type.test", "id", "duration_unit", "year"),
					checkEntityAttr(t, "MembershipType", "civicrm_membership_type.test", "id", "period_type", "rolling"),
					checkEntityAttr(t, "MembershipType", "civicrm_membership_type.test", "id", "is_active", "true"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_membership_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update name and duration_interval; verify both are reflected in the DB.
			{
				Config: providerConfig() + `
resource "civicrm_membership_type" "test" {
  name                 = "tf_acc_membership_type_updated"
  member_of_contact_id = 1
  financial_type_id    = 2
  duration_unit        = "year"
  duration_interval    = 2
  period_type          = "rolling"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_membership_type.test", "duration_interval", "2"),
					resource.TestCheckResourceAttr("civicrm_membership_type.test", "name", "tf_acc_membership_type_updated"),
					// DB verification after update
					checkEntityAttr(t, "MembershipType", "civicrm_membership_type.test", "id", "name", "tf_acc_membership_type_updated"),
					checkEntityAttr(t, "MembershipType", "civicrm_membership_type.test", "id", "duration_interval", "2"),
				),
			},
			// Deactivate: is_active=false must be persisted.
			{
				Config: providerConfig() + `
resource "civicrm_membership_type" "test" {
  name                 = "tf_acc_membership_type_updated"
  member_of_contact_id = 1
  financial_type_id    = 2
  duration_unit        = "year"
  duration_interval    = 2
  period_type          = "rolling"
  is_active            = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_membership_type.test", "is_active", "false"),
					checkEntityAttr(t, "MembershipType", "civicrm_membership_type.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
