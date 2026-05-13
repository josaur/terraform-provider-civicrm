package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const tfAccCustomFieldGroupConfig = `
resource "civicrm_custom_group" "field_test" {
  name       = "tf_acc_cf"
  title      = "TF Acceptance Custom Field Group"
  extends    = "Contact"
  table_name = "civicrm_value_tf_acc_cf"
}
`

func TestAccCustomFieldResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "CustomField", "civicrm_custom_field"),
		Steps: []resource.TestStep{
			// Create: verify state and DB.
			{
				Config: providerConfig() + tfAccCustomFieldGroupConfig + `
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
					// column_name is auto-generated; must be non-empty after create
					resource.TestCheckResourceAttrSet("civicrm_custom_field.test", "column_name"),
					// DB verification
					checkEntityAttr(t, "CustomField", "civicrm_custom_field.test", "id", "label", "TF Acceptance Custom Field"),
					checkEntityAttr(t, "CustomField", "civicrm_custom_field.test", "id", "data_type", "String"),
					checkEntityAttr(t, "CustomField", "civicrm_custom_field.test", "id", "html_type", "Text"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_custom_field.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update label; column_name must remain unchanged (write-once via UseStateForUnknown).
			{
				Config: providerConfig() + tfAccCustomFieldGroupConfig + `
resource "civicrm_custom_field" "test" {
  custom_group_id = civicrm_custom_group.field_test.id
  name            = "tf_acc_cf_field"
  label           = "TF Acceptance Custom Field Updated"
  data_type       = "String"
  html_type       = "Text"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_custom_field.test", "label", "TF Acceptance Custom Field Updated"),
					// column_name must still be non-empty and unchanged
					resource.TestCheckResourceAttrSet("civicrm_custom_field.test", "column_name"),
					// DB verification after update
					checkEntityAttr(t, "CustomField", "civicrm_custom_field.test", "id", "label", "TF Acceptance Custom Field Updated"),
				),
			},
		},
	})
}

// TestAccCustomFieldResourceColumnNameStability explicitly supplies a
// column_name on create and then updates the field, verifying that CiviCRM
// has not changed the column name in the database.
func TestAccCustomFieldResourceColumnNameStability(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "CustomField", "civicrm_custom_field"),
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_custom_group" "coltest" {
  name       = "tf_acc_cf_col"
  title      = "TF Acceptance Column Name Stability"
  extends    = "Contact"
  table_name = "civicrm_value_tf_acc_cf_col"
}

resource "civicrm_custom_field" "coltest" {
  custom_group_id = civicrm_custom_group.coltest.id
  name            = "tf_acc_col_field"
  label           = "Column Name Stability Field"
  data_type       = "String"
  html_type       = "Text"
  column_name     = "tf_stable_col"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_custom_field.coltest", "id"),
					resource.TestCheckResourceAttr("civicrm_custom_field.coltest", "column_name", "tf_stable_col"),
					checkEntityAttr(t, "CustomField", "civicrm_custom_field.coltest", "id", "column_name", "tf_stable_col"),
				),
			},
			// Update label but omit column_name from config; UseStateForUnknown keeps it.
			{
				Config: providerConfig() + `
resource "civicrm_custom_group" "coltest" {
  name       = "tf_acc_cf_col"
  title      = "TF Acceptance Column Name Stability"
  extends    = "Contact"
  table_name = "civicrm_value_tf_acc_cf_col"
}

resource "civicrm_custom_field" "coltest" {
  custom_group_id = civicrm_custom_group.coltest.id
  name            = "tf_acc_col_field"
  label           = "Column Name Stability Field Updated"
  data_type       = "String"
  html_type       = "Text"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_custom_field.coltest", "label", "Column Name Stability Field Updated"),
					// column_name must still equal the value from the initial create
					resource.TestCheckResourceAttr("civicrm_custom_field.coltest", "column_name", "tf_stable_col"),
					checkEntityAttr(t, "CustomField", "civicrm_custom_field.coltest", "id", "column_name", "tf_stable_col"),
				),
			},
		},
	})
}
