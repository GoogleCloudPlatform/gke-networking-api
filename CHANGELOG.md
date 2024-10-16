# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

### Added

- Status conditions to be used for validation updates for GNPs with `NetworkAttachment`.
- Innterface-status annotation was added to pod.
- Network annotation to hold egress NAT gateway IP.
- The `NodeTopology` CRD definition.
- The `description` field in the GCPFirewall CRD.

### Removed

- Unused v1beta GCPFirewall CRD.

## [0.1.1] - 2024-07-02

### Changed

- The `status` field in the GCPFirewall CRD is now a subresource.

## [0.1.0] - 2024-06-18

### Changed

- CRD definitions were migrated from the `kubernetes/ingress-gce`
  repository to simplify cross-project maintanance.

- CRD generation scripts were updated to use the `kube_generator`
  script istead of the deprecated `generate_group` script.

[unreleased]: https://github.com/GoogleCloudPlatform/gke-networking-api/compare/v0.1.1...HEAD
[0.1.1]: https://github.com/GoogleCloudPlatform/gke-networking-api/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/GoogleCloudPlatform/gke-networking-api/releases/v0.1.0
