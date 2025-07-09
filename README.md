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


## Usage

### `Get`
```shell
$ toml-cli get ./sample/example.toml owner.dob
1979-05-27 07:32:00 -0800 -0800
```

### `Set`
```shell
$ toml-cli set ./sample/example.toml owner.name MinseokOh
# modify owner.name to MinseokOh
```

## Development

This project is actively developed with [Claude Code](https://claude.ai/code) for enhanced productivity and code quality. The codebase includes comprehensive documentation in `CLAUDE.md` to facilitate AI-assisted development.

## Contributors

<a href="https://github.com/MinseokOh/toml-cli/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=MinseokOh/toml-cli" />
</a>

Made with [contributors-img](https://contrib.rocks).
