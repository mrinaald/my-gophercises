package clitaskmanager

import (
	"fmt"
	"os"
	"strconv"

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
		taskIds := make([]int, 0)
		for _, arg := range args {
			taskId, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid Task Id: \"%s\". Ignoring\n", arg)
			} else {
				taskIds = append(taskIds, taskId)
			}
		}

		CompleteTaskInDB(taskIds)
	},
}
