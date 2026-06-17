# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a CLI tool for editing and querying TOML files, written in Go. It provides three main commands:
- `get`: Extract values from TOML files using dot notation queries
- `set`: Update values in TOML files with optional output file specification
- `merge`: Merge two TOML files with recursive object merging

## Development Commands

### Build and Run
```bash
go build -o toml-cli
./toml-cli get ./sample/get-set/app.toml server.port
./toml-cli set ./sample/get-set/app.toml server.port 3000 [-o output_file]
./toml-cli merge ./sample/merge/base.toml ./sample/merge/override.toml [-o output_file]
```

### Testing
```bash
go test -v ./...           # Run all tests
go test -v ./toml          # Run specific package tests
```

### Standard Go Commands
```bash
go run main.go <args>      # Run without building
go build                  # Build binary
go mod tidy               # Clean up dependencies
go fmt ./...              # Format code
```

## Architecture

### Core Components

1. **Command Structure** (`cmd/`):
   - `root.go`: Main cobra command setup and execution
   - `get.go`: Implementation of get command for querying TOML values
   - `set.go`: Implementation of set command for updating TOML values
   - `merge.go`: Implementation of merge command for combining TOML files

2. **TOML Processing** (`toml/`):
   - `toml.go`: Core TOML wrapper. Documents are unmarshalled into a
     `map[string]any` tree; `Get`/`Has`/`Set`/`Merge` operate on it via
     dot-notation key traversal.
   - `tomlw.go`: File I/O operations for reading/writing TOML files

3. **Entry Point**:
   - `main.go`: Simple entry point that calls cmd.Execute()

### Key Dependencies
- `github.com/pelletier/go-toml/v2`: Core TOML parsing library (marshal/unmarshal)
- `github.com/spf13/cobra`: CLI framework
- `github.com/stretchr/testify`: Testing framework

### Data Flow
1. Commands parse arguments using cobra
2. Create `Toml` struct instance from file path
3. For `get`: Query value using dot notation and print result
4. For `set`: Update value, optionally specify output file, then write changes
5. For `merge`: Load two TOML files, merge recursively, then write combined result

### Value Type Handling
The `set` command infers the type of the input value by delegating to go-toml:
`parseInput` parses the argument as the value of a single TOML key
(`v = <input>`) and returns whatever go-toml produces. This covers the full
TOML type system in one step:
- Booleans (`true`/`false`)
- Integers including hex/octal/binary literals and underscores (`0xFF`, `1_000`)
- Floats (float64)
- Date/time values (LocalDate, LocalDateTime, LocalTime)
- Arrays and inline tables
- Any bare or unparseable token falls back to a plain string

Note: output formatting follows go-toml v2 conventions — keys are emitted in
alphabetical order and strings use single quotes; comments are not preserved on
round-trip.

## Testing Strategy

Tests live in `toml/toml_test.go` and cover the core wrapper operations:
- `TestNewToml`: loading a file (and the error path for a missing file)
- `TestGet`: dot-notation queries with type assertions (e.g. `int64` ports)
- `TestSet`: updating a value and writing it back out
- `TestMerge`: recursive merge with override-wins and base-preserved semantics