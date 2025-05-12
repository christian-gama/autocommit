package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

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

	clearScreen()

	if err := cli.AutoCommit.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}

// clearScreen clears the terminal screen. It uses platform-specific
// commands based on the runtime environment.
func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}
}
