# groq-go

[![Go Reference](https://pkg.go.dev/badge/github.com/hasitpbhatt/groq-go.svg)](https://pkg.go.dev/github.com/hasitpbhatt/groq-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/hasitpbhatt/groq-go)](https://goreportcard.com/report/github.com/hasitpbhatt/groq-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

groq-go is a Go client library for interacting with the Groq API. This library provides a simple and efficient way to integrate Groq's powerful language models into your Go applications.

## Features

- Easy-to-use client for Groq API
- Support for chat completions
- Customizable API requests

### Test Example

Here's an example of how to use the `ChatCompletion` function to get chat completions:
```go
package main

import (
	"fmt"

	groq "github.com/hasitpbhatt/groq-go"
)

func main() {

	// Automatically read API key from GROQ_API_KEY environment variable
	client := groq.NewClient()

	resp, err := client.ChatCompletion([]groq.Message{
		{
			Content: "You're a seasoned developer",
			Role:    "system",
		},
		{
			Content: "What is groq cloud?",
			Role:    "user",
		},
	})
	if err != nil {
		fmt.Println("Error occurred")
	}

	for _, c := range resp.Choices {
		fmt.Println(c.Message.Content)
	}
}
```


## Installation

To install the package, use `go get github.com/hasitpbhatt/groq-go`:
