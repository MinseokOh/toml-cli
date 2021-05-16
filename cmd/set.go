package cmd

import (
	"fmt"
	"strconv"

	"github.com/MinseokOh/toml-cli/toml"
	lib "github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
)

const (
	flagData = "v"
)

// SetTomlCommand returns set command
func SetTomlCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set [path] [query]",
		Short: "Edit the file to set some data",
		Long: `
e.g.
toml-cli set ./sample/example.toml title 123456
`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			query := args[1]
			data := parseInput(args[2])

			if data == nil {
				return fmt.Errorf("data is nill")
			}

			toml, err := toml.NewToml(path)
			if err != nil {
				return err
			}

			if err := toml.Set(query, data); err != nil {
				return err
			}

			if err := toml.Write(); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}

func parseInput(str string) interface{} {
	if val, err := strconv.ParseBool(str); err == nil {
		return val
	}

	if val, err := strconv.ParseInt(str, 0, 64); err == nil {
		return val
	}

	if val, err := strconv.ParseFloat(str, 64); err == nil {
		return val
	}

	if val, err := lib.ParseLocalDate(str); err == nil {
		return val
	}

	if val, err := lib.ParseLocalDateTime(str); err == nil {
		return val
	}

	if val, err := lib.ParseLocalTime(str); err == nil {
		return val
	}

	return str
}
