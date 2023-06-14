package main

import (
	"log"

	"github.com/christian-gama/autocommit/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
