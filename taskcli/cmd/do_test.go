package cmd

import (
	"errors"
	"testing"

	"github.com/CLI/gophercises/db"
)

// This function covers the unit tests when we have pass blank with do command
func TestDoCmd(t *testing.T) {

	tmplist := FunctionList
	defer func() {
		FunctionList = tmplist
	}()

	var temp []db.Task
	FunctionList = func() ([]db.Task, error) {
		return temp, errors.New("Error while Listing")
	}
	a := []string{""}
	ListCmd.Run(ListCmd, a)

	tmpdo := FunctionDo
	defer func() {
		FunctionDo = tmpdo
	}()

	FunctionDo = func(int) error {
		return errors.New("Error while completing task")
	}

	DoCmd.Run(DoCmd, a)
}

// This function covers the unit tests when have pass any value (i.e '1') with do command
func TestDoCmdPositive(t *testing.T) {

	tmpdo := FunctionDo
	defer func() {
		FunctionDo = tmpdo
	}()

	FunctionDo = func(int) error {
		return errors.New("Error while completing task")
	}
	a := []string{"1"}
	DoCmd.Run(DoCmd, a)

}

// This function covers the unit tests when we have pass a negative value(i.e 0) with do command

func TestDoCmdNegative(t *testing.T) {

	tmpdo := FunctionDo
	defer func() {
		FunctionDo = tmpdo
	}()

	FunctionDo = func(int) error {
		return errors.New("Error while completing task")
	}
	a := []string{"0"}
	DoCmd.Run(DoCmd, a)

}

/*func TestDoCmdMark(t *testing.T) {

	tmpdo := FunctionDo
	defer func() {
		FunctionDo = tmpdo
	}()

	FunctionDo = func(int) error {
		return errors.New("Error while completing task")
	}
	a := []string{}
	DoCmd.Run(DoCmd, a)

}*/
