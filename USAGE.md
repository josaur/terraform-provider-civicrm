# Using This Provider From Other Projects

This document explains how to use the CiviCRM Terraform Provider in your Terraform projects.

## Installation

The provider is automatically downloaded from **GitHub releases** (not the official Terraform Registry) when you initialize Terraform. No manual installation or registry registration is required.

### How It Works

This provider uses Terraform's **implicit provider installation** mechanism:

- When you specify `source = "Caritas-Deutschland-Digitallabor/civicrm"`, Terraform automatically discovers and downloads the provider from GitHub
- No registry.terraform.io registration needed
- Binaries are fetched directly from: https://github.com/Caritas-Deutschland-Digitallabor/civicrm-terraform/releases
- Terraform verifies checksums and signatures automatically

This is a standard distribution method for Terraform providers and works exactly like providers from the official registry from the user's perspective.

## Basic Setup

### 1. Create a Terraform Configuration

Create a new directory for your Terraform configuration and add a `main.tf` file:

```hcl
terraform {
  required_version = ">= 1.0"
  
  required_providers {
    civicrm = {
      source  = "Caritas-Deutschland-Digitallabor/civicrm"
      version = "~> 0.1"
    }
  }
}

provider "civicrm" {
  url     = var.civicrm_url
  api_key = var.civicrm_api_key
}
```

### 2. Configure Variables (Optional but Recommended)

Create a `variables.tf` file:

```hcl
variable "civicrm_url" {
  description = "The URL of your CiviCRM instance"
  type        = string
}

variable "civicrm_api_key" {
  description = "The API key for CiviCRM authentication"
  type        = string
  sensitive   = true
}
```

### 3. Initialize Terraform

Run the following command to download the provider:

```bash
terraform init
```

Terraform will automatically:
- Download the correct version of the provider for your platform
- Verify checksums
- Cache the provider locally

## Providing Credentials

### Option 1: Environment Variables (Recommended for CI/CD)

Set environment variables before running Terraform:

```bash
export CIVICRM_URL="https://your-civicrm-instance.org"
export CIVICRM_API_KEY="your-api-key"

terraform plan
terraform apply
```

Or use Terraform's variable environment variables:

```bash
export TF_VAR_civicrm_url="https://your-civicrm-instance.org"
export TF_VAR_civicrm_api_key="your-api-key"

terraform plan
terraform apply
```

### Option 2: Variables File (for local development)

Create a `terraform.tfvars` file (add this to `.gitignore`!):

```hcl
civicrm_url     = "https://your-civicrm-instance.org"
civicrm_api_key = "your-api-key"
```

Then run:

```bash
terraform plan
terraform apply
```

### Option 3: Command Line Arguments

Pass variables directly on the command line:

```bash
terraform apply \
  -var="civicrm_url=https://your-civicrm-instance.org" \
  -var="civicrm_api_key=your-api-key"
```

## Example: Complete Project Structure

```
my-civicrm-infrastructure/
├── .gitignore           # Include terraform.tfvars!
├── main.tf              # Provider and resource definitions
├── variables.tf         # Variable declarations
├── terraform.tfvars     # Variable values (DO NOT COMMIT)
├── outputs.tf           # Output definitions
└── README.md            # Project documentation
```

### Example `.gitignore`

```
# Local .terraform directories
**/.terraform/*

# .tfstate files
*.tfstate
*.tfstate.*

# Crash log files
crash.log
crash.*.log

# Variable files that may contain sensitive data
terraform.tfvars
*.auto.tfvars

# Override files
override.tf
override.tf.json
*_override.tf
*_override.tf.json

# CLI configuration files
.terraformrc
terraform.rc
```

### Example `main.tf`

```hcl
terraform {
  required_version = ">= 1.0"
  
  required_providers {
    civicrm = {
      source  = "Caritas-Deutschland-Digitallabor/civicrm"
      version = "~> 0.1"
    }
  }
}

provider "civicrm" {
  url     = var.civicrm_url
  api_key = var.civicrm_api_key
}

# Create a group for volunteers
resource "civicrm_group" "volunteers" {
  name        = "volunteers"
  title       = "Volunteers"
  description = "Active volunteers in the organization"
  is_active   = true
}

# Create a group for team leaders
resource "civicrm_group" "team_leaders" {
  name        = "team_leaders"
  title       = "Team Leaders"
  description = "Staff members with team lead responsibilities"
  is_active   = true
}

# Create an ACL role
resource "civicrm_acl_role" "volunteer_manager" {
  name        = "volunteer_manager"
  label       = "Volunteer Manager"
  description = "Can view and edit volunteers"
  is_active   = true
}

# Create an ACL rule
resource "civicrm_acl" "vm_edit_volunteers" {
  name         = "vm_edit_volunteers"
  entity_id    = civicrm_acl_role.volunteer_manager.id
  operation    = "Edit"
  object_table = "civicrm_group"
  object_id    = civicrm_group.volunteers.id
  is_active    = true
}

# Assign role to group
resource "civicrm_acl_entity_role" "team_leaders_vm" {
  acl_role_id  = civicrm_acl_role.volunteer_manager.id
  entity_table = "civicrm_group"
  entity_id    = civicrm_group.team_leaders.id
  is_active    = true
}
```

### Example `outputs.tf`

```hcl
output "volunteers_group_id" {
  description = "ID of the volunteers group"
  value       = civicrm_group.volunteers.id
}

output "team_leaders_group_id" {
  description = "ID of the team leaders group"
  value       = civicrm_group.team_leaders.id
}
```

## CI/CD Integration

### GitHub Actions Example

```yaml
name: Terraform

on:
  push:
    branches: [main]
  pull_request:

jobs:
  terraform:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Setup Terraform
        uses: hashicorp/setup-terraform@v3
        with:
          terraform_version: 1.5.0
      
      - name: Terraform Init
        run: terraform init
        
      - name: Terraform Plan
        env:
          CIVICRM_URL: ${{ secrets.CIVICRM_URL }}
          CIVICRM_API_KEY: ${{ secrets.CIVICRM_API_KEY }}
        run: terraform plan
        
      - name: Terraform Apply
        if: github.ref == 'refs/heads/main' && github.event_name == 'push'
        env:
          CIVICRM_URL: ${{ secrets.CIVICRM_URL }}
          CIVICRM_API_KEY: ${{ secrets.CIVICRM_API_KEY }}
        run: terraform apply -auto-approve
```

## Version Pinning

For production use, pin to a specific version:

```hcl
terraform {
  required_providers {
    civicrm = {
      source  = "Caritas-Deutschland-Digitallabor/civicrm"
      version = "0.1.0"  # Exact version
    }
  }
}
```

Or use version constraints:

```hcl
version = "~> 0.1.0"  # Allow patch updates (0.1.x)
version = ">= 0.1.0, < 0.2.0"  # Range
version = "~> 0.1"    # Allow minor updates (0.x.y)
```

## Troubleshooting

### Provider Installation from GitHub

This provider is distributed via GitHub releases, not the official Terraform Registry. Here are common issues and solutions:

#### Provider Not Found Error

```
Error: Failed to query available provider packages
Could not retrieve the list of available versions for provider Caritas-Deutschland-Digitallabor/civicrm
```

**Causes and Solutions**:

1. **No internet access to GitHub**
   - Ensure you can access https://github.com
   - Check firewall/proxy settings
   - For air-gapped environments, see "Air-Gapped Installation" below

2. **No releases published yet**
   - Check [GitHub releases](https://github.com/Caritas-Deutschland-Digitallabor/civicrm-terraform/releases)
   - Ensure at least one release is published and not in draft state

3. **Version constraint too restrictive**
   - Check available versions in releases
   - Try loosening the version constraint: `version = "~> 0.1"` instead of exact version

4. **Stale cache**
   - Clear Terraform's cache: `rm -rf .terraform .terraform.lock.hcl`
   - Run `terraform init` again

#### Checksum Verification Failed

```
Error: Failed to install provider
checksum mismatch
```

**Solutions**:
1. Clear the plugin cache: `rm -rf ~/.terraform.d/plugins/`
2. Remove lock file: `rm .terraform.lock.hcl`
3. Re-run `terraform init`

This can happen if a release was updated after you first downloaded it.

#### Air-Gapped Installation

For environments without internet access:

1. **Download on a machine with internet**:
   ```bash
   # Go to: https://github.com/Caritas-Deutschland-Digitallabor/civicrm-terraform/releases
   # Download the appropriate file for your platform, e.g.:
   # - terraform-provider-civicrm_0.1.0_linux_amd64.zip
   # - terraform-provider-civicrm_0.1.0_SHA256SUMS
   ```

2. **Transfer files to air-gapped machine**

3. **Install manually**:
   ```bash
   # Extract the ZIP
   unzip terraform-provider-civicrm_0.1.0_linux_amd64.zip
   
   # Create plugin directory
   mkdir -p ~/.terraform.d/plugins/registry.terraform.io/Caritas-Deutschland-Digitallabor/civicrm/0.1.0/linux_amd64/
   
   # Move binary (note: binary name includes version)
   mv terraform-provider-civicrm_v0.1.0 ~/.terraform.d/plugins/registry.terraform.io/Caritas-Deutschland-Digitallabor/civicrm/0.1.0/linux_amd64/
   
   # Make executable
   chmod +x ~/.terraform.d/plugins/registry.terraform.io/Caritas-Deutschland-Digitallabor/civicrm/0.1.0/linux_amd64/terraform-provider-civicrm_v0.1.0
   ```

4. **Run terraform init** - it will detect and use the local binary

For Windows (PowerShell):
```powershell
# Extract and move
Expand-Archive terraform-provider-civicrm_0.1.0_windows_amd64.zip
$pluginDir = "$env:APPDATA\terraform.d\plugins\registry.terraform.io\Caritas-Deutschland-Digitallabor\civicrm\0.1.0\windows_amd64"
New-Item -ItemType Directory -Force -Path $pluginDir
Move-Item terraform-provider-civicrm_v0.1.0.exe $pluginDir
```

### Authentication Issues

If you get authentication errors:

1. Verify your CiviCRM URL is correct (no trailing slash)
2. Verify your API key is valid
3. Ensure your CiviCRM user has the required permissions
4. Check that the AuthX extension is enabled in CiviCRM

### Platform Not Supported

If you get an error about your platform not being supported:

1. Check the [releases](https://github.com/Caritas-Deutschland-Digitallabor/civicrm-terraform/releases) for available platforms
2. You may need to build the provider from source (see main README.md)

## Getting Help

- **Documentation**: See the main [README.md](README.md) for complete resource documentation
- **Issues**: Report bugs or request features on [GitHub Issues](https://github.com/Caritas-Deutschland-Digitallabor/civicrm-terraform/issues)
- **Examples**: Check the [examples/](examples/) directory for more usage examples

## Security Best Practices

1. **Never commit credentials**: Always use `.gitignore` for `terraform.tfvars`
2. **Use environment variables**: Especially in CI/CD environments
3. **Rotate API keys**: Regularly rotate your CiviCRM API keys
4. **Use separate keys**: Use different API keys for development and production
5. **Limit permissions**: Ensure the API key has only the permissions it needs
