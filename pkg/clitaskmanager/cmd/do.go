package cmd

import (
	"fmt"
	"strconv"

	"github.com/mrinaald/my-gophercises/pkg/clitaskmanager/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			taskId, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Invalid Task Id: \"%s\". Ignoring\n", arg)
			} else {
				ids = append(ids, taskId)
			}
		}

		tasks, err := db.GetTaskList()
		if err != nil {
			fmt.Println("Something went wrong:", err)
		}

		var taskIds []int
		var taskDescriptions []string
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Printf("Task number \"%d\" out of range. Ignoring!\n", id)
				continue
			}
			taskIds = append(taskIds, tasks[id-1].Key)
			taskDescriptions = append(taskDescriptions, tasks[id-1].Task)
		}

		if err := db.CompleteTask(taskIds); err != nil {
			fmt.Println("Something went wrong while deleting tasks:", err)
			return
		}

		for _, task := range taskDescriptions {
			fmt.Printf("You have completed the \"%s\" task.\n", task)
		}
	},
}
