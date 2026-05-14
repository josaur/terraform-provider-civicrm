package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// TestAccACLFullChain tests the complete CiviCRM permission model:
//   civicrm_acl_role  → defines a role
//   civicrm_acl_entity_role → assigns the role to a group (group members carry the role)
//   civicrm_acl        → grants an operation on an object to that role
//
// All three resources are created and linked together; the test verifies the
// relationships are stored correctly in CiviCRM via the API.
func TestAccACLFullChain(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy: resource.ComposeAggregateTestCheckFunc(
			checkDestroyByID(t, "ACL", "civicrm_acl.chain_rule"),
			checkDestroyByID(t, "ACLEntityRole", "civicrm_acl_entity_role.chain_assign"),
			checkDestroyByID(t, "OptionValue", "civicrm_acl_role.chain_role"),
			checkDestroyByID(t, "Group", "civicrm_group.chain_group"),
		),
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + testAccACLFullChainConfig("View"),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Role exists and has a value
					resource.TestCheckResourceAttrSet("civicrm_acl_role.chain_role", "id"),
					resource.TestCheckResourceAttrSet("civicrm_acl_role.chain_role", "value"),
					resource.TestCheckResourceAttr("civicrm_acl_role.chain_role", "name", "tf_acc_chain_role"),
					resource.TestCheckResourceAttr("civicrm_acl_role.chain_role", "is_active", "true"),

					// Group exists
					resource.TestCheckResourceAttrSet("civicrm_acl_entity_role.chain_assign", "id"),
					resource.TestCheckResourceAttr("civicrm_acl_entity_role.chain_assign", "entity_table", "civicrm_group"),
					resource.TestCheckResourceAttr("civicrm_acl_entity_role.chain_assign", "is_active", "true"),

					// ACL rule is linked to the role and targets the right object
					resource.TestCheckResourceAttrSet("civicrm_acl.chain_rule", "id"),
					resource.TestCheckResourceAttr("civicrm_acl.chain_rule", "operation", "View"),
					resource.TestCheckResourceAttr("civicrm_acl.chain_rule", "object_table", "civicrm_group"),
					resource.TestCheckResourceAttr("civicrm_acl.chain_rule", "is_active", "true"),
					resource.TestCheckResourceAttr("civicrm_acl.chain_rule", "deny", "false"),
					resource.TestCheckResourceAttrSet("civicrm_acl.chain_rule", "priority"),

					// DB verification: role fields stored correctly
					checkEntityAttr(t, "OptionValue", "civicrm_acl_role.chain_role", "id", "name", "tf_acc_chain_role"),
					checkEntityAttr(t, "OptionValue", "civicrm_acl_role.chain_role", "id", "label", "TF Acceptance Chain Role"),

					// DB verification: entity role links the right group
					checkEntityAttr(t, "ACLEntityRole", "civicrm_acl_entity_role.chain_assign", "id", "entity_table", "civicrm_group"),
					checkEntityAttr(t, "ACLEntityRole", "civicrm_acl_entity_role.chain_assign", "id", "is_active", "true"),

					// DB verification: ACL rule operation and object
					checkEntityAttr(t, "ACL", "civicrm_acl.chain_rule", "id", "operation", "View"),
					checkEntityAttr(t, "ACL", "civicrm_acl.chain_rule", "id", "object_table", "civicrm_group"),
					checkEntityAttr(t, "ACL", "civicrm_acl.chain_rule", "id", "is_active", "true"),

					// entity_id in ACL must equal the role's value field
					testCheckACLEntityIDMatchesRoleValue("civicrm_acl.chain_rule", "civicrm_acl_role.chain_role"),
					// acl_role_id in entity_role must equal the role's value field
					testCheckEntityRoleMatchesRoleValue("civicrm_acl_entity_role.chain_assign", "civicrm_acl_role.chain_role"),
					// entity_id in entity_role must equal the group's id
					testCheckEntityRoleEntityIDMatchesGroup("civicrm_acl_entity_role.chain_assign", "civicrm_group.chain_group"),
				),
			},
			// Update operation — verifies that changing the permission works correctly
			{
				Config: providerConfig() + testAccACLFullChainConfig("Edit"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_acl.chain_rule", "operation", "Edit"),
					checkEntityAttr(t, "ACL", "civicrm_acl.chain_rule", "id", "operation", "Edit"),
					resource.TestCheckResourceAttrSet("civicrm_acl.chain_rule", "priority"),
				),
			},
		},
	})
}

// TestAccACLDenyRule verifies that a deny ACL rule (deny=true) is created and stored correctly.
func TestAccACLDenyRule(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "ACL", "civicrm_acl.deny_rule"),
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_acl" "deny_rule" {
  name         = "tf_acc_deny_rule"
  entity_table = "civicrm_acl_role"
  entity_id    = 1
  operation    = "Edit"
  object_table = "civicrm_group"
  deny         = true
  is_active    = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_acl.deny_rule", "deny", "true"),
					resource.TestCheckResourceAttr("civicrm_acl.deny_rule", "operation", "Edit"),
					checkEntityAttr(t, "ACL", "civicrm_acl.deny_rule", "id", "deny", "true"),
					checkEntityAttr(t, "ACL", "civicrm_acl.deny_rule", "id", "operation", "Edit"),
				),
			},
		},
	})
}

// TestAccACLDeactivation verifies that a permission rule can be deactivated via is_active=false.
func TestAccACLDeactivation(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             checkDestroyByID(t, "ACL", "civicrm_acl.deactivate_test"),
		Steps: []resource.TestStep{
			// Create active
			{
				Config: providerConfig() + `
resource "civicrm_acl" "deactivate_test" {
  name         = "tf_acc_deactivate"
  entity_table = "civicrm_acl_role"
  entity_id    = 1
  operation    = "View"
  object_table = "civicrm_group"
  is_active    = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_acl.deactivate_test", "is_active", "true"),
					checkEntityAttr(t, "ACL", "civicrm_acl.deactivate_test", "id", "is_active", "true"),
				),
			},
			// Deactivate
			{
				Config: providerConfig() + `
resource "civicrm_acl" "deactivate_test" {
  name         = "tf_acc_deactivate"
  entity_table = "civicrm_acl_role"
  entity_id    = 1
  operation    = "View"
  object_table = "civicrm_group"
  is_active    = false
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("civicrm_acl.deactivate_test", "is_active", "false"),
					checkEntityAttr(t, "ACL", "civicrm_acl.deactivate_test", "id", "is_active", "false"),
				),
			},
		},
	})
}

// TestAccACLObjectIDLink tests that object_id correctly references an existing group, and
// that the ACL is scoped to that specific group (not all groups).
func TestAccACLObjectIDLink(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy: resource.ComposeAggregateTestCheckFunc(
			checkDestroyByID(t, "ACL", "civicrm_acl.scoped_rule"),
			checkDestroyByID(t, "Group", "civicrm_group.acl_target_group"),
		),
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_group" "acl_target_group" {
  name  = "tf_acc_acl_target_group"
  title = "TF Acceptance ACL Target Group"
}

resource "civicrm_acl" "scoped_rule" {
  name         = "tf_acc_scoped_acl"
  entity_table = "civicrm_acl_role"
  entity_id    = 1
  operation    = "View"
  object_table = "civicrm_group"
  object_id    = civicrm_group.acl_target_group.id
  is_active    = true
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("civicrm_acl.scoped_rule", "object_id"),
					// object_id in ACL must equal the referenced group's id
					testCheckACLObjectIDMatchesGroup("civicrm_acl.scoped_rule", "civicrm_group.acl_target_group"),
					checkEntityAttr(t, "ACL", "civicrm_acl.scoped_rule", "id", "object_table", "civicrm_group"),
					checkEntityAttr(t, "ACL", "civicrm_acl.scoped_rule", "id", "operation", "View"),
				),
			},
		},
	})
}

// TestAccACLAllOperations ensures all valid operation values are accepted by CiviCRM.
func TestAccACLAllOperations(t *testing.T) {
	operations := []string{"Edit", "View", "Create", "Delete", "Search", "All"}

	for _, op := range operations {
		op := op // capture loop variable
		t.Run(op, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck:                 func() { testAccPreCheck(t) },
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				CheckDestroy:             checkDestroyByID(t, "ACL", "civicrm_acl.op_test"),
				Steps: []resource.TestStep{
					{
						Config: providerConfig() + fmt.Sprintf(`
resource "civicrm_acl" "op_test" {
  name         = "tf_acc_op_%s"
  entity_table = "civicrm_acl_role"
  entity_id    = 1
  operation    = %q
  object_table = "civicrm_group"
  is_active    = true
}`, op, op),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("civicrm_acl.op_test", "operation", op),
							checkEntityAttr(t, "ACL", "civicrm_acl.op_test", "id", "operation", op),
						),
					},
				},
			})
		})
	}
}

// TestAccACLAllObjectTables verifies that the object_table values known from the
// CiviCRM admin UI are actually accepted by the live CiviCRM instance.
// This is the acceptance-test complement to the informational CI log step
// "Log valid ACL object_table values from CiviCRM schema".
//
// object_table is a free varchar in civicrm_acl — CiviCRM does not enforce an enum
// at the DB or API level.  This test creates one ACL per known value and confirms
// CiviCRM stores it without error.  If a future CiviCRM version removes a value,
// this test will catch it so the provider documentation can be updated.
func TestAccACLAllObjectTables(t *testing.T) {
	// Values shown in the CiviCRM ACL admin UI under "Type of Data".
	// Do NOT add values here without confirming against a running CiviCRM instance.
	knownObjectTables := []string{
		"civicrm_group",         // static group contacts
		"civicrm_saved_search",  // smart group contacts
		"civicrm_uf_group",      // profiles
		"civicrm_custom_group",  // custom data groups
	}

	for _, tbl := range knownObjectTables {
		tbl := tbl // capture loop variable
		t.Run(tbl, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck:                 func() { testAccPreCheck(t) },
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				CheckDestroy:             checkDestroyByID(t, "ACL", "civicrm_acl.objtbl_test"),
				Steps: []resource.TestStep{
					{
						Config: providerConfig() + fmt.Sprintf(`
resource "civicrm_acl" "objtbl_test" {
  name         = "tf_acc_objtbl_%s"
  entity_table = "civicrm_acl_role"
  entity_id    = 1
  operation    = "View"
  object_table = %q
  is_active    = true
}`, tbl, tbl),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("civicrm_acl.objtbl_test", "object_table", tbl),
							checkEntityAttr(t, "ACL", "civicrm_acl.objtbl_test", "id", "object_table", tbl),
						),
					},
				},
			})
		})
	}
}

// --- helpers ---

func testAccACLFullChainConfig(operation string) string {
	return fmt.Sprintf(`
resource "civicrm_group" "chain_group" {
  name  = "tf_acc_chain_group"
  title = "TF Acceptance Chain Group"
}

resource "civicrm_acl_role" "chain_role" {
  name      = "tf_acc_chain_role"
  label     = "TF Acceptance Chain Role"
  is_active = true
}

resource "civicrm_acl_entity_role" "chain_assign" {
  acl_role_id  = civicrm_acl_role.chain_role.value
  entity_table = "civicrm_group"
  entity_id    = civicrm_group.chain_group.id
  is_active    = true
}

resource "civicrm_acl" "chain_rule" {
  name         = "tf_acc_chain_acl"
  entity_table = "civicrm_acl_role"
  entity_id    = civicrm_acl_role.chain_role.value
  operation    = %q
  object_table = "civicrm_group"
  object_id    = civicrm_group.chain_group.id
  is_active    = true

  depends_on = [civicrm_acl_entity_role.chain_assign]
}
`, operation)
}

// testCheckACLEntityIDMatchesRoleValue verifies that the ACL's entity_id equals the role's value.
func testCheckACLEntityIDMatchesRoleValue(aclAddr, roleAddr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		aclRS, ok := s.RootModule().Resources[aclAddr]
		if !ok {
			return fmt.Errorf("resource %q not in state", aclAddr)
		}
		roleRS, ok := s.RootModule().Resources[roleAddr]
		if !ok {
			return fmt.Errorf("resource %q not in state", roleAddr)
		}
		aclEntityID := aclRS.Primary.Attributes["entity_id"]
		roleValue := roleRS.Primary.Attributes["value"]
		if aclEntityID != roleValue {
			return fmt.Errorf("ACL entity_id (%s) does not match acl_role.value (%s) — the ACL rule is not linked to the role", aclEntityID, roleValue)
		}
		return nil
	}
}

// testCheckEntityRoleMatchesRoleValue verifies that the entity_role's acl_role_id equals the role's value.
func testCheckEntityRoleMatchesRoleValue(entityRoleAddr, roleAddr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		erRS, ok := s.RootModule().Resources[entityRoleAddr]
		if !ok {
			return fmt.Errorf("resource %q not in state", entityRoleAddr)
		}
		roleRS, ok := s.RootModule().Resources[roleAddr]
		if !ok {
			return fmt.Errorf("resource %q not in state", roleAddr)
		}
		erRoleID := erRS.Primary.Attributes["acl_role_id"]
		roleValue := roleRS.Primary.Attributes["value"]
		if erRoleID != roleValue {
			return fmt.Errorf("acl_entity_role.acl_role_id (%s) does not match acl_role.value (%s)", erRoleID, roleValue)
		}
		return nil
	}
}

// testCheckEntityRoleEntityIDMatchesGroup verifies that the entity_role's entity_id equals the group's id.
func testCheckEntityRoleEntityIDMatchesGroup(entityRoleAddr, groupAddr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		erRS, ok := s.RootModule().Resources[entityRoleAddr]
		if !ok {
			return fmt.Errorf("resource %q not in state", entityRoleAddr)
		}
		grpRS, ok := s.RootModule().Resources[groupAddr]
		if !ok {
			return fmt.Errorf("resource %q not in state", groupAddr)
		}
		erEntityID := erRS.Primary.Attributes["entity_id"]
		groupID := grpRS.Primary.ID
		if erEntityID != groupID {
			return fmt.Errorf("acl_entity_role.entity_id (%s) does not match group.id (%s)", erEntityID, groupID)
		}
		return nil
	}
}

// testCheckACLObjectIDMatchesGroup verifies that the ACL's object_id equals the group's id.
func testCheckACLObjectIDMatchesGroup(aclAddr, groupAddr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		aclRS, ok := s.RootModule().Resources[aclAddr]
		if !ok {
			return fmt.Errorf("resource %q not in state", aclAddr)
		}
		grpRS, ok := s.RootModule().Resources[groupAddr]
		if !ok {
			return fmt.Errorf("resource %q not in state", groupAddr)
		}
		objectID := aclRS.Primary.Attributes["object_id"]
		groupID := grpRS.Primary.ID
		if objectID != groupID {
			return fmt.Errorf("acl.object_id (%s) does not match group.id (%s) — cross-resource link is broken", objectID, groupID)
		}
		return nil
	}
}
