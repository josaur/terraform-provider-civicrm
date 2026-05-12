package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCiviRulesRuleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
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
  label      = "TF Acceptance CiviRules Rule"
  trigger_id = civicrm_civirules_trigger.rule_test.id
  is_active  = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_civirules_rule.test", "id"),
					resource.TestCheckResourceAttr("civicrm_civirules_rule.test", "name", "tf_acc_civirules_rule"),
					resource.TestCheckResourceAttr("civicrm_civirules_rule.test", "label", "TF Acceptance CiviRules Rule"),
					resource.TestCheckResourceAttr("civicrm_civirules_rule.test", "is_active", "true"),
				),
			},
			{
				ResourceName:      "civicrm_civirules_rule.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
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
  is_active  = true
}`,
				Check: resource.TestCheckResourceAttr("civicrm_civirules_rule.test", "label", "TF Acceptance CiviRules Rule Updated"),
			},
		},
	})
}
