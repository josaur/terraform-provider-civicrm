# Terraform Provider for CiviCRM

[![Build](https://github.com/example/terraform-provider-civicrm/actions/workflows/build.yml/badge.svg)](https://github.com/example/terraform-provider-civicrm/actions/workflows/build.yml)
[![Release](https://github.com/example/terraform-provider-civicrm/actions/workflows/release.yml/badge.svg)](https://github.com/example/terraform-provider-civicrm/actions/workflows/release.yml)

A Terraform provider for managing CiviCRM access control resources via API v4.

## Features

This provider supports managing the following CiviCRM resources:

- **Groups** (`civicrm_group`) - CiviCRM groups that can be assigned ACL roles
- **ACL Roles** (`civicrm_acl_role`) - Permission roles that define access levels
- **ACLs** (`civicrm_acl`) - Access control rules defining what operations a role can perform
- **ACL Entity Roles** (`civicrm_acl_entity_role`) - Assigns ACL roles to groups (role bindings)

Each resource also has a corresponding data source for read-only lookups.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.21 (for building from source)
- CiviCRM >= 5.47 (required for full REST API v4 support)

## Building the Provider

1. Clone the repository:
```bash
git clone https://github.com/example/terraform-provider-civicrm.git
cd terraform-provider-civicrm
```

2. Build the provider:
```bash
go mod tidy
go build -o terraform-provider-civicrm
```

3. Install for local development:
```bash
# Linux/Mac
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/example/civicrm/0.1.0/$(go env GOOS)_$(go env GOARCH)
cp terraform-provider-civicrm ~/.terraform.d/plugins/registry.terraform.io/example/civicrm/0.1.0/$(go env GOOS)_$(go env GOARCH)/

# Windows (PowerShell)
$installDir = "$env:APPDATA\terraform.d\plugins\registry.terraform.io\example\civicrm\0.1.0\windows_amd64"
New-Item -ItemType Directory -Force -Path $installDir
Copy-Item terraform-provider-civicrm.exe $installDir
```

## Configuration

### Provider Configuration

```hcl
provider "civicrm" {
  url     = "https://your-civicrm-instance.org"  # Required
  api_key = "your-api-key"                        # Required

  # Optional
  insecure = false  # Skip TLS verification (for development only)
}
```

### Environment Variables

The provider supports the following environment variables:

| Variable | Description |
|----------|-------------|
| `CIVICRM_URL` | The base URL of your CiviCRM instance |
| `CIVICRM_API_KEY` | Your CiviCRM API key |

### CiviCRM Setup

1. Ensure your CiviCRM instance is version 5.47 or later.

2. Enable the AuthX extension if not already enabled.

3. Generate an API key for a user with appropriate permissions:
   - Navigate to **Administer > Users and Permissions > API Key**
   - Or use cv: `cv api4 Contact.create '{"id": YOUR_CONTACT_ID, "api_key": "your-generated-key"}'`

4. Ensure the user has the following permissions:
   - "access CiviCRM"
   - "authenticate with api key"
   - Appropriate ACL permissions for managing groups and ACLs

## Usage Examples

### Create Groups and ACL Roles

```hcl
# Create a group for team members
resource "civicrm_group" "team_leaders" {
  name        = "team_leaders"
  title       = "Team Leaders"
  description = "Staff members with team lead responsibilities"
  is_active   = true
}

# Create a group for volunteers
resource "civicrm_group" "volunteers" {
  name        = "volunteers"
  title       = "Volunteers"
  description = "Active volunteers"
  is_active   = true
}

# Create an ACL role
resource "civicrm_acl_role" "volunteer_manager" {
  name        = "volunteer_manager"
  label       = "Volunteer Manager"
  description = "Can view and edit volunteers"
  is_active   = true
}
```

### Create ACL Rules

```hcl
# Grant Edit permission on volunteers group to volunteer_manager role
resource "civicrm_acl" "vm_edit_volunteers" {
  name         = "vm_edit_volunteers"
  entity_id    = civicrm_acl_role.volunteer_manager.id
  operation    = "Edit"
  object_table = "civicrm_group"
  object_id    = civicrm_group.volunteers.id
  is_active    = true
}
```

### Assign Roles to Groups

```hcl
# Give team leaders the volunteer manager role
resource "civicrm_acl_entity_role" "team_leaders_vm" {
  acl_role_id  = civicrm_acl_role.volunteer_manager.id
  entity_table = "civicrm_group"
  entity_id    = civicrm_group.team_leaders.id
  is_active    = true
}
```

### Look Up Existing Resources

```hcl
# Find an existing group by name
data "civicrm_group" "administrators" {
  name = "Administrators"
}

# Find an existing ACL role by name
data "civicrm_acl_role" "admin_role" {
  name = "Administrator"
}
```

## Resource Reference

### civicrm_group

Manages a CiviCRM Group.

#### Arguments

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string | Yes | The machine name of the group (must be unique) |
| `title` | string | Yes | The display title of the group |
| `description` | string | No | A description of the group |
| `is_active` | bool | No | Whether the group is active (default: true) |
| `visibility` | string | No | Visibility setting (default: "User and User Admin Only") |
| `group_type` | list(string) | No | The types of the group |

#### Attributes

| Name | Type | Description |
|------|------|-------------|
| `id` | int | The unique identifier of the group |

### civicrm_acl_role

Manages a CiviCRM ACL Role.

#### Arguments

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string | Yes | The machine name of the ACL role |
| `label` | string | Yes | The display label of the ACL role |
| `description` | string | No | A description of the ACL role |
| `is_active` | bool | No | Whether the ACL role is active (default: true) |
| `weight` | int | No | The sort weight of the ACL role |

#### Attributes

| Name | Type | Description |
|------|------|-------------|
| `id` | int | The unique identifier of the ACL role |
| `value` | string | The internal value (auto-generated) |

### civicrm_acl

Manages a CiviCRM ACL rule.

#### Arguments

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string | Yes | The name of the ACL rule |
| `entity_id` | int | Yes | The ID of the ACL role this rule belongs to |
| `operation` | string | Yes | The operation: "Edit", "View", "Create", "Delete", "Search", "All" |
| `object_table` | string | Yes | The type of object being permissioned |
| `entity_table` | string | No | The entity table (default: "civicrm_acl_role") |
| `object_id` | int | No | The specific object ID (null = all) |
| `is_active` | bool | No | Whether the ACL rule is active (default: true) |
| `deny` | bool | No | Deny instead of allow (default: false) |
| `priority` | int | No | Rule priority |

#### Attributes

| Name | Type | Description |
|------|------|-------------|
| `id` | int | The unique identifier of the ACL |

### civicrm_acl_entity_role

Manages a CiviCRM ACL Entity Role assignment.

#### Arguments

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `acl_role_id` | int | Yes | The ID of the ACL role to assign |
| `entity_id` | int | Yes | The ID of the group to assign the role to |
| `entity_table` | string | No | The entity table (default: "civicrm_group") |
| `is_active` | bool | No | Whether the assignment is active (default: true) |

#### Attributes

| Name | Type | Description |
|------|------|-------------|
| `id` | int | The unique identifier of the assignment |

## Importing Existing Resources

All resources support importing by ID:

```bash
# Import a group
terraform import civicrm_group.example 123

# Import an ACL role
terraform import civicrm_acl_role.example 456

# Import an ACL
terraform import civicrm_acl.example 789

# Import an ACL entity role
terraform import civicrm_acl_entity_role.example 101
```

## Development

### Building

```bash
make build
```

### Testing

```bash
make test
```

### Local Installation

```bash
make install
```

### Code Formatting

```bash
make fmt
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
