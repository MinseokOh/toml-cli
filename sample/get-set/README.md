# Get-Set Examples

This directory contains examples for the `get` and `set` commands.

## Files

- `app.toml` - Base configuration file
- `app_output.toml` - Example output after running set command

## Usage Examples

### Get Command
```bash
./toml-cli get ./sample/get-set/app.toml server.port
# Output: 8080
```

### Set Command
```bash
./toml-cli set ./sample/get-set/app.toml server.port 3000 -o ./sample/get-set/app_output.toml
# Creates app_output.toml with server.port changed to 3000
```