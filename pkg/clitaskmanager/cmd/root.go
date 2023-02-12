package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI for managing your TODOs.",
	Long: `task is a command-line interface tool for
managing TODO lists.`,
}

func Execute() error {
	return rootCmd.Execute()
}
