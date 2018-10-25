package cmd

import (
	"fmt"
	"strings"

	"github.com/CLI/gophercises/db"
	"github.com/spf13/cobra"
)

// Created the function as variable for test coverage
var FunctionCreate = db.CreateTask

// adds command to list & store in it DB.
// Here we have used cobra libraries to add the tasks
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the todo list tasks",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, "")
		_, err := FunctionCreate(task)
		if err != nil {
			fmt.Println("something went wrong", err)
			return
		}
		fmt.Printf("Added \"%s\" task to your list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
