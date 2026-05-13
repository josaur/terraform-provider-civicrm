package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccActivityTypeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "OptionValue", "civicrm_activity_type"),
		Steps: []resource.TestStep{
			// Create: verify state and DB.
			{
				Config: providerConfig() + `
resource "civicrm_activity_type" "test" {
  name  = "tf_acc_activity_type"
  label = "TF Acceptance Activity Type"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_activity_type.test", "id"),
					resource.TestCheckResourceAttr("civicrm_activity_type.test", "name", "tf_acc_activity_type"),
					resource.TestCheckResourceAttr("civicrm_activity_type.test", "label", "TF Acceptance Activity Type"),
					resource.TestCheckResourceAttr("civicrm_activity_type.test", "is_active", "true"),
					// value is auto-assigned by CiviCRM; must be non-empty
					resource.TestCheckResourceAttrSet("civicrm_activity_type.test", "value"),
					// DB verification
					checkEntityAttr(t, "OptionValue", "civicrm_activity_type.test", "id", "label", "TF Acceptance Activity Type"),
					checkEntityAttr(t, "OptionValue", "civicrm_activity_type.test", "id", "name", "tf_acc_activity_type"),
					checkEntityAttr(t, "OptionValue", "civicrm_activity_type.test", "id", "is_active", "true"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_activity_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update label; the auto-assigned `value` must remain stable (UseStateForUnknown).
			{
				Config: providerConfig() + `
resource "civicrm_activity_type" "test" {
  name  = "tf_acc_activity_type"
  label = "TF Acceptance Activity Type Updated"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_activity_type.test", "label", "TF Acceptance Activity Type Updated"),
					// value must still be set and unchanged
					resource.TestCheckResourceAttrSet("civicrm_activity_type.test", "value"),
					// DB verification after update
					checkEntityAttr(t, "OptionValue", "civicrm_activity_type.test", "id", "label", "TF Acceptance Activity Type Updated"),
				),
			},
		},
	})
}

// TestAccActivityTypeResourceValueStability verifies that the `value` field
// (UseStateForUnknown) set on create is never overwritten by subsequent updates.
func TestAccActivityTypeResourceValueStability(t *testing.T) {
	var capturedValue string

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "OptionValue", "civicrm_activity_type"),
		Steps: []resource.TestStep{
			// Create without explicit value; capture the auto-assigned one.
			{
				Config: providerConfig() + `
resource "civicrm_activity_type" "stable" {
  name  = "tf_acc_at_stable"
  label = "TF Value Stability"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_activity_type.stable", "value"),
					func(s *terraform.State) error {
						rs, ok := s.RootModule().Resources["civicrm_activity_type.stable"]
						if ok {
							capturedValue = rs.Primary.Attributes["value"]
						}
						return nil
					},
				),
			},
			// Update description only; value must equal what was captured above.
			{
				Config: providerConfig() + `
resource "civicrm_activity_type" "stable" {
  name        = "tf_acc_at_stable"
  label       = "TF Value Stability"
  description = "updated description"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrWith("civicrm_activity_type.stable", "value", func(got string) error {
						if got != capturedValue {
							return fmt.Errorf("value changed after update: was %q, now %q", capturedValue, got)
						}
						return nil
					}),
					// DB must store the same value
					checkEntityAttr(t, "OptionValue", "civicrm_activity_type.stable", "id", "description", "updated description"),
				),
			},
		},
	})
}
