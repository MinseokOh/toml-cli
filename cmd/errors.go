package cmd

import "fmt"

var (
	ErrLoad     = fmt.Errorf("failed to load toml")
	ErrBadQuery = fmt.Errorf("asdf")
)
