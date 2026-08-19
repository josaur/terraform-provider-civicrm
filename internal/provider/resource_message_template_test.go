package provider_test

import (
	"strconv"
	"testing"

	"github.com/Caritas-Deutschland-Digitallabor/civicrm-terraform/internal/provider"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccMessageTemplateResource_ImportPreExisting verifies the documented
// best-practice path for adopting a MessageTemplate that already exists in
// CiviCRM (e.g. one of CiviCRM's built-in workflow templates such as the
// Invoice template): import it by ID rather than letting Terraform create a
// duplicate. This is standard `terraform import`, exercised the same way
// TestAccMessageTemplateResource below exercises import for a
// Terraform-created resource — the only difference here is that the row
// exists *before* any Terraform config references it.
func TestAccMessageTemplateResource_ImportPreExisting(t *testing.T) {
	const workflowName = "tf_acc_import_workflow"

	client := newTestClient(t)
	preExisting, err := client.Create("MessageTemplate", map[string]any{
		"msg_title":     "tf_acc_import_pre_existing",
		"workflow_name": workflowName,
		"is_reserved":   false,
		"msg_subject":   "pre-existing subject",
	})
	if err != nil {
		t.Fatalf("failed to seed pre-existing MessageTemplate: %s", err)
	}
	preExistingID, ok := provider.GetInt64(preExisting, "id")
	if !ok {
		t.Fatalf("seeded MessageTemplate has no id: %v", preExisting)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "MessageTemplate", "civicrm_message_template"),
		Steps: []resource.TestStep{
			// Import the pre-existing row into a resource address that matches it,
			// using a Terraform 1.5+ `import` block (ImportBlockWithID) rather than
			// the legacy `terraform import` command: `import` blocks plan and apply
			// together with the resource config in a single `terraform apply`, which
			// is the actual documented workflow (see message_template.md) — unlike
			// the legacy command-based import test mode, it cannot accidentally
			// create a duplicate if misconfigured, since there is only ever one
			// apply, not a separate import-then-compare pass.
			{
				Config: providerConfig() + `
import {
  to = civicrm_message_template.imported
  id = "` + strconv.FormatInt(preExistingID, 10) + `"
}

resource "civicrm_message_template" "imported" {
  msg_title     = "tf_acc_import_pre_existing"
  workflow_name = "` + workflowName + `"
  msg_subject   = "pre-existing subject"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_message_template.imported", "id", strconv.FormatInt(preExistingID, 10)),
					resource.TestCheckResourceAttr("civicrm_message_template.imported", "msg_subject", "pre-existing subject"),
					checkEntityAttr(t, "MessageTemplate", "civicrm_message_template.imported", "id", "msg_subject", "pre-existing subject"),
				),
			},
			// Once imported, it's managed like any other resource: updates apply,
			// and the eventual `terraform destroy` (driven by CheckDestroy below)
			// removes it from CiviCRM like normal — no special adopt/protect
			// behavior, per the standard import workflow.
			{
				Config: providerConfig() + `
resource "civicrm_message_template" "imported" {
  msg_title     = "tf_acc_import_pre_existing"
  workflow_name = "` + workflowName + `"
  msg_subject   = "updated after import"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_message_template.imported", "msg_subject", "updated after import"),
					checkEntityAttr(t, "MessageTemplate", "civicrm_message_template.imported", "id", "msg_subject", "updated after import"),
				),
			},
		},
	})
}

func TestAccMessageTemplateResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "MessageTemplate", "civicrm_message_template"),
		Steps: []resource.TestStep{
			// Create: verify all set fields land in DB.
			{
				Config: providerConfig() + `
resource "civicrm_message_template" "test" {
  msg_title   = "tf_acc_message_template"
  msg_subject = "TF Acceptance Test Subject"
  msg_html    = "<p>Hello {contact.display_name}</p>"
  msg_text    = "Hello {contact.display_name}"
  is_active   = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_message_template.test", "id"),
					resource.TestCheckResourceAttr("civicrm_message_template.test", "msg_title", "tf_acc_message_template"),
					resource.TestCheckResourceAttr("civicrm_message_template.test", "msg_subject", "TF Acceptance Test Subject"),
					resource.TestCheckResourceAttr("civicrm_message_template.test", "is_active", "true"),
					checkEntityAttr(t, "MessageTemplate", "civicrm_message_template.test", "id", "msg_title", "tf_acc_message_template"),
					checkEntityAttr(t, "MessageTemplate", "civicrm_message_template.test", "id", "msg_subject", "TF Acceptance Test Subject"),
					checkEntityAttr(t, "MessageTemplate", "civicrm_message_template.test", "id", "is_active", "true"),
				),
			},
			// Import
			{
				ResourceName:      "civicrm_message_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: all changed fields must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_message_template" "test" {
  msg_title   = "tf_acc_message_template"
  msg_subject = "TF Acceptance Test Subject Updated"
  msg_html    = "<p>Hi {contact.display_name}</p>"
  msg_text    = "Hi {contact.display_name}"
  is_active   = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_message_template.test", "msg_subject", "TF Acceptance Test Subject Updated"),
					checkEntityAttr(t, "MessageTemplate", "civicrm_message_template.test", "id", "msg_subject", "TF Acceptance Test Subject Updated"),
					checkEntityAttr(t, "MessageTemplate", "civicrm_message_template.test", "id", "msg_html", "<p>Hi {contact.display_name}</p>"),
				),
			},
			// Set is_active=false: must be persisted in DB.
			{
				Config: providerConfig() + `
resource "civicrm_message_template" "test" {
  msg_title   = "tf_acc_message_template"
  msg_subject = "TF Acceptance Test Subject Updated"
  msg_html    = "<p>Hi {contact.display_name}</p>"
  msg_text    = "Hi {contact.display_name}"
  is_active   = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_message_template.test", "is_active", "false"),
					checkEntityAttr(t, "MessageTemplate", "civicrm_message_template.test", "id", "is_active", "false"),
				),
			},
		},
	})
}
