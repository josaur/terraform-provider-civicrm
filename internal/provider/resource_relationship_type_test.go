package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRelationshipTypeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "RelationshipType", "civicrm_relationship_type"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + `
resource "civicrm_relationship_type" "test" {
  name_a_b    = "tf_acc_rel_type_a_b"
  label_a_b   = "TF Acc has relationship with"
  name_b_a    = "tf_acc_rel_type_b_a"
  label_b_a   = "TF Acc is related to"
  description = "tf acc relationship type description"
  is_active   = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_relationship_type.test", "id"),
					resource.TestCheckResourceAttr("civicrm_relationship_type.test", "name_a_b", "tf_acc_rel_type_a_b"),
					resource.TestCheckResourceAttr("civicrm_relationship_type.test", "label_a_b", "TF Acc has relationship with"),
					resource.TestCheckResourceAttr("civicrm_relationship_type.test", "is_active", "true"),
					checkEntityAttr(t, "RelationshipType", "civicrm_relationship_type.test", "id", "name_a_b", "tf_acc_rel_type_a_b"),
					checkEntityAttr(t, "RelationshipType", "civicrm_relationship_type.test", "id", "label_a_b", "TF Acc has relationship with"),
					checkEntityAttr(t, "RelationshipType", "civicrm_relationship_type.test", "id", "name_b_a", "tf_acc_rel_type_b_a"),
					checkEntityAttr(t, "RelationshipType", "civicrm_relationship_type.test", "id", "label_b_a", "TF Acc is related to"),
					checkEntityAttr(t, "RelationshipType", "civicrm_relationship_type.test", "id", "description", "tf acc relationship type description"),
					checkEntityAttr(t, "RelationshipType", "civicrm_relationship_type.test", "id", "is_active", "true"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_relationship_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_relationship_type" "test" {
  name_a_b    = "tf_acc_rel_type_a_b"
  label_a_b   = "TF Acc has relationship with (updated)"
  name_b_a    = "tf_acc_rel_type_b_a"
  label_b_a   = "TF Acc is related to (updated)"
  description = "tf acc relationship type description updated"
  is_active   = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_relationship_type.test", "label_a_b", "TF Acc has relationship with (updated)"),
					checkEntityAttr(t, "RelationshipType", "civicrm_relationship_type.test", "id", "label_a_b", "TF Acc has relationship with (updated)"),
					checkEntityAttr(t, "RelationshipType", "civicrm_relationship_type.test", "id", "label_b_a", "TF Acc is related to (updated)"),
					checkEntityAttr(t, "RelationshipType", "civicrm_relationship_type.test", "id", "description", "tf acc relationship type description updated"),
					checkEntityAttr(t, "RelationshipType", "civicrm_relationship_type.test", "id", "is_active", "true"),
				),
			},
			// Deactivate: is_active=false must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_relationship_type" "test" {
  name_a_b  = "tf_acc_rel_type_a_b"
  label_a_b = "TF Acc has relationship with (updated)"
  name_b_a  = "tf_acc_rel_type_b_a"
  label_b_a = "TF Acc is related to (updated)"
  is_active = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_relationship_type.test", "is_active", "false"),
					checkEntityAttr(t, "RelationshipType", "civicrm_relationship_type.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
