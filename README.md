<p align="center">
  <h2 align="center">toml-cli</h2>
  <p align="center">A simple CLI for editing and querying TOML files.</p>
  <p align="center">✨ <strong>Enhanced with Claude Code</strong> ✨</p>
</p>

<p align="center">
<a href="https://pkg.go.dev/github.com/MinseokOh/toml-cli" target="blank">
<img src="https://pkg.go.dev/badge/github.com/MinseokOh/toml-cli.svg"/>
</a>
<a href="https://github.com/MinseokOh/toml-cli/blob/master/LICENSE" target="blank">
<img src="https://img.shields.io/badge/licence-MIT-green?style=flat-square"/>
</a>
<a href="https://goreportcard.com/report/github.com/MinseokOh/toml-cli" target="blank">
<img src="https://goreportcard.com/badge/github.com/MinseokOh/toml-cli"/>
</a>  
</p>
This is the home of the toml command, a simple CLI for editing and querying TOML files.

Idea from [gnprince/toml-cli](https://github.com/gnprice/toml-cli), written in golang. and depends on [pelletier/go-toml](https://github.com/pelletier/go-toml).

The intent of the toml command is to be useful
- in shell scripts, for consulting or editing a config file
- and in instructions a human can follow for editing a config file, as a command to copy-paste and run.


## Sample Files Structure

The project includes well-organized sample files with comprehensive test cases:

```
sample/
├── get-set/          # Examples for get and set operations
│   ├── README.md     # Detailed test cases and usage examples
│   └── config.toml   # Simple configuration example
└── merge/            # Examples for merge operations
    ├── README.md     # Detailed test cases and usage examples
    ├── base.toml     # Base configuration
    └── override.toml # Override configuration
```

## Usage

### `Get`
```shell
$ toml-cli get ./sample/get-set/config.toml server.port
8080

$ toml-cli get ./sample/get-set/config.toml app.name
toml-cli-demo
```

### `Set`
```shell
$ toml-cli set ./sample/get-set/config.toml server.port 3000
# modify server.port to 3000

$ toml-cli set ./sample/get-set/config.toml features.maintenance_mode true -o ./sample/get-set/config_output.toml
# modify and save to new file
```

### `Merge`
```shell
$ toml-cli merge ./sample/merge/base.toml ./sample/merge/override.toml -o ./sample/merge/base_output.toml
# merge base with override configuration

$ toml-cli merge ./sample/merge/base.toml ./sample/merge/override.toml
# merge and save to base_output.toml
```

> 💡 **Tip**: Check the `sample/*/README.md` files for comprehensive test cases and examples!

## Development

This project is actively developed with [Claude Code](https://claude.ai/code) for enhanced productivity and code quality. The codebase includes comprehensive documentation in `CLAUDE.md` to facilitate AI-assisted development.

## Contributors

<a href="https://github.com/MinseokOh/toml-cli/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=MinseokOh/toml-cli" />
</a>

Made with [contributors-img](https://contrib.rocks).
