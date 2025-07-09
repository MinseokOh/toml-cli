package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/MinseokOh/toml-cli/toml"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge <source1> <source2>",
	Short: "Merge two TOML files",
	Long: `Merge two TOML files. The second file's values will overwrite the first file's values.
Nested objects are merged recursively. Arrays are replaced entirely.

Examples:
  toml-cli merge config.toml defaults.toml
  toml-cli merge config.toml override.toml -o merged.toml
`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		source1Path := args[0]
		source2Path := args[1]

		outDir, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		// Load first TOML file (base)
		base, err := toml.NewToml(source1Path)
		if err != nil {
			return fmt.Errorf("failed to load base file %s: %v", source1Path, err)
		}

		// Load second TOML file (overlay)
		overlay, err := toml.NewToml(source2Path)
		if err != nil {
			return fmt.Errorf("failed to load overlay file %s: %v", source2Path, err)
		}

		// Set output path
		if outDir != "" {
			base.Out(outDir)
		} else {
			// Default to first file's name with _merged suffix
			ext := filepath.Ext(source1Path)
			name := source1Path[:len(source1Path)-len(ext)]
			base.Out(name + "_merged" + ext)
		}

		// Merge overlay into base
		if err := base.Merge(&overlay); err != nil {
			return fmt.Errorf("failed to merge files: %v", err)
		}

		// Write the merged result
		if err := base.Write(); err != nil {
			return fmt.Errorf("failed to write merged file: %v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
	mergeCmd.Flags().StringP("output", "o", "", "Output file path")
}
