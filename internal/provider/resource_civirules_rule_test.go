package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCiviRulesRuleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "CiviRulesRule", "civicrm_civirules_rule"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + `
resource "civicrm_civirules_trigger" "rule_test" {
  name       = "tf_acc_rule_trigger"
  label      = "TF Acceptance Rule Trigger"
  class_name = "CRM_Civirules_Trigger_Cron"
  cron       = true
}

resource "civicrm_civirules_rule" "test" {
  name        = "tf_acc_civirules_rule"
  label       = "TF Acceptance CiviRules Rule"
  description = "tf acc civirules rule description"
  trigger_id  = civicrm_civirules_trigger.rule_test.id
  is_active   = true
  is_debug    = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_civirules_rule.test", "id"),
					resource.TestCheckResourceAttr("civicrm_civirules_rule.test", "name", "tf_acc_civirules_rule"),
					resource.TestCheckResourceAttr("civicrm_civirules_rule.test", "label", "TF Acceptance CiviRules Rule"),
					resource.TestCheckResourceAttr("civicrm_civirules_rule.test", "is_active", "true"),
					checkEntityAttr(t, "CiviRulesRule", "civicrm_civirules_rule.test", "id", "name", "tf_acc_civirules_rule"),
					checkEntityAttr(t, "CiviRulesRule", "civicrm_civirules_rule.test", "id", "label", "TF Acceptance CiviRules Rule"),
					checkEntityAttr(t, "CiviRulesRule", "civicrm_civirules_rule.test", "id", "description", "tf acc civirules rule description"),
					checkEntityAttr(t, "CiviRulesRule", "civicrm_civirules_rule.test", "id", "is_active", "true"),
					checkEntityAttr(t, "CiviRulesRule", "civicrm_civirules_rule.test", "id", "is_debug", "false"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_civirules_rule.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_civirules_trigger" "rule_test" {
  name       = "tf_acc_rule_trigger"
  label      = "TF Acceptance Rule Trigger"
  class_name = "CRM_Civirules_Trigger_Cron"
  cron       = true
}

resource "civicrm_civirules_rule" "test" {
  name        = "tf_acc_civirules_rule"
  label       = "TF Acceptance CiviRules Rule Updated"
  description = "tf acc civirules rule description updated"
  trigger_id  = civicrm_civirules_trigger.rule_test.id
  is_active   = true
  is_debug    = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_civirules_rule.test", "label", "TF Acceptance CiviRules Rule Updated"),
					checkEntityAttr(t, "CiviRulesRule", "civicrm_civirules_rule.test", "id", "label", "TF Acceptance CiviRules Rule Updated"),
					checkEntityAttr(t, "CiviRulesRule", "civicrm_civirules_rule.test", "id", "description", "tf acc civirules rule description updated"),
					checkEntityAttr(t, "CiviRulesRule", "civicrm_civirules_rule.test", "id", "is_active", "true"),
					checkEntityAttr(t, "CiviRulesRule", "civicrm_civirules_rule.test", "id", "is_debug", "false"),
				),
			},
			// Deactivate: is_active=false must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_civirules_trigger" "rule_test" {
  name       = "tf_acc_rule_trigger"
  label      = "TF Acceptance Rule Trigger"
  class_name = "CRM_Civirules_Trigger_Cron"
  cron       = true
}

resource "civicrm_civirules_rule" "test" {
  name       = "tf_acc_civirules_rule"
  label      = "TF Acceptance CiviRules Rule Updated"
  trigger_id = civicrm_civirules_trigger.rule_test.id
  is_active  = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_civirules_rule.test", "is_active", "false"),
					checkEntityAttr(t, "CiviRulesRule", "civicrm_civirules_rule.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
