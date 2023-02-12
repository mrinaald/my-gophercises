package cmd

import (
	"fmt"
	"os"

	"github.com/mrinaald/my-gophercises/pkg/clitaskmanager/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Long:  `List all of your incomplete tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetTaskList()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You don't have any tasks. Why not take a vacation?")
			return
		}

		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Task)
		}
	},
}
