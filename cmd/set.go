package cmd

import (
	"github.com/MinseokOh/toml-cli/toml"
	"github.com/spf13/cobra"
)

func SetTomlCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set [path] [query] [data]",
		Short: "Edit the file to set some data",
		Long: `
e.g.
toml-cli set ./sample/example.toml title 123456
`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			query := args[1]
			data := args[2]

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
