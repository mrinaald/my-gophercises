package clitaskmanager

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task [command]",
	Short: "task is a CLI for managing your TODOs.",
	Long: `task is a command-line interface tool for
managing TODO lists.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
