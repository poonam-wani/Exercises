package cmd

import (
	"fmt"
	"testing"
)

func TestRootCmd(t *testing.T) {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println("Error in root")
	}
}
