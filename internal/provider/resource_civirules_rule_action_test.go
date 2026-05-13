package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccCiviRulesRuleActionResource tests attaching an action to a rule.
// action_id 1 is assumed to exist in the test CiviCRM instance.
func TestAccCiviRulesRuleActionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "CiviRulesRuleAction", "civicrm_civirules_rule_action"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + `
resource "civicrm_civirules_trigger" "action_test" {
  name       = "tf_acc_action_trigger"
  label      = "TF Acceptance Action Trigger"
  class_name = "CRM_Civirules_Trigger_Cron"
  cron       = true
}

resource "civicrm_civirules_rule" "action_test" {
  name       = "tf_acc_action_rule"
  label      = "TF Acceptance Action Rule"
  trigger_id = civicrm_civirules_trigger.action_test.id
  is_active  = true
}

resource "civicrm_civirules_rule_action" "test" {
  rule_id   = civicrm_civirules_rule.action_test.id
  action_id = 1
  is_active = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_civirules_rule_action.test", "id"),
					resource.TestCheckResourceAttr("civicrm_civirules_rule_action.test", "is_active", "true"),
					checkEntityAttr(t, "CiviRulesRuleAction", "civicrm_civirules_rule_action.test", "id", "is_active", "true"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_civirules_rule_action.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update is_active: must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_civirules_trigger" "action_test" {
  name       = "tf_acc_action_trigger"
  label      = "TF Acceptance Action Trigger"
  class_name = "CRM_Civirules_Trigger_Cron"
  cron       = true
}

resource "civicrm_civirules_rule" "action_test" {
  name       = "tf_acc_action_rule"
  label      = "TF Acceptance Action Rule"
  trigger_id = civicrm_civirules_trigger.action_test.id
  is_active  = true
}

resource "civicrm_civirules_rule_action" "test" {
  rule_id   = civicrm_civirules_rule.action_test.id
  action_id = 1
  is_active = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_civirules_rule_action.test", "is_active", "false"),
					checkEntityAttr(t, "CiviRulesRuleAction", "civicrm_civirules_rule_action.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
