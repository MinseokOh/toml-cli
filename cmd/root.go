package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = GetRootCommand()

// GetRootCommand returns root command
func GetRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "toml-cli",
		Short: "toml-cli",
		SilenceUsage: true,
		Long: `A simple CLI for editing and querying TOML files.
	`,
	}

	return rootCmd
}

func init() {
	rootCmd.AddCommand(GetTomlCommand())
	rootCmd.AddCommand(SetTomlCommand())
}

// Execute commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
