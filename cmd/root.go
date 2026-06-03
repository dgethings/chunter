package cmd

import (
	"github.com/spf13/cobra"
)

var logLevel string

var rootCmd = &cobra.Command{
	Use:   "chunter",
	Short: "Language server for network operating systems",
}

var Version string

func init() {
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Log level")
}

func Execute() error {
	return rootCmd.Execute()
}
