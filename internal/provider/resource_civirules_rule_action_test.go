package provider_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// checkActionParamsIsPHPSerialized asserts that CiviRulesRuleAction.action_params,
// as stored in CiviCRM, is PHP serialize() data (starts with the "a:" array
// marker) rather than a raw JSON object — the exact bug this provider fixes:
// CiviRules reads action_params with PHP's unserialize(), which silently
// fails on a JSON string. A round-trip through this provider's own Read
// path cannot catch a regression here, since Read decodes whatever format
// is stored; this check bypasses the provider and inspects the raw API
// response directly.
func checkActionParamsIsPHPSerialized(t *testing.T, resourceAddr string) resource.TestCheckFunc {
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
		result, err := client.GetByID("CiviRulesRuleAction", id, []string{"action_params"})
		if err != nil {
			t.Fatalf("GetByID(CiviRulesRuleAction, %d): %s", id, err)
		}

		raw, ok := result["action_params"]
		if !ok || raw == nil {
			t.Fatalf("CiviRulesRuleAction %d has no action_params", id)
		}
		str, ok := raw.(string)
		if !ok {
			t.Fatalf("CiviRulesRuleAction %d action_params is %T, not a string: %#v", id, raw, raw)
		}
		if !strings.HasPrefix(str, "a:") {
			t.Fatalf("CiviRulesRuleAction %d action_params = %q, want PHP serialize() data (prefix \"a:\")", id, str)
		}
		return nil
	}
}

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
			// action_params as JSON: must be readable back as the same JSON,
			// and stored in CiviCRM as PHP serialize() data (not raw JSON) so
			// CiviRules' own unserialize() call can read it at runtime.
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
  action_params = jsonencode({
    status_id = 2
    tag_ids   = [1, 2]
  })
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_civirules_rule_action.test", "action_params", `{"status_id":2,"tag_ids":[1,2]}`),
					checkActionParamsIsPHPSerialized(t, "civicrm_civirules_rule_action.test"),
				),
			},
			// Import must survive the JSON <-> PHP serialize() round trip.
			{
				ResourceName:      "civicrm_civirules_rule_action.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
