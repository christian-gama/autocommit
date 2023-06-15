package main

import (
	"fmt"

	"github.com/christian-gama/autocommit/cmd"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			const redColor = "\033[31m"
			const resetColor = "\033[0m"
			fmt.Printf("%sX Sorry, something went wrong: %s%s\n", redColor, err, resetColor)
		}
	}()

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
