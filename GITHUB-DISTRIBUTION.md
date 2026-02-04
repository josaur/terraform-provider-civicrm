# Quick Reference: GitHub-Based Provider Distribution

This document provides a quick reference for understanding how this provider is distributed via GitHub (not the official Terraform Registry).

## Key Points

✅ **No Terraform Registry registration required**  
✅ **Automatic installation via `terraform init`**  
✅ **Standard Terraform workflow for users**  
✅ **Distributed through GitHub Releases**

## For Users

### Using the Provider

Simply add to your Terraform configuration:

```hcl
terraform {
  required_providers {
    civicrm = {
      source  = "Caritas-Deutschland-Digitallabor/civicrm"
      version = "~> 0.1"
    }
  }
}
```

Run `terraform init` - Terraform automatically downloads from GitHub releases.

### Provider Source Format

`Caritas-Deutschland-Digitallabor/civicrm` means:
- **Namespace**: `Caritas-Deutschland-Digitallabor` (GitHub organization)
- **Type**: `civicrm` (provider name)
- **Repository**: `https://github.com/Caritas-Deutschland-Digitallabor/civicrm-terraform`

### What Happens During `terraform init`

1. Terraform queries GitHub API for available versions
2. Selects version matching your constraint (e.g., `~> 0.1`)
3. Downloads the appropriate binary for your OS/architecture
4. Verifies SHA256 checksum
5. Verifies GPG signature (if available)
6. Installs to local plugin cache
7. Ready to use!

## For Maintainers

### Publishing a New Version

1. Ensure all changes are committed and tests pass
2. Update CHANGELOG.md
3. Create and push a semantic version tag:
   ```bash
   git tag v0.1.0
   git push origin v0.1.0
   ```
4. GitHub Actions automatically:
   - Builds binaries for all platforms
   - Creates release with artifacts
   - Generates checksums and signatures

### Required Release Artifacts

Each release must include:
- Platform binaries (e.g., `terraform-provider-civicrm_v0.1.0_linux_amd64.zip`)
- Checksum file: `terraform-provider-civicrm_0.1.0_SHA256SUMS`
- Signature file: `terraform-provider-civicrm_0.1.0_SHA256SUMS.sig` (optional)

### Supported Platforms

- Linux: amd64, arm64
- macOS/Darwin: amd64, arm64
- Windows: amd64, arm64

## GitHub vs Official Registry

| Feature | GitHub Releases | Terraform Registry |
|---------|----------------|-------------------|
| Setup | Push tags | Sign up + verify namespace |
| Publishing | Automatic (CI/CD) | Manual or API |
| User experience | Identical | Identical |
| Discoverability | Via GitHub org | Via registry search |
| Documentation | README/GitHub Wiki | Registry docs page |

## Troubleshooting

### Provider not found
- Check: https://github.com/Caritas-Deutschland-Digitallabor/civicrm-terraform/releases
- Ensure at least one release exists
- Verify internet access to GitHub

### Checksum verification failed
```bash
rm -rf ~/.terraform.d/plugins/
rm .terraform.lock.hcl
terraform init
```

### Air-gapped installation
1. Download ZIP from releases
2. Extract to plugin directory:
   ```bash
   ~/.terraform.d/plugins/registry.terraform.io/Caritas-Deutschland-Digitallabor/civicrm/<VERSION>/<OS>_<ARCH>/
   ```
3. Run `terraform init`

## Documentation

- **Full Documentation**: See [README.md](README.md)
- **Usage Guide**: See [USAGE.md](USAGE.md)
- **Publishing Guide**: See [PUBLISHING.md](PUBLISHING.md)
- **GitHub Releases**: https://github.com/Caritas-Deutschland-Digitallabor/civicrm-terraform/releases

## Technical Details

### Implicit Provider Installation

Terraform's implicit provider installation mechanism allows providers to be distributed without registry registration:

1. When `source` doesn't include a hostname, defaults to `registry.terraform.io`
2. If not found in registry, Terraform queries GitHub for organization matching the namespace
3. Discovers releases from `terraform-provider-<TYPE>` repository
4. Downloads and verifies artifacts following Terraform's provider protocol

### Release Naming Conventions

- **Tag**: `v0.1.0` (semantic versioning with `v` prefix)
- **Binary**: `terraform-provider-civicrm_v0.1.0`
- **Archive**: `terraform-provider-civicrm_0.1.0_linux_amd64.zip`
- **Checksums**: `terraform-provider-civicrm_0.1.0_SHA256SUMS`
- **Signature**: `terraform-provider-civicrm_0.1.0_SHA256SUMS.sig`

### GoReleaser Configuration

This provider uses GoReleaser for automated releases:
- See [.goreleaser.yml](.goreleaser.yml) for configuration
- Triggered by GitHub Actions on tag push
- Builds, archives, checksums, and signs automatically
