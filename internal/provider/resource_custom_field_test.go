package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCustomFieldResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_custom_group" "field_test" {
  name       = "tf_acc_cf"
  title      = "TF Acceptance Custom Field Group"
  extends    = "Contact"
  table_name = "civicrm_value_tf_acc_cf"
}

resource "civicrm_custom_field" "test" {
  custom_group_id = civicrm_custom_group.field_test.id
  name            = "tf_acc_cf_field"
  label           = "TF Acceptance Custom Field"
  data_type       = "String"
  html_type       = "Text"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_custom_field.test", "id"),
					resource.TestCheckResourceAttr("civicrm_custom_field.test", "name", "tf_acc_cf_field"),
					resource.TestCheckResourceAttr("civicrm_custom_field.test", "label", "TF Acceptance Custom Field"),
					resource.TestCheckResourceAttr("civicrm_custom_field.test", "data_type", "String"),
					resource.TestCheckResourceAttr("civicrm_custom_field.test", "html_type", "Text"),
				),
			},
			{
				ResourceName:      "civicrm_custom_field.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_custom_group" "field_test" {
  name       = "tf_acc_cf"
  title      = "TF Acceptance Custom Field Group"
  extends    = "Contact"
  table_name = "civicrm_value_tf_acc_cf"
}

resource "civicrm_custom_field" "test" {
  custom_group_id = civicrm_custom_group.field_test.id
  name            = "tf_acc_cf_field"
  label           = "TF Acceptance Custom Field Updated"
  data_type       = "String"
  html_type       = "Text"
}`,
				Check: resource.TestCheckResourceAttr("civicrm_custom_field.test", "label", "TF Acceptance Custom Field Updated"),
			},
		},
	})
}
