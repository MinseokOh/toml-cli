package cmd

import (
	"fmt"

	"github.com/MinseokOh/toml-cli/toml"
	lib "github.com/pelletier/go-toml/v2"
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

// parseInput infers the TOML type of str by letting go-toml parse it as the
// value of a single key. This covers integers (incl. hex/octal/binary and
// underscores), floats, booleans, dates/times, arrays and inline tables.
// A bare or otherwise unparseable token is treated as a plain string.
func parseInput(str string) any {
	var holder map[string]any
	if err := lib.Unmarshal([]byte("v = "+str), &holder); err == nil {
		return holder["v"]
	}

	return str
}
