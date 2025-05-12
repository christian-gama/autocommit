package main

import (
	"fmt"
	"os"

	"github.com/christian-gama/autocommit/cli"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			const redColor = "\033[31m"
			const resetColor = "\033[0m"
			fmt.Printf(
				"%sX Sorry, something went wrong: %s%s\n",
				redColor,
				err,
				resetColor,
			)
		}
	}()

	if err := cli.AutoCommit.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
