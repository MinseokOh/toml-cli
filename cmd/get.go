package cmd

import (
	"fmt"

	"github.com/MinseokOh/toml-cli/toml"
	"github.com/spf13/cobra"
)

// GetTomlCommand returns get command
func GetTomlCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [path] [query]",
		Short: "Print some data from the file",
		Long: `
e.g.
toml-cli get ./sample/example.toml title
TOML Example
`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			toml, err := toml.NewToml(args[0])
			if err != nil {
				return err
			}

			res := toml.Get(args[1])
			if res == nil {
				return fmt.Errorf("Key %v does not exist in %v", args[1], args[0])
			}

			fmt.Println(res)
			return nil
		},
	}

	return cmd
}
