package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccCiviRulesRuleConditionResource tests attaching a condition to a rule.
// It depends on CiviRules being installed on the CiviCRM test instance and requires
// a valid condition_id. The condition_id 1 is assumed to exist; adjust if needed.
func TestAccCiviRulesRuleConditionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_civirules_trigger" "cond_test" {
  name       = "tf_acc_cond_trigger"
  label      = "TF Acceptance Condition Trigger"
  class_name = "CRM_Civirules_Trigger_Cron"
  cron       = true
}

resource "civicrm_civirules_rule" "cond_test" {
  name       = "tf_acc_cond_rule"
  label      = "TF Acceptance Condition Rule"
  trigger_id = civicrm_civirules_trigger.cond_test.id
  is_active  = true
}

resource "civicrm_civirules_rule_condition" "test" {
  rule_id      = civicrm_civirules_rule.cond_test.id
  condition_id = 1
  is_active    = true
  negate       = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_civirules_rule_condition.test", "id"),
					resource.TestCheckResourceAttr("civicrm_civirules_rule_condition.test", "is_active", "true"),
					resource.TestCheckResourceAttr("civicrm_civirules_rule_condition.test", "negate", "false"),
				),
			},
			{
				ResourceName:            "civicrm_civirules_rule_condition.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"negate"}, // CiviCRM may not return negate=false explicitly
			},
		},
	})
}
