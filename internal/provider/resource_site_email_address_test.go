package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSiteEmailAddressResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_site_email_address" "test" {
  display_name = "TF Acceptance Test"
  email        = "tf-acc-test@example.org"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_site_email_address.test", "id"),
					resource.TestCheckResourceAttr("civicrm_site_email_address.test", "display_name", "TF Acceptance Test"),
					resource.TestCheckResourceAttr("civicrm_site_email_address.test", "email", "tf-acc-test@example.org"),
					resource.TestCheckResourceAttr("civicrm_site_email_address.test", "is_active", "true"),
				),
			},
			{
				ResourceName:      "civicrm_site_email_address.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_site_email_address" "test" {
  display_name = "TF Acceptance Test Updated"
  email        = "tf-acc-test@example.org"
}`,
				Check: resource.TestCheckResourceAttr("civicrm_site_email_address.test", "display_name", "TF Acceptance Test Updated"),
			},
		},
	})
}
