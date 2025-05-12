package main

import (
	"fmt"

	"github.com/christian-gama/autocommit/cli"
)

func main() {
	if err := cli.AutoCommit.Execute(); err != nil {
		fmt.Println(err)
	}
}
