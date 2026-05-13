package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSiteEmailAddressResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "SiteEmailAddress", "civicrm_site_email_address"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + `
resource "civicrm_site_email_address" "test" {
  display_name = "TF Acceptance Test"
  email        = "tf-acc-test@example.org"
  is_active    = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_site_email_address.test", "id"),
					resource.TestCheckResourceAttr("civicrm_site_email_address.test", "display_name", "TF Acceptance Test"),
					resource.TestCheckResourceAttr("civicrm_site_email_address.test", "email", "tf-acc-test@example.org"),
					resource.TestCheckResourceAttr("civicrm_site_email_address.test", "is_active", "true"),
					checkEntityAttr(t, "SiteEmailAddress", "civicrm_site_email_address.test", "id", "display_name", "TF Acceptance Test"),
					checkEntityAttr(t, "SiteEmailAddress", "civicrm_site_email_address.test", "id", "email", "tf-acc-test@example.org"),
					checkEntityAttr(t, "SiteEmailAddress", "civicrm_site_email_address.test", "id", "is_active", "true"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_site_email_address.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_site_email_address" "test" {
  display_name = "TF Acceptance Test Updated"
  email        = "tf-acc-test@example.org"
  is_active    = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_site_email_address.test", "display_name", "TF Acceptance Test Updated"),
					checkEntityAttr(t, "SiteEmailAddress", "civicrm_site_email_address.test", "id", "display_name", "TF Acceptance Test Updated"),
					checkEntityAttr(t, "SiteEmailAddress", "civicrm_site_email_address.test", "id", "email", "tf-acc-test@example.org"),
					checkEntityAttr(t, "SiteEmailAddress", "civicrm_site_email_address.test", "id", "is_active", "true"),
				),
			},
			// Deactivate: is_active=false must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_site_email_address" "test" {
  display_name = "TF Acceptance Test Updated"
  email        = "tf-acc-test@example.org"
  is_active    = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_site_email_address.test", "is_active", "false"),
					checkEntityAttr(t, "SiteEmailAddress", "civicrm_site_email_address.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
