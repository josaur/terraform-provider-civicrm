package provider_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Caritas-Deutschland-Digitallabor/civicrm-terraform/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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

// newTestClient creates a CiviCRM API client from the test environment variables.
func newTestClient(t *testing.T) *provider.Client {
	t.Helper()
	c, err := provider.NewClient(
		os.Getenv("CIVICRM_URL"),
		os.Getenv("CIVICRM_API_KEY"),
		os.Getenv("CIVICRM_INSECURE") == "true",
	)
	if err != nil {
		t.Fatalf("failed to create test client: %s", err)
	}
	return c
}

// checkDestroyByID returns a TestCheckFunc that verifies an entity no longer
// exists in the CiviCRM database after Terraform has destroyed it.
func checkDestroyByID(t *testing.T, entity, resourceType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := newTestClient(t)
		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceType {
				continue
			}
			var id int64
			fmt.Sscan(rs.Primary.ID, &id) //nolint:errcheck
			if id == 0 {
				continue
			}
			_, err := client.GetByID(entity, id, []string{"id"})
			if err == nil {
				return fmt.Errorf("%s %s (ID %d) still exists after destroy", entity, rs.Primary.ID, id)
			}
		}
		return nil
	}
}

// checkEntityAttr fetches the entity by the ID stored in the given resource
// attribute and asserts that the remote field equals the expected value.
func checkEntityAttr(t *testing.T, entity, resourceAddr, idAttr, field, expected string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceAddr]
		if !ok {
			return fmt.Errorf("resource %q not found in state", resourceAddr)
		}
		var id int64
		fmt.Sscan(rs.Primary.Attributes[idAttr], &id) //nolint:errcheck
		if id == 0 {
			return fmt.Errorf("resource %q has no %s", resourceAddr, idAttr)
		}
		client := newTestClient(t)
		result, err := client.GetByID(entity, id, []string{field})
		if err != nil {
			return fmt.Errorf("GetByID(%s, %d): %w", entity, id, err)
		}
		got, ok := provider.GetString(result, field)
		if !ok {
			// try bool
			b, ok2 := provider.GetBool(result, field)
			if !ok2 {
				return fmt.Errorf("field %q not found in API response for %s %d", field, entity, id)
			}
			if expected == "true" {
				got = "true"
				if !b {
					got = "false"
				}
			} else {
				got = "false"
				if b {
					got = "true"
				}
			}
		}
		if got != expected {
			return fmt.Errorf("%s.%s: expected %q, got %q (entity ID %d)", entity, field, expected, got, id)
		}
		return nil
	}
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
