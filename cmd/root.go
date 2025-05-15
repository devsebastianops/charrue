package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	verbose bool
)

var rootCmd = &cobra.Command{
	Short:   "Generates Terraform configuration files",
	Long:    `Charrue is a tool that generates Terraform configuration files from a higher-level config and template directory.`,
	Use:     "charrue",
	Version: "0.1.0",
}

func Execute() {
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add commands
	rootCmd.AddCommand(BuildCmd)

	// Global flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "", false, "Enable verbose output")
}
