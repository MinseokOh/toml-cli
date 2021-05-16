package cmd

import (
	"fmt"

	"github.com/MinseokOh/toml-cli/toml"
	"github.com/spf13/cobra"
)

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
			path := args[0]
			query := args[1]

			toml, err := toml.NewToml(path)
			if err != nil {
				return err
			}

			res := toml.Get(query)
			if res == nil {
				return fmt.Errorf("Not exist data")
			}

			fmt.Println(res)
			return nil
		},
	}

	return cmd
}
