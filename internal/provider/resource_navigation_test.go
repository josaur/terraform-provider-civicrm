package provider_test

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// checkNavigationPermissionIsList asserts that Navigation.permission, as
// returned by Navigation.get, is a list with exactly the given elements —
// not a single string containing a comma. permission is stored as
// SERIALIZE_COMMA internally; a round-trip through this provider's own Read
// path cannot distinguish a correctly-stored list from a single
// comma-joined string, since both decode to the same list shape here. This
// check bypasses the provider and re-reads the raw API response directly.
func checkNavigationPermissionIsList(t *testing.T, resourceAddr string, want []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		t.Helper()
		rs, ok := s.RootModule().Resources[resourceAddr]
		if !ok {
			t.Fatalf("resource %q not found in state", resourceAddr)
		}
		id, err := strconv.ParseInt(rs.Primary.Attributes["id"], 10, 64)
		if err != nil {
			t.Fatalf("resource %q has no valid id: %s", resourceAddr, err)
		}

		client := newTestClient(t)
		result, err := client.GetByID("Navigation", id, []string{"permission"})
		if err != nil {
			t.Fatalf("GetByID(Navigation, %d): %s", id, err)
		}

		raw, ok := result["permission"]
		if !ok || raw == nil {
			t.Fatalf("Navigation %d has no permission", id)
		}
		got, ok := raw.([]any)
		if !ok {
			t.Fatalf("Navigation %d permission is %T, not a list: %#v", id, raw, raw)
		}
		if len(got) != len(want) {
			t.Fatalf("Navigation %d permission has %d entries, want %d: %#v", id, len(got), len(want), got)
		}
		for i, w := range want {
			gs, ok := got[i].(string)
			if !ok || gs != w {
				t.Fatalf("Navigation %d permission[%d] = %#v, want %q", id, i, got[i], w)
			}
		}
		return nil
	}
}

func TestAccNavigationResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "Navigation", "civicrm_navigation"),
		Steps: []resource.TestStep{
			// Create: single permission, no operator needed.
			{
				Config: providerConfig() + `
resource "civicrm_navigation" "test" {
  label      = "TF Acceptance Test Nav"
  name       = "tf_acc_navigation"
  url        = "civicrm/tf-acc-test"
  permission = ["access CiviCRM"]
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_navigation.test", "id"),
					resource.TestCheckResourceAttr("civicrm_navigation.test", "label", "TF Acceptance Test Nav"),
					resource.TestCheckResourceAttr("civicrm_navigation.test", "name", "tf_acc_navigation"),
					resource.TestCheckResourceAttr("civicrm_navigation.test", "is_active", "true"),
					resource.TestCheckResourceAttr("civicrm_navigation.test", "has_separator", "0"),
					resource.TestCheckResourceAttrSet("civicrm_navigation.test", "weight"),
					resource.TestCheckResourceAttr("civicrm_navigation.test", "permission.#", "1"),
					resource.TestCheckResourceAttr("civicrm_navigation.test", "permission.0", "access CiviCRM"),
					checkEntityAttr(t, "Navigation", "civicrm_navigation.test", "id", "name", "tf_acc_navigation"),
					checkNavigationPermissionIsList(t, "civicrm_navigation.test", []string{"access CiviCRM"}),
				),
			},
			// Import: Terraform state must survive a round-trip through the API.
			{
				ResourceName:      "civicrm_navigation.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update: two permissions with an explicit operator — the
			// meaningful test per the issue, since a two-element list only
			// survives a broken comma-join if it's stored/read correctly.
			{
				Config: providerConfig() + `
resource "civicrm_navigation" "test" {
  label                = "TF Acceptance Test Nav"
  name                 = "tf_acc_navigation"
  url                  = "civicrm/tf-acc-test"
  permission            = ["access CiviCRM", "administer CiviCRM"]
  permission_operator   = "AND"
  has_separator         = 1
  weight                = 42
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_navigation.test", "permission.#", "2"),
					resource.TestCheckResourceAttr("civicrm_navigation.test", "permission.0", "access CiviCRM"),
					resource.TestCheckResourceAttr("civicrm_navigation.test", "permission.1", "administer CiviCRM"),
					resource.TestCheckResourceAttr("civicrm_navigation.test", "permission_operator", "AND"),
					resource.TestCheckResourceAttr("civicrm_navigation.test", "has_separator", "1"),
					resource.TestCheckResourceAttr("civicrm_navigation.test", "weight", "42"),
					checkNavigationPermissionIsList(t, "civicrm_navigation.test", []string{"access CiviCRM", "administer CiviCRM"}),
				),
			},
		},
	})
}

func TestAccNavigationResource_nested(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "Navigation", "civicrm_navigation"),
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_navigation" "parent" {
  label = "TF Acceptance Test Parent"
  name  = "tf_acc_navigation_parent"
}

resource "civicrm_navigation" "child" {
  label     = "TF Acceptance Test Child"
  name      = "tf_acc_navigation_child"
  parent_id = civicrm_navigation.parent.id
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("civicrm_navigation.child", "parent_id", "civicrm_navigation.parent", "id"),
				),
			},
		},
	})
}

func TestAccNavigationDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "Navigation", "civicrm_navigation"),
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_navigation" "test" {
  label = "TF Acceptance Test Nav DS"
  name  = "tf_acc_navigation_ds"
}

data "civicrm_navigation" "lookup" {
  name = civicrm_navigation.test.name
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.civicrm_navigation.lookup", "id", "civicrm_navigation.test", "id"),
					resource.TestCheckResourceAttr("data.civicrm_navigation.lookup", "label", "TF Acceptance Test Nav DS"),
				),
			},
		},
	})
}
