terraform {
  required_providers {
    civicrm = {
      source  = "registry.terraform.io/example/civicrm"
      version = "~> 0.1"
    }
  }
}

# Configure the CiviCRM provider
# Set CIVICRM_URL and CIVICRM_API_KEY environment variables
# or provide them directly in the provider block
provider "civicrm" {
  # url     = "https://your-civicrm-instance.org"
  # api_key = "your-api-key"
}

# Create a group for volunteers
resource "civicrm_group" "volunteers" {
  name        = "volunteers"
  title       = "Volunteers"
  description = "Active volunteers in the organization"
  is_active   = true
  visibility  = "User and User Admin Only"
}

# Create a group for team leaders
resource "civicrm_group" "team_leaders" {
  name        = "team_leaders"
  title       = "Team Leaders"
  description = "Staff members with team lead responsibilities"
  is_active   = true
  visibility  = "User and User Admin Only"
}

# Create an ACL role for volunteer managers
resource "civicrm_acl_role" "volunteer_manager" {
  name        = "volunteer_manager"
  label       = "Volunteer Manager"
  description = "Can view and edit volunteers"
  is_active   = true
}

# Create an ACL role for viewers
resource "civicrm_acl_role" "volunteer_viewer" {
  name        = "volunteer_viewer"
  label       = "Volunteer Viewer"
  description = "Can only view volunteers"
  is_active   = true
}

# ACL rule: Volunteer managers can edit the volunteers group
resource "civicrm_acl" "vm_edit_volunteers" {
  name         = "vm_edit_volunteers"
  entity_id    = civicrm_acl_role.volunteer_manager.id
  operation    = "Edit"
  object_table = "civicrm_group"
  object_id    = civicrm_group.volunteers.id
  is_active    = true
}

# ACL rule: Volunteer viewers can view the volunteers group
resource "civicrm_acl" "vv_view_volunteers" {
  name         = "vv_view_volunteers"
  entity_id    = civicrm_acl_role.volunteer_viewer.id
  operation    = "View"
  object_table = "civicrm_group"
  object_id    = civicrm_group.volunteers.id
  is_active    = true
}

# Assign the volunteer manager role to team leaders
resource "civicrm_acl_entity_role" "team_leaders_vm" {
  acl_role_id  = civicrm_acl_role.volunteer_manager.id
  entity_table = "civicrm_group"
  entity_id    = civicrm_group.team_leaders.id
  is_active    = true
}

# Data source: Look up an existing group
data "civicrm_group" "administrators" {
  name = "Administrators"
}

# Output examples
output "volunteers_group_id" {
  description = "ID of the volunteers group"
  value       = civicrm_group.volunteers.id
}

output "team_leaders_group_id" {
  description = "ID of the team leaders group"
  value       = civicrm_group.team_leaders.id
}

output "volunteer_manager_role_id" {
  description = "ID of the volunteer manager ACL role"
  value       = civicrm_acl_role.volunteer_manager.id
}

output "administrators_group_id" {
  description = "ID of the existing administrators group (from data source)"
  value       = data.civicrm_group.administrators.id
}
