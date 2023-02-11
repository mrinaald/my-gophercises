package clitaskmanager

import (
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Long:  `Add a new task to your TODO list`,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		AddTaskInDB(task)
	},
}
