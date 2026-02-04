# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Documentation for publishing and using the provider
- PUBLISHING.md with detailed instructions for maintainers
- terraform-registry-manifest.json for Terraform Registry compatibility
- CHANGELOG.md for tracking releases

### Changed
- Updated module path from `github.com/example/terraform-provider-civicrm` to `github.com/Caritas-Deutschland-Digitallabor/civicrm-terraform`
- Updated provider source from `registry.terraform.io/example/civicrm` to `Caritas-Deutschland-Digitallabor/civicrm`
- Improved README with clear instructions for using the provider from GitHub releases
- Updated all examples to use the correct provider source

## [0.1.0] - Initial Release (Planned)

### Added
- CiviCRM Group resource and data source
- CiviCRM ACL Role resource and data source
- CiviCRM ACL resource and data source
- CiviCRM ACL Entity Role resource and data source
- API v4 client for CiviCRM integration
- Comprehensive documentation and examples
- GitHub Actions workflow for automated releases
- GoReleaser configuration for multi-platform builds

[Unreleased]: https://github.com/Caritas-Deutschland-Digitallabor/civicrm-terraform/compare/v0.1.0...HEAD
