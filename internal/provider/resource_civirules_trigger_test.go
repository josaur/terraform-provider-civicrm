package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCiviRulesTriggerResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "CiviRulesTrigger", "civicrm_civirules_trigger"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + `
resource "civicrm_civirules_trigger" "test" {
  name       = "tf_acc_civirules_trigger"
  label      = "TF Acceptance CiviRules Trigger"
  class_name = "CRM_Civirules_Trigger_Cron"
  cron       = true
  is_active  = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_civirules_trigger.test", "id"),
					resource.TestCheckResourceAttr("civicrm_civirules_trigger.test", "name", "tf_acc_civirules_trigger"),
					resource.TestCheckResourceAttr("civicrm_civirules_trigger.test", "label", "TF Acceptance CiviRules Trigger"),
					resource.TestCheckResourceAttr("civicrm_civirules_trigger.test", "cron", "true"),
					resource.TestCheckResourceAttr("civicrm_civirules_trigger.test", "is_active", "true"),
					checkEntityAttr(t, "CiviRulesTrigger", "civicrm_civirules_trigger.test", "id", "name", "tf_acc_civirules_trigger"),
					checkEntityAttr(t, "CiviRulesTrigger", "civicrm_civirules_trigger.test", "id", "label", "TF Acceptance CiviRules Trigger"),
					checkEntityAttr(t, "CiviRulesTrigger", "civicrm_civirules_trigger.test", "id", "cron", "true"),
					checkEntityAttr(t, "CiviRulesTrigger", "civicrm_civirules_trigger.test", "id", "is_active", "true"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_civirules_trigger.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_civirules_trigger" "test" {
  name       = "tf_acc_civirules_trigger"
  label      = "TF Acceptance CiviRules Trigger Updated"
  class_name = "CRM_Civirules_Trigger_Cron"
  cron       = true
  is_active  = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_civirules_trigger.test", "label", "TF Acceptance CiviRules Trigger Updated"),
					checkEntityAttr(t, "CiviRulesTrigger", "civicrm_civirules_trigger.test", "id", "label", "TF Acceptance CiviRules Trigger Updated"),
					checkEntityAttr(t, "CiviRulesTrigger", "civicrm_civirules_trigger.test", "id", "cron", "true"),
					checkEntityAttr(t, "CiviRulesTrigger", "civicrm_civirules_trigger.test", "id", "is_active", "true"),
				),
			},
			// Deactivate: is_active=false must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_civirules_trigger" "test" {
  name       = "tf_acc_civirules_trigger"
  label      = "TF Acceptance CiviRules Trigger Updated"
  class_name = "CRM_Civirules_Trigger_Cron"
  cron       = true
  is_active  = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_civirules_trigger.test", "is_active", "false"),
					checkEntityAttr(t, "CiviRulesTrigger", "civicrm_civirules_trigger.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
