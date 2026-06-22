package cmd

import (
	"fmt"
	"os"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/features/cisco_ios"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check <file>",
	Short: "Analyze a configuration file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading file: %w", err)
		}

		feature := cisco_ios.New()
		defer feature.Close()

	doc := document.New(path, "cisco_ios", 0, content)
	diagnostics, err := feature.DidOpen(cmd.Context(), doc)
	if err != nil {
		return fmt.Errorf("parsing file: %w", err)
	}

	if len(diagnostics) == 0 {
			fmt.Println("No issues found.")
			return nil
		}

		for _, d := range diagnostics {
			fmt.Printf("%s:%d:%d: [%s] %s\n",
				path,
				d.Range.Start.Line+1,
				d.Range.Start.Character+1,
				d.Source,
				d.Message,
			)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
