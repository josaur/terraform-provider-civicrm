package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCiviRulesTriggerResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_civirules_trigger" "test" {
  name       = "tf_acc_civirules_trigger"
  label      = "TF Acceptance CiviRules Trigger"
  class_name = "CRM_Civirules_Trigger_Cron"
  cron       = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_civirules_trigger.test", "id"),
					resource.TestCheckResourceAttr("civicrm_civirules_trigger.test", "name", "tf_acc_civirules_trigger"),
					resource.TestCheckResourceAttr("civicrm_civirules_trigger.test", "label", "TF Acceptance CiviRules Trigger"),
					resource.TestCheckResourceAttr("civicrm_civirules_trigger.test", "cron", "true"),
				),
			},
			{
				ResourceName:      "civicrm_civirules_trigger.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_civirules_trigger" "test" {
  name       = "tf_acc_civirules_trigger"
  label      = "TF Acceptance CiviRules Trigger Updated"
  class_name = "CRM_Civirules_Trigger_Cron"
  cron       = true
}`,
				Check: resource.TestCheckResourceAttr("civicrm_civirules_trigger.test", "label", "TF Acceptance CiviRules Trigger Updated"),
			},
		},
	})
}
