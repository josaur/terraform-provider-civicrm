package provider_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// checkConditionParamsIsPHPSerialized is the CiviRulesRuleCondition analogue
// of checkActionParamsIsPHPSerialized in resource_civirules_rule_action_test.go.
func checkConditionParamsIsPHPSerialized(t *testing.T, resourceAddr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		t.Helper()
		rs, ok := s.RootModule().Resources[resourceAddr]
		if !ok {
			t.Fatalf("resource %q not found in state", resourceAddr)
		}
		id, err := strconv.ParseInt(rs.Primary.Attributes["id"], 10, 64)
		if err != nil {
			t.Fatalf("resource %q has no valid id: %s", resourceAddr, err)
		}

		client := newTestClient(t)
		result, err := client.GetByID("CiviRulesRuleCondition", id, []string{"condition_params"})
		if err != nil {
			t.Fatalf("GetByID(CiviRulesRuleCondition, %d): %s", id, err)
		}

		raw, ok := result["condition_params"]
		if !ok || raw == nil {
			t.Fatalf("CiviRulesRuleCondition %d has no condition_params", id)
		}
		str, ok := raw.(string)
		if !ok {
			t.Fatalf("CiviRulesRuleCondition %d condition_params is %T, not a string: %#v", id, raw, raw)
		}
		if !strings.HasPrefix(str, "a:") {
			t.Fatalf("CiviRulesRuleCondition %d condition_params = %q, want PHP serialize() data (prefix \"a:\")", id, str)
		}
		return nil
	}
}

// TestAccCiviRulesRuleConditionResource tests attaching a condition to a rule.
// condition_id 1 is assumed to exist in the test CiviCRM instance.
func TestAccCiviRulesRuleConditionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "CiviRulesRuleCondition", "civicrm_civirules_rule_condition"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
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
					checkEntityAttr(t, "CiviRulesRuleCondition", "civicrm_civirules_rule_condition.test", "id", "is_active", "true"),
				),
			},
			// Import
			{
				ResourceName:            "civicrm_civirules_rule_condition.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"negate"},
			},
			// Update is_active and negate: must be persisted in DB.
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
  is_active    = false
  negate       = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_civirules_rule_condition.test", "is_active", "false"),
					checkEntityAttr(t, "CiviRulesRuleCondition", "civicrm_civirules_rule_condition.test", "id", "is_active", "false"),
				),
			},
			// condition_params as JSON: must be readable back as the same
			// JSON, and stored in CiviCRM as PHP serialize() data (not raw
			// JSON) so CiviRules' own unserialize() call can read it at
			// runtime.
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
  condition_params = jsonencode({
    activity_type_id = "Follow up"
  })
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_civirules_rule_condition.test", "condition_params", `{"activity_type_id":"Follow up"}`),
					checkConditionParamsIsPHPSerialized(t, "civicrm_civirules_rule_condition.test"),
				),
			},
			// Import must survive the JSON <-> PHP serialize() round trip.
			{
				ResourceName:            "civicrm_civirules_rule_condition.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"negate"},
			},
		},
	})
}
