package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMailSettingsResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "MailSettings", "civicrm_mail_settings"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + `
resource "civicrm_mail_settings" "test" {
  name      = "tf_acc_mail_settings"
  domain    = "tf-acc-test.example.org"
  localpart = "tf-acc"
  protocol  = "IMAP"
  server    = "imap.tf-acc-test.example.org"
  is_active = true
  is_ssl    = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_mail_settings.test", "id"),
					resource.TestCheckResourceAttr("civicrm_mail_settings.test", "name", "tf_acc_mail_settings"),
					resource.TestCheckResourceAttr("civicrm_mail_settings.test", "domain", "tf-acc-test.example.org"),
					resource.TestCheckResourceAttr("civicrm_mail_settings.test", "protocol", "IMAP"),
					resource.TestCheckResourceAttr("civicrm_mail_settings.test", "is_active", "true"),
					checkEntityAttr(t, "MailSettings", "civicrm_mail_settings.test", "id", "name", "tf_acc_mail_settings"),
					checkEntityAttr(t, "MailSettings", "civicrm_mail_settings.test", "id", "domain", "tf-acc-test.example.org"),
					checkEntityAttr(t, "MailSettings", "civicrm_mail_settings.test", "id", "localpart", "tf-acc"),
					checkEntityAttr(t, "MailSettings", "civicrm_mail_settings.test", "id", "protocol", "IMAP"),
					checkEntityAttr(t, "MailSettings", "civicrm_mail_settings.test", "id", "server", "imap.tf-acc-test.example.org"),
					checkEntityAttr(t, "MailSettings", "civicrm_mail_settings.test", "id", "is_active", "true"),
					checkEntityAttr(t, "MailSettings", "civicrm_mail_settings.test", "id", "is_ssl", "false"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_mail_settings.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_mail_settings" "test" {
  name      = "tf_acc_mail_settings"
  domain    = "tf-acc-test-updated.example.org"
  localpart = "tf-acc-updated"
  protocol  = "IMAP"
  server    = "imap.tf-acc-test.example.org"
  is_active = true
  is_ssl    = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_mail_settings.test", "domain", "tf-acc-test-updated.example.org"),
					checkEntityAttr(t, "MailSettings", "civicrm_mail_settings.test", "id", "domain", "tf-acc-test-updated.example.org"),
					checkEntityAttr(t, "MailSettings", "civicrm_mail_settings.test", "id", "localpart", "tf-acc-updated"),
					checkEntityAttr(t, "MailSettings", "civicrm_mail_settings.test", "id", "is_active", "true"),
					checkEntityAttr(t, "MailSettings", "civicrm_mail_settings.test", "id", "is_ssl", "false"),
				),
			},
			// Deactivate: is_active=false must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_mail_settings" "test" {
  name      = "tf_acc_mail_settings"
  domain    = "tf-acc-test-updated.example.org"
  localpart = "tf-acc-updated"
  protocol  = "IMAP"
  server    = "imap.tf-acc-test.example.org"
  is_active = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_mail_settings.test", "is_active", "false"),
					checkEntityAttr(t, "MailSettings", "civicrm_mail_settings.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
