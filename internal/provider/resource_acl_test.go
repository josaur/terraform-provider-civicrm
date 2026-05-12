package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccACLResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
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
				),
			},
			{
				ResourceName:      "civicrm_acl.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
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
				Check: resource.TestCheckResourceAttr("civicrm_acl.test", "operation", "Edit"),
			},
		},
	})
}
