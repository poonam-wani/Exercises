package cmd

import (
	"errors"
	"testing"

	"github.com/CLI/gophercises/db"
)

// This function covers the unit tests when haven't pass any value with list command
func TestListCmd(t *testing.T) {

	tmpList := FunctionList
	defer func() {
		FunctionList = tmpList
	}()

	var temp []db.Task
	FunctionList = func() ([]db.Task, error) {
		return temp, errors.New("Error while Listing")
	}
	a := []string{}
	ListCmd.Run(ListCmd, a)
}

// This function covers the unit tests when we have passed the blank string as argument
func TestListCmdEmpty(t *testing.T) {

	tmpList := FunctionList
	defer func() {
		FunctionList = tmpList
	}()

	var temp []db.Task
	FunctionList = func() ([]db.Task, error) {
		return temp, errors.New("Error while Listing passing empty list")
	}
	a := []string{""}
	ListCmd.Run(ListCmd, a)
}

// This function covers the positive part of list command
func TestListCmdPositive(t *testing.T) {

	a := []string{}
	ListCmd.Run(ListCmd, a)

}
