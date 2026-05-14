package provider_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Unit tests for schema-level validation. These run without TF_ACC and without a real CiviCRM
// instance because the errors are raised during plan validation, before any API call is made.

func TestACLResource_Validate_InvalidOperation(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_acl" "test" {
  name         = "tf_unit_acl"
  entity_id    = 1
  operation    = "BadOp"
  object_table = "civicrm_group"
}`,
				ExpectError: regexp.MustCompile(`(?i)BadOp.*not a valid|not a valid.*BadOp`),
			},
		},
	})
}

// Note: object_table is a free varchar in CiviCRM with no DB-level enum.
// The valid values depend on the CiviCRM version and installed extensions.
// Validation is intentionally not an enum — CiviCRM itself is the authority.
// The acceptance test TestAccACLAllObjectTables verifies the known values against
// the live CiviCRM instance.

func TestACLResource_Validate_EmptyObjectTable(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_acl" "test" {
  name         = "tf_unit_acl"
  entity_id    = 1
  operation    = "View"
  object_table = ""
}`,
				ExpectError: regexp.MustCompile(`(?i)too short|at least 1`),
			},
		},
	})
}

func TestACLResource_Validate_InvalidEntityTable(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_acl" "test" {
  name         = "tf_unit_acl"
  entity_table = "civicrm_contact"
  entity_id    = 1
  operation    = "View"
  object_table = "civicrm_group"
}`,
				ExpectError: regexp.MustCompile(`(?i)civicrm_contact.*not a valid|not a valid.*civicrm_contact`),
			},
		},
	})
}

func TestACLResource_Validate_EmptyName(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_acl" "test" {
  name         = ""
  entity_id    = 1
  operation    = "View"
  object_table = "civicrm_group"
}`,
				ExpectError: regexp.MustCompile(`(?i)too short|at least 1`),
			},
		},
	})
}

func TestACLResource_Validate_EntityIDZero(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_acl" "test" {
  name         = "tf_unit_acl"
  entity_id    = 0
  operation    = "View"
  object_table = "civicrm_group"
}`,
				ExpectError: regexp.MustCompile(`(?i)at least 1|too small`),
			},
		},
	})
}

// TestACLResource_Validate_AclTableWithoutAclID checks that providing acl_table without acl_id
// is caught by the cross-field ValidateConfig.
func TestACLResource_Validate_AclTableWithoutAclID(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_acl" "test" {
  name         = "tf_unit_acl"
  entity_id    = 1
  operation    = "View"
  object_table = "civicrm_group"
  acl_table    = "civicrm_acl"
}`,
				ExpectError: regexp.MustCompile(`(?i)acl_id.*must be set|Missing acl_id`),
			},
		},
	})
}

// TestACLResource_Validate_AclIDWithoutAclTable checks the inverse cross-field constraint.
func TestACLResource_Validate_AclIDWithoutAclTable(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_acl" "test" {
  name         = "tf_unit_acl"
  entity_id    = 1
  operation    = "View"
  object_table = "civicrm_group"
  acl_id       = 42
}`,
				ExpectError: regexp.MustCompile(`(?i)acl_table.*must be set|Missing acl_table`),
			},
		},
	})
}

// --- ACL Role unit validation tests ---

func TestACLRoleResource_Validate_EmptyName(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_acl_role" "test" {
  name  = ""
  label = "My Role"
}`,
				ExpectError: regexp.MustCompile(`(?i)too short|at least 1`),
			},
		},
	})
}

func TestACLRoleResource_Validate_EmptyLabel(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_acl_role" "test" {
  name  = "my_role"
  label = ""
}`,
				ExpectError: regexp.MustCompile(`(?i)too short|at least 1`),
			},
		},
	})
}

// --- ACL Entity Role unit validation tests ---

func TestACLEntityRoleResource_Validate_InvalidEntityTable(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_acl_entity_role" "test" {
  acl_role_id  = 5
  entity_table = "civicrm_contact"
  entity_id    = 1
}`,
				ExpectError: regexp.MustCompile(`(?i)civicrm_contact.*not a valid|not a valid.*civicrm_contact`),
			},
		},
	})
}

func TestACLEntityRoleResource_Validate_AclRoleIDZero(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_acl_entity_role" "test" {
  acl_role_id = 0
  entity_id   = 1
}`,
				ExpectError: regexp.MustCompile(`(?i)at least 1|too small`),
			},
		},
	})
}

func TestACLEntityRoleResource_Validate_EntityIDZero(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "civicrm_acl_entity_role" "test" {
  acl_role_id = 5
  entity_id   = 0
}`,
				ExpectError: regexp.MustCompile(`(?i)at least 1|too small`),
			},
		},
	})
}
