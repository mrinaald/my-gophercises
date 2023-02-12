package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mrinaald/my-gophercises/pkg/clitaskmanager/db"
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

		_, err := db.AddTask(task)
		if err != nil {
			fmt.Println("Something went wrong while creating task:", err)
			os.Exit(1)
		}

		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}
