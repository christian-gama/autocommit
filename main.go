package main

import (
	"fmt"
	"os"

	"github.com/christian-gama/autocommit/v2/cli"
)

func main() {
	if err := cli.AutoCommitCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
