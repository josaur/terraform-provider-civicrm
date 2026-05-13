package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTagsetResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "Tag", "civicrm_tagset"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + `
resource "civicrm_tagset" "test" {
  name        = "tf_acc_tagset"
  label       = "TF Acceptance Test Tagset"
  description = "tf acc tagset description"
  used_for    = ["civicrm_contact"]
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_tagset.test", "id"),
					resource.TestCheckResourceAttr("civicrm_tagset.test", "name", "tf_acc_tagset"),
					resource.TestCheckResourceAttr("civicrm_tagset.test", "label", "TF Acceptance Test Tagset"),
					checkEntityAttr(t, "Tag", "civicrm_tagset.test", "id", "name", "tf_acc_tagset"),
					checkEntityAttr(t, "Tag", "civicrm_tagset.test", "id", "label", "TF Acceptance Test Tagset"),
					checkEntityAttr(t, "Tag", "civicrm_tagset.test", "id", "description", "tf acc tagset description"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_tagset.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_tagset" "test" {
  name        = "tf_acc_tagset"
  label       = "TF Acceptance Test Tagset Updated"
  description = "tf acc tagset description updated"
  used_for    = ["civicrm_contact"]
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_tagset.test", "label", "TF Acceptance Test Tagset Updated"),
					checkEntityAttr(t, "Tag", "civicrm_tagset.test", "id", "label", "TF Acceptance Test Tagset Updated"),
					checkEntityAttr(t, "Tag", "civicrm_tagset.test", "id", "description", "tf acc tagset description updated"),
				),
			},
		},
	})
}
