package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccACLRoleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "OptionValue", "civicrm_acl_role"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + `
resource "civicrm_acl_role" "test" {
  name        = "tf_acc_acl_role"
  label       = "TF Acceptance ACL Role"
  description = "tf acc description"
  is_active   = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_acl_role.test", "id"),
					resource.TestCheckResourceAttr("civicrm_acl_role.test", "name", "tf_acc_acl_role"),
					resource.TestCheckResourceAttr("civicrm_acl_role.test", "label", "TF Acceptance ACL Role"),
					resource.TestCheckResourceAttr("civicrm_acl_role.test", "is_active", "true"),
					checkEntityAttr(t, "OptionValue", "civicrm_acl_role.test", "id", "name", "tf_acc_acl_role"),
					checkEntityAttr(t, "OptionValue", "civicrm_acl_role.test", "id", "label", "TF Acceptance ACL Role"),
					checkEntityAttr(t, "OptionValue", "civicrm_acl_role.test", "id", "description", "tf acc description"),
					checkEntityAttr(t, "OptionValue", "civicrm_acl_role.test", "id", "is_active", "true"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_acl_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_acl_role" "test" {
  name        = "tf_acc_acl_role"
  label       = "TF Acceptance ACL Role Updated"
  description = "tf acc description updated"
  is_active   = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_acl_role.test", "label", "TF Acceptance ACL Role Updated"),
					checkEntityAttr(t, "OptionValue", "civicrm_acl_role.test", "id", "label", "TF Acceptance ACL Role Updated"),
					checkEntityAttr(t, "OptionValue", "civicrm_acl_role.test", "id", "description", "tf acc description updated"),
					checkEntityAttr(t, "OptionValue", "civicrm_acl_role.test", "id", "is_active", "true"),
				),
			},
			// Deactivate: is_active=false must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_acl_role" "test" {
  name      = "tf_acc_acl_role"
  label     = "TF Acceptance ACL Role Updated"
  is_active = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_acl_role.test", "is_active", "false"),
					checkEntityAttr(t, "OptionValue", "civicrm_acl_role.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
