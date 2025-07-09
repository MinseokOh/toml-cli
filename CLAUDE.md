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
   - `toml.go`: Core TOML wrapper around pelletier/go-toml library
   - `tomlw.go`: File I/O operations for reading/writing TOML files
   - `keysparsing.go`: Key parsing utilities for dot notation queries
   - `lexer.go`: TOML lexical analysis
   - `position.go` & `token.go`: Supporting lexer infrastructure

3. **Entry Point**:
   - `main.go`: Simple entry point that calls cmd.Execute()

### Key Dependencies
- `github.com/pelletier/go-toml`: Core TOML parsing library
- `github.com/spf13/cobra`: CLI framework
- `github.com/stretchr/testify`: Testing framework

### Data Flow
1. Commands parse arguments using cobra
2. Create `Toml` struct instance from file path
3. For `get`: Query value using dot notation and print result
4. For `set`: Update value, optionally specify output file, then write changes
5. For `merge`: Load two TOML files, merge recursively, then write combined result

### Value Type Handling
The `set` command automatically detects and converts input types:
- Boolean values (true/false)
- Integer values (int64)
- Float values (float64)
- TOML date/time formats (LocalDate, LocalDateTime, LocalTime)
- Strings (fallback)

## Testing Strategy

The codebase includes comprehensive tests in the `toml/` package covering:
- Key parsing with various formats (bare keys, quoted keys, dotted keys)
- TOML lexing and tokenization
- Core TOML operations (get/set)
- Date/time parsing
- Array and inline table handling
- Unicode support
- Error cases and edge conditions