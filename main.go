package main

import (
	"log"

	"github.com/christian-gama/autocommit/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
