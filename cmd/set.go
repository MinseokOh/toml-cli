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
toml-cli set ./sample/example.toml title 123456 -o ./sample/example_out.toml
`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			data := parseInput(args[2])
			if data == nil {
				return fmt.Errorf("data is nil")
			}

			outDir, err := cmd.Flags().GetString(flagOut)
			if err != nil {
				return err
			}

			toml, err := toml.NewToml(args[0])
			if err != nil {
				return err
			}

			toml.Out(outDir)

			if err := toml.Set(args[1], data); err != nil {
				return err
			}

			return toml.Write()
		},
	}

	cmd.Flags().StringP(flagOut, "o", "", "set output directory")
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
