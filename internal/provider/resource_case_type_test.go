package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCaseTypeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_case_type" "test" {
  name  = "tf_acc_case_type"
  title = "TF Acceptance Case Type"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_case_type.test", "id"),
					resource.TestCheckResourceAttr("civicrm_case_type.test", "name", "tf_acc_case_type"),
					resource.TestCheckResourceAttr("civicrm_case_type.test", "title", "TF Acceptance Case Type"),
					resource.TestCheckResourceAttr("civicrm_case_type.test", "is_active", "true"),
				),
			},
			{
				ResourceName:      "civicrm_case_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_case_type" "test" {
  name  = "tf_acc_case_type"
  title = "TF Acceptance Case Type Updated"
}`,
				Check: resource.TestCheckResourceAttr("civicrm_case_type.test", "title", "TF Acceptance Case Type Updated"),
			},
		},
	})
}
