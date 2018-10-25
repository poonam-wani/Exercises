package main

import (
	"path/filepath"

	"github.com/CLI/gophercises/cmd"
	"github.com/CLI/gophercises/db"
	homeDir "github.com/mitchellh/go-homedir"
)

func main() {

	home, _ := homeDir.Dir()
	dbPath := filepath.Join(home, "tests.db") // Created the database with name tests.db
	db.Init(dbPath)
	cmd.RootCmd.Execute()

}
