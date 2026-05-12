package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCaseStatusResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_case_status" "test" {
  name  = "tf_acc_case_status"
  label = "TF Acceptance Case Status"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_case_status.test", "id"),
					resource.TestCheckResourceAttr("civicrm_case_status.test", "name", "tf_acc_case_status"),
					resource.TestCheckResourceAttr("civicrm_case_status.test", "label", "TF Acceptance Case Status"),
					resource.TestCheckResourceAttr("civicrm_case_status.test", "is_active", "true"),
				),
			},
			{
				ResourceName:      "civicrm_case_status.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_case_status" "test" {
  name  = "tf_acc_case_status"
  label = "TF Acceptance Case Status Updated"
}`,
				Check: resource.TestCheckResourceAttr("civicrm_case_status.test", "label", "TF Acceptance Case Status Updated"),
			},
		},
	})
}
