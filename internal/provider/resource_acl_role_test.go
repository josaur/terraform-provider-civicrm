package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccACLRoleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_acl_role" "test" {
  name  = "tf_acc_acl_role"
  label = "TF Acceptance ACL Role"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_acl_role.test", "id"),
					resource.TestCheckResourceAttr("civicrm_acl_role.test", "name", "tf_acc_acl_role"),
					resource.TestCheckResourceAttr("civicrm_acl_role.test", "label", "TF Acceptance ACL Role"),
					resource.TestCheckResourceAttr("civicrm_acl_role.test", "is_active", "true"),
				),
			},
			{
				ResourceName:      "civicrm_acl_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_acl_role" "test" {
  name  = "tf_acc_acl_role"
  label = "TF Acceptance ACL Role Updated"
}`,
				Check: resource.TestCheckResourceAttr("civicrm_acl_role.test", "label", "TF Acceptance ACL Role Updated"),
			},
		},
	})
}
