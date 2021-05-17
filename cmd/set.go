package cmd

import (
	"fmt"
	"strconv"

	"github.com/MinseokOh/toml-cli/toml"
	lib "github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
)

const (
	flagOut = "out"
)

// SetTomlCommand returns set command
func SetTomlCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set [path] [query] [data]",
		Short: "Edit the file to set some data",
		Long: `
e.g.
toml-cli set ./sample/example.toml title 123456

e.g.
toml-cli set ./sample/example.toml title 123456 --out=./sample/example_out.toml
`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			query := args[1]
			data := parseInput(args[2])
			outDir, _ := cmd.Flags().GetString(flagOut)

			if data == nil {
				return fmt.Errorf("data is nill")
			}

			toml, err := toml.NewToml(path)
			if err != nil {
				return err
			}

			toml.Dest(outDir)

			if err := toml.Set(query, data); err != nil {
				return err
			}

			if err := toml.Write(); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringP(flagOut, "o", "asdf", "set output directory")
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
