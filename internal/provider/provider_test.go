package provider_test

import (
	"os"
	"testing"

	"github.com/Caritas-Deutschland-Digitallabor/civicrm-terraform/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"civicrm": providerserver.NewProtocol6WithError(provider.New("test")()),
}

func testAccPreCheck(t *testing.T) {
	t.Helper()
	if os.Getenv("CIVICRM_URL") == "" {
		t.Fatal("CIVICRM_URL must be set for acceptance tests")
	}
	if os.Getenv("CIVICRM_API_KEY") == "" {
		t.Fatal("CIVICRM_API_KEY must be set for acceptance tests")
	}
}

func providerConfig() string {
	return `provider "civicrm" {}`
}

func TestAccProvider(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig(),
			},
		},
	})
}
