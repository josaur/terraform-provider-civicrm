package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccACLResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "ACL", "civicrm_acl"),
		Steps: []resource.TestStep{
			// Create without explicit priority; CiviCRM auto-assigns it.
			{
				Config: providerConfig() + `
resource "civicrm_group" "acl_test" {
  name  = "tf_acc_acl_group"
  title = "TF Acceptance ACL Group"
}

resource "civicrm_acl_role" "acl_test" {
  name  = "tf_acc_acl_role_for_acl"
  label = "TF Acceptance ACL Role for ACL"
}

resource "civicrm_acl" "test" {
  name         = "tf_acc_acl"
  entity_table = "civicrm_group"
  entity_id    = civicrm_group.acl_test.id
  operation    = "View"
  object_table = "civicrm_saved_search"
  is_active    = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_acl.test", "id"),
					resource.TestCheckResourceAttr("civicrm_acl.test", "name", "tf_acc_acl"),
					resource.TestCheckResourceAttr("civicrm_acl.test", "operation", "View"),
					resource.TestCheckResourceAttr("civicrm_acl.test", "is_active", "true"),
					// priority must be set (auto-assigned by CiviCRM)
					resource.TestCheckResourceAttrSet("civicrm_acl.test", "priority"),
					// DB verification
					checkEntityAttr(t, "ACL", "civicrm_acl.test", "id", "name", "tf_acc_acl"),
					checkEntityAttr(t, "ACL", "civicrm_acl.test", "id", "operation", "View"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_acl.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update operation without setting priority — priority must stay stable in DB.
			// This is the regression test for the "inconsistent result after apply" bug:
			// the provider was sending priority=null on update, causing CiviCRM to reassign it.
			{
				Config: providerConfig() + `
resource "civicrm_group" "acl_test" {
  name  = "tf_acc_acl_group"
  title = "TF Acceptance ACL Group"
}

resource "civicrm_acl_role" "acl_test" {
  name  = "tf_acc_acl_role_for_acl"
  label = "TF Acceptance ACL Role for ACL"
}

resource "civicrm_acl" "test" {
  name         = "tf_acc_acl_updated"
  entity_table = "civicrm_group"
  entity_id    = civicrm_group.acl_test.id
  operation    = "Edit"
  object_table = "civicrm_saved_search"
  is_active    = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_acl.test", "operation", "Edit"),
					// priority must still be set and must match what CiviCRM has in DB
					resource.TestCheckResourceAttrSet("civicrm_acl.test", "priority"),
					checkEntityAttr(t, "ACL", "civicrm_acl.test", "id", "name", "tf_acc_acl_updated"),
					checkEntityAttr(t, "ACL", "civicrm_acl.test", "id", "operation", "Edit"),
				),
			},
		},
	})
}

// TestAccACLResourcePriorityStability is a targeted regression test for the bug where
// updating an ACL without an explicit priority in config caused the provider to send
// priority=null to CiviCRM, which reassigned it to a new value — producing the
// "inconsistent result after apply" error.
func TestAccACLResourcePriorityStability(t *testing.T) {
	var capturedPriority string

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "ACL", "civicrm_acl"),
		Steps: []resource.TestStep{
			// Create without explicit priority; capture what CiviCRM assigns.
			{
				Config: providerConfig() + `
resource "civicrm_acl" "prio_stable" {
  name         = "tf_acc_acl_prio_stable"
  entity_table = "civicrm_acl_role"
  entity_id    = 1
  operation    = "View"
  object_table = "civicrm_group"
  is_active    = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_acl.prio_stable", "priority"),
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["civicrm_acl.prio_stable"]
						if ok {
							capturedPriority = rs.Primary.Attributes["priority"]
						}
						return nil
					},
				),
			},
			// Update name only (no priority in config) — priority must not change.
			{
				Config: providerConfig() + `
resource "civicrm_acl" "prio_stable" {
  name         = "tf_acc_acl_prio_stable_updated"
  entity_table = "civicrm_acl_role"
  entity_id    = 1
  operation    = "View"
  object_table = "civicrm_group"
  is_active    = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_acl.prio_stable", "name", "tf_acc_acl_prio_stable_updated"),
					resource.TestCheckResourceAttrWith("civicrm_acl.prio_stable", "priority", func(got string) error {
						if got != capturedPriority {
							return fmt.Errorf("priority changed after update: was %q, now %q", capturedPriority, got)
						}
						return nil
					}),
					// DB must also have the original priority
					checkEntityAttr(t, "ACL", "civicrm_acl.prio_stable", "id", "name", "tf_acc_acl_prio_stable_updated"),
				),
			},
		},
	})
}
