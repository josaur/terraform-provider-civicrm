package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCaseStatusResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "OptionValue", "civicrm_case_status"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + `
resource "civicrm_case_status" "test" {
  name      = "tf_acc_case_status"
  label     = "TF Acceptance Case Status"
  is_active = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_case_status.test", "id"),
					resource.TestCheckResourceAttr("civicrm_case_status.test", "name", "tf_acc_case_status"),
					resource.TestCheckResourceAttr("civicrm_case_status.test", "label", "TF Acceptance Case Status"),
					resource.TestCheckResourceAttr("civicrm_case_status.test", "is_active", "true"),
					checkEntityAttr(t, "OptionValue", "civicrm_case_status.test", "id", "name", "tf_acc_case_status"),
					checkEntityAttr(t, "OptionValue", "civicrm_case_status.test", "id", "label", "TF Acceptance Case Status"),
					checkEntityAttr(t, "OptionValue", "civicrm_case_status.test", "id", "is_active", "true"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_case_status.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_case_status" "test" {
  name      = "tf_acc_case_status"
  label     = "TF Acceptance Case Status Updated"
  is_active = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_case_status.test", "label", "TF Acceptance Case Status Updated"),
					checkEntityAttr(t, "OptionValue", "civicrm_case_status.test", "id", "label", "TF Acceptance Case Status Updated"),
					checkEntityAttr(t, "OptionValue", "civicrm_case_status.test", "id", "name", "tf_acc_case_status"),
					checkEntityAttr(t, "OptionValue", "civicrm_case_status.test", "id", "is_active", "true"),
				),
			},
			// Deactivate: is_active=false must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_case_status" "test" {
  name      = "tf_acc_case_status"
  label     = "TF Acceptance Case Status Updated"
  is_active = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_case_status.test", "is_active", "false"),
					checkEntityAttr(t, "OptionValue", "civicrm_case_status.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
