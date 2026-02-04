# Publishing Guide

This guide explains how to publish new releases of the CiviCRM Terraform Provider.

## Overview

This provider uses [GoReleaser](https://goreleaser.com/) to automate the release process. Releases are triggered automatically when a new version tag is pushed to the repository.

## Prerequisites

Before publishing a release, ensure you have:

1. **GPG Key** (optional but recommended for signing releases):
   - Generate a GPG key if you don't have one: `gpg --full-generate-key`
   - Add the GPG private key to GitHub Secrets as `GPG_PRIVATE_KEY`
   - Add the passphrase to GitHub Secrets as `GPG_PASSPHRASE`
   - Export your public key for verification: `gpg --armor --export YOUR_EMAIL`

2. **GitHub Token**: The workflow uses `GITHUB_TOKEN` which is automatically provided by GitHub Actions.

3. **Write Permissions**: Ensure you have write access to the repository.

## Release Process

### 1. Prepare the Release

Before creating a release, ensure:

- All changes are committed and merged to the main branch
- Tests pass successfully (`make test` or similar)
- The code builds without errors (`make build` or `go build`)
- Update CHANGELOG.md with the new version and changes (if it exists)
- Update version references in documentation if needed

### 2. Create and Push a Version Tag

The release workflow is triggered by pushing a tag that starts with `v`:

```bash
# Create a new tag (use semantic versioning)
git tag v0.1.0

# Push the tag to GitHub
git push origin v0.1.0
```

Version Format:
- Use [Semantic Versioning](https://semver.org/): `vMAJOR.MINOR.PATCH`
- Example: `v0.1.0`, `v1.0.0`, `v1.2.3`
- For pre-releases: `v0.1.0-rc1`, `v0.1.0-alpha`, `v0.1.0-beta.1`

### 3. Monitor the Release

Once the tag is pushed:

1. Go to the **Actions** tab in GitHub
2. Watch the "Release" workflow execution
3. The workflow will:
   - Build binaries for multiple platforms (Linux, macOS, Windows)
   - Create ZIP archives for each platform
   - Generate SHA256 checksums
   - Sign the checksums with GPG (if configured)
   - Create a GitHub Release with all artifacts
   - Generate a changelog from commit messages

### 4. Verify the Release

After the workflow completes:

1. Go to the **Releases** section in GitHub
2. Verify the new release is published
3. Check that all platform binaries are attached
4. Verify checksums and signatures (if applicable)

### 5. Update Documentation

After a successful release:

1. Update the main README.md if needed
2. Announce the release (if applicable)
3. Update any external documentation referencing version numbers

## Using the Published Provider

Once published, users can reference the provider in their Terraform configurations:

```hcl
terraform {
  required_providers {
    civicrm = {
      source  = "Caritas-Deutschland-Digitallabor/civicrm"
      version = "~> 0.1.0"
    }
  }
}

provider "civicrm" {
  url     = "https://your-civicrm-instance.org"
  api_key = "your-api-key"
}
```

When users run `terraform init`, Terraform will:
1. Download the provider from the GitHub releases
2. Verify checksums (and signatures if available)
3. Install the provider in their local plugin cache

## Platform Support

The provider is built for the following platforms:

- **Linux**: amd64, arm64
- **macOS/Darwin**: amd64, arm64
- **Windows**: amd64, arm64

## GoReleaser Configuration

The release process is configured in `.goreleaser.yml`. Key features:

- **Binary Naming**: `terraform-provider-civicrm_vX.Y.Z`
- **Archive Format**: ZIP files
- **Checksums**: SHA256SUMS file
- **Signing**: GPG signature of checksums (if configured)
- **Changelog**: Auto-generated from git commits

## Troubleshooting

### Release Workflow Fails

If the release workflow fails:

1. Check the workflow logs in the Actions tab
2. Common issues:
   - Missing or incorrect GPG secrets (if using signing)
   - Build errors in the code
   - Network issues downloading dependencies

### GPG Signing Issues

If you encounter GPG signing errors:

1. Verify GPG secrets are correctly set in repository settings
2. Test GPG signing locally: `gpg --sign --local-user YOUR_KEY test.txt`
3. You can temporarily disable signing by removing the `signs` section from `.goreleaser.yml`

### Creating a Release Without GPG Signing

If you don't have GPG configured:

1. The workflow will check for `GPG_PRIVATE_KEY` and skip signing if not present
2. Releases will still be created, just without cryptographic signatures
3. Add the secrets later to enable signing for future releases

## Rollback

If you need to remove a release:

1. Go to the Releases section in GitHub
2. Click on the release you want to remove
3. Click "Delete this release"
4. Optionally delete the tag: `git push --delete origin vX.Y.Z`

Note: Deleting a release doesn't prevent users who already downloaded it from using it.

## Best Practices

1. **Test Before Release**: Always test thoroughly on the main branch before tagging
2. **Semantic Versioning**: Follow semver strictly to help users understand changes
3. **Changelog**: Keep a CHANGELOG.md file documenting all changes
4. **Communication**: Announce significant releases to users
5. **Security**: Use GPG signing for production releases
6. **Backward Compatibility**: Avoid breaking changes in minor versions
