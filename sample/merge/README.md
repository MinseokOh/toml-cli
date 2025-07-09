# Merge Examples

This directory contains examples for the `merge` command.

## Files

- `base.toml` - Base configuration file
- `override.toml` - Override configuration file
- `merged_output.toml` - Example output after merging

## Usage Examples

### Merge Command
```bash
./toml-cli merge ./sample/merge/base.toml ./sample/merge/override.toml -o ./sample/merge/merged_output.toml
# Creates merged_output.toml with configurations merged recursively
```

The merge operation combines the two TOML files, with values from `override.toml` taking precedence over `base.toml`.