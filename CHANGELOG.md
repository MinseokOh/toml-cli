# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2026-06-17

### Changed
- Migrated from `pelletier/go-toml` v1 to `go-toml/v2`. The internal wrapper now
  unmarshals documents into a `map[string]any` tree and performs `get`/`set`/
  `merge` via dot-notation key traversal instead of the removed v1 `Tree` API.
- `set` now infers value types by delegating to go-toml (parsing `v = <input>`),
  covering booleans, integers (incl. hex/octal/binary and underscores), floats,
  date/time values, arrays and inline tables, with a string fallback.
- Replaced `interface{}` with the `any` alias throughout the codebase.
- Updated `CLAUDE.md` to match the current package layout and behavior.

### Fixed
- `set` no longer parses integer inputs `1` and `0` as the booleans `true` and
  `false`; they are now stored as integers.
- Removed the Homebrew formula from the GoReleaser configuration.

### Note
- Output formatting follows go-toml v2 conventions: keys are emitted in
  alphabetical order and strings use single quotes; comments are not preserved
  on round-trip.

## [0.1.3] - 2025-07-10

### Added
- `merge` command for combining two TOML files with recursive object merging,
  along with a comprehensive sample structure.

### Fixed
- Updated test file paths and added a `.gitignore` for test outputs.

## [0.1.2] - 2025-07-10

### Added
- GoReleaser configuration and a GitHub Actions release workflow.

### Changed
- Modernized the codebase and added Claude Code integration.
- Updated GitHub Actions to use Go 1.23 and the latest action versions.

### Fixed
- Added the missing `stretchr/objx` dependency.
- Reduced GoReleaser deprecation warnings.

## [0.1.1] - 2024-03-24

### Added
- `set` command output flags `--out` / `-o` to write results to a separate file.

### Changed
- Improved CLI error reporting: print only errors via cobra and stop printing
  usage on every error.

## [0.1.0-beta] - 2021-05-16

### Added
- Initial release with `get` and `set` commands for querying and editing TOML
  files using dot-notation queries.

[Unreleased]: https://github.com/MinseokOh/toml-cli/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/MinseokOh/toml-cli/compare/v0.1.3...v0.2.0
[0.1.3]: https://github.com/MinseokOh/toml-cli/compare/v0.1.2...v0.1.3
[0.1.2]: https://github.com/MinseokOh/toml-cli/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/MinseokOh/toml-cli/compare/v0.1.0-beta...v0.1.1
[0.1.0-beta]: https://github.com/MinseokOh/toml-cli/releases/tag/v0.1.0-beta
