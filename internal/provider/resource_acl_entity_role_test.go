package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccACLEntityRoleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_group" "entity_role_test" {
  name  = "tf_acc_entity_role_group"
  title = "TF Acceptance Entity Role Group"
}

resource "civicrm_acl_role" "entity_role_test" {
  name  = "tf_acc_entity_role_role"
  label = "TF Acceptance Entity Role ACL Role"
  value = 99
}

resource "civicrm_acl_entity_role" "test" {
  acl_role_id  = civicrm_acl_role.entity_role_test.value
  entity_table = "civicrm_group"
  entity_id    = civicrm_group.entity_role_test.id
  is_active    = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_acl_entity_role.test", "id"),
					resource.TestCheckResourceAttr("civicrm_acl_entity_role.test", "entity_table", "civicrm_group"),
					resource.TestCheckResourceAttr("civicrm_acl_entity_role.test", "is_active", "true"),
				),
			},
			{
				ResourceName:      "civicrm_acl_entity_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
