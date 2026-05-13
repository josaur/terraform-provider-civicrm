package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccContactTypeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "ContactType", "civicrm_contact_type"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			// parent_id 1 = Individual (built-in CiviCRM contact type)
			{
				Config: providerConfig() + `
resource "civicrm_contact_type" "test" {
  name        = "tf_acc_contact_type"
  label       = "TF Acceptance Contact Type"
  description = "tf acc contact type description"
  parent_id   = 1
  is_active   = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_contact_type.test", "id"),
					resource.TestCheckResourceAttr("civicrm_contact_type.test", "name", "tf_acc_contact_type"),
					resource.TestCheckResourceAttr("civicrm_contact_type.test", "label", "TF Acceptance Contact Type"),
					resource.TestCheckResourceAttr("civicrm_contact_type.test", "is_active", "true"),
					checkEntityAttr(t, "ContactType", "civicrm_contact_type.test", "id", "name", "tf_acc_contact_type"),
					checkEntityAttr(t, "ContactType", "civicrm_contact_type.test", "id", "label", "TF Acceptance Contact Type"),
					checkEntityAttr(t, "ContactType", "civicrm_contact_type.test", "id", "description", "tf acc contact type description"),
					checkEntityAttr(t, "ContactType", "civicrm_contact_type.test", "id", "is_active", "true"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_contact_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_contact_type" "test" {
  name        = "tf_acc_contact_type"
  label       = "TF Acceptance Contact Type Updated"
  description = "tf acc contact type description updated"
  parent_id   = 1
  is_active   = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_contact_type.test", "label", "TF Acceptance Contact Type Updated"),
					checkEntityAttr(t, "ContactType", "civicrm_contact_type.test", "id", "label", "TF Acceptance Contact Type Updated"),
					checkEntityAttr(t, "ContactType", "civicrm_contact_type.test", "id", "description", "tf acc contact type description updated"),
					checkEntityAttr(t, "ContactType", "civicrm_contact_type.test", "id", "is_active", "true"),
				),
			},
			// Deactivate: is_active=false must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_contact_type" "test" {
  name      = "tf_acc_contact_type"
  label     = "TF Acceptance Contact Type Updated"
  parent_id = 1
  is_active = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_contact_type.test", "is_active", "false"),
					checkEntityAttr(t, "ContactType", "civicrm_contact_type.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
