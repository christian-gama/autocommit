package autocommit

import "github.com/atotto/clipboard"

// Clipboard is an interface that defines the Copy method.
type Clipboard interface {
	// Copy copies the given message to the clipboard.
	Copy(message string) error
}

// clipboardImpl is an implementation of Clipboard.
type clipboardImpl struct{}

// Copy implements the Clipboard interface.
func (c *clipboardImpl) Copy(message string) error {
	return clipboard.WriteAll(message)
}

// NewClipboard creates a new instance of Clipboard.
func NewClipboard() Clipboard {
	return &clipboardImpl{}
}
