package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRelationshipTypeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_relationship_type" "test" {
  name_a_b  = "tf_acc_rel_type_a_b"
  label_a_b = "TF Acc has relationship with"
  name_b_a  = "tf_acc_rel_type_b_a"
  label_b_a = "TF Acc is related to"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_relationship_type.test", "id"),
					resource.TestCheckResourceAttr("civicrm_relationship_type.test", "name_a_b", "tf_acc_rel_type_a_b"),
					resource.TestCheckResourceAttr("civicrm_relationship_type.test", "label_a_b", "TF Acc has relationship with"),
					resource.TestCheckResourceAttr("civicrm_relationship_type.test", "is_active", "true"),
				),
			},
			{
				ResourceName:      "civicrm_relationship_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_relationship_type" "test" {
  name_a_b  = "tf_acc_rel_type_a_b"
  label_a_b = "TF Acc has relationship with (updated)"
  name_b_a  = "tf_acc_rel_type_b_a"
  label_b_a = "TF Acc is related to (updated)"
}`,
				Check: resource.TestCheckResourceAttr("civicrm_relationship_type.test", "label_a_b", "TF Acc has relationship with (updated)"),
			},
		},
	})
}
