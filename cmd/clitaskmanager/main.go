package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/mrinaald/my-gophercises/pkg/clitaskmanager/cmd"
	"github.com/mrinaald/my-gophercises/pkg/clitaskmanager/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")

	must(db.InitailizeDB(dbPath))
	defer db.CloseDB()

	must(cmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
