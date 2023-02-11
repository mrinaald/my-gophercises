package main

import (
	"github.com/mrinaald/my-gophercises/pkg/clitaskmanager"
)

func main() {
	clitaskmanager.InitailizeDB()
	defer clitaskmanager.FinalizeDB()

	clitaskmanager.Execute()
}
