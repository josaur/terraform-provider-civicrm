package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCaseTypeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "CaseType", "civicrm_case_type"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + `
resource "civicrm_case_type" "test" {
  name        = "tf_acc_case_type"
  title       = "TF Acceptance Case Type"
  description = "tf acc case type description"
  is_active   = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_case_type.test", "id"),
					resource.TestCheckResourceAttr("civicrm_case_type.test", "name", "tf_acc_case_type"),
					resource.TestCheckResourceAttr("civicrm_case_type.test", "title", "TF Acceptance Case Type"),
					resource.TestCheckResourceAttr("civicrm_case_type.test", "is_active", "true"),
					checkEntityAttr(t, "CaseType", "civicrm_case_type.test", "id", "name", "tf_acc_case_type"),
					checkEntityAttr(t, "CaseType", "civicrm_case_type.test", "id", "title", "TF Acceptance Case Type"),
					checkEntityAttr(t, "CaseType", "civicrm_case_type.test", "id", "description", "tf acc case type description"),
					checkEntityAttr(t, "CaseType", "civicrm_case_type.test", "id", "is_active", "true"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_case_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_case_type" "test" {
  name        = "tf_acc_case_type"
  title       = "TF Acceptance Case Type Updated"
  description = "tf acc case type description updated"
  is_active   = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_case_type.test", "title", "TF Acceptance Case Type Updated"),
					checkEntityAttr(t, "CaseType", "civicrm_case_type.test", "id", "title", "TF Acceptance Case Type Updated"),
					checkEntityAttr(t, "CaseType", "civicrm_case_type.test", "id", "description", "tf acc case type description updated"),
					checkEntityAttr(t, "CaseType", "civicrm_case_type.test", "id", "is_active", "true"),
				),
			},
			// Deactivate: is_active=false must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_case_type" "test" {
  name      = "tf_acc_case_type"
  title     = "TF Acceptance Case Type Updated"
  is_active = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_case_type.test", "is_active", "false"),
					checkEntityAttr(t, "CaseType", "civicrm_case_type.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
