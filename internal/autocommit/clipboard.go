package autocommit

import "github.com/atotto/clipboard"

type Clipboard interface {
	Copy(message string) error
}

type clipboardImpl struct{}

func (c *clipboardImpl) Copy(message string) error {
	return clipboard.WriteAll(message)
}

func NewClipboard() Clipboard {
	return &clipboardImpl{}
}
