package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMailSettingsResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_mail_settings" "test" {
  name      = "tf_acc_mail_settings"
  domain    = "tf-acc-test.example.org"
  localpart = "tf-acc"
  protocol  = "IMAP"
  server    = "imap.tf-acc-test.example.org"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_mail_settings.test", "id"),
					resource.TestCheckResourceAttr("civicrm_mail_settings.test", "name", "tf_acc_mail_settings"),
					resource.TestCheckResourceAttr("civicrm_mail_settings.test", "domain", "tf-acc-test.example.org"),
					resource.TestCheckResourceAttr("civicrm_mail_settings.test", "protocol", "IMAP"),
				),
			},
			{
				ResourceName:      "civicrm_mail_settings.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: providerConfig() + `
resource "civicrm_mail_settings" "test" {
  name      = "tf_acc_mail_settings"
  domain    = "tf-acc-test-updated.example.org"
  localpart = "tf-acc"
  protocol  = "IMAP"
  server    = "imap.tf-acc-test.example.org"
}`,
				Check: resource.TestCheckResourceAttr("civicrm_mail_settings.test", "domain", "tf-acc-test-updated.example.org"),
			},
		},
	})
}
