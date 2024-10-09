# RELEASE PROCESS

This document outlines the process for creating releases for
`gke-networking-api`, including versioning conventions and a changelog of notable
changes in each release.

## Module Versioning

When merging changes, please request a tag update from an owner in OWNERS
for the merged changes when needed.

This project uses Go module versioning
(https://go.dev/doc/modules/version-numbers), which uses semantic versioning
(vMAJOR.MINOR.PATCH) to indicate the nature of changes.

However, as the project has the v0 major version, so it makes no stability or
backward compatibility guarantees.

* MINOR: Adding functionality.
* PATCH: Bug fixes.

## Releases

Please visit the [changelog](CHANGELOG.md) for a detailed changelog. Every
release should update the CHANGELOG.md

**Release tags are created manually from the CLI using the following commands:**

Before creating a new release tag, please ensure you have updated the
CHANGELOG.md file.

```sh
git tag <tag>
git push origin <tag>
```
