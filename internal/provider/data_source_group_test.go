package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGroupDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_group" "ds_test" {
  name  = "tf_acc_ds_group"
  title = "TF Acceptance DS Group"
}

data "civicrm_group" "test" {
  name = civicrm_group.ds_test.name
  depends_on = [civicrm_group.ds_test]
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.civicrm_group.test", "id"),
					resource.TestCheckResourceAttr("data.civicrm_group.test", "name", "tf_acc_ds_group"),
					resource.TestCheckResourceAttr("data.civicrm_group.test", "title", "TF Acceptance DS Group"),
				),
			},
		},
	})
}
