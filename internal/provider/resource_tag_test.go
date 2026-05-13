package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTagResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "Tag", "civicrm_tag"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + `
resource "civicrm_tag" "test" {
  name        = "tf_acc_tag"
  label       = "TF Acceptance Test Tag"
  description = "tf acc tag description"
  used_for    = ["civicrm_contact"]
  is_selectable = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_tag.test", "id"),
					resource.TestCheckResourceAttr("civicrm_tag.test", "name", "tf_acc_tag"),
					resource.TestCheckResourceAttr("civicrm_tag.test", "label", "TF Acceptance Test Tag"),
					resource.TestCheckResourceAttr("civicrm_tag.test", "is_selectable", "true"),
					checkEntityAttr(t, "Tag", "civicrm_tag.test", "id", "name", "tf_acc_tag"),
					checkEntityAttr(t, "Tag", "civicrm_tag.test", "id", "label", "TF Acceptance Test Tag"),
					checkEntityAttr(t, "Tag", "civicrm_tag.test", "id", "description", "tf acc tag description"),
					checkEntityAttr(t, "Tag", "civicrm_tag.test", "id", "is_selectable", "true"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_tag.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_tag" "test" {
  name          = "tf_acc_tag"
  label         = "TF Acceptance Test Tag Updated"
  description   = "tf acc tag description updated"
  used_for      = ["civicrm_contact"]
  is_selectable = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_tag.test", "label", "TF Acceptance Test Tag Updated"),
					checkEntityAttr(t, "Tag", "civicrm_tag.test", "id", "label", "TF Acceptance Test Tag Updated"),
					checkEntityAttr(t, "Tag", "civicrm_tag.test", "id", "description", "tf acc tag description updated"),
					checkEntityAttr(t, "Tag", "civicrm_tag.test", "id", "is_selectable", "true"),
				),
			},
			// Set is_selectable=false: must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_tag" "test" {
  name          = "tf_acc_tag"
  label         = "TF Acceptance Test Tag Updated"
  used_for      = ["civicrm_contact"]
  is_selectable = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_tag.test", "is_selectable", "false"),
					checkEntityAttr(t, "Tag", "civicrm_tag.test", "id", "is_selectable", "false"),
				),
			},
		},
	})
}
