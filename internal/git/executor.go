package git

import (
	"bytes"
	"errors"
	"os/exec"
)

// Executor is an interface that defines the contract for the git command executor.
type Executor interface {
	Command(arg ...string) (string, error)
	CommandInDir(dir string, arg ...string) (string, error)
}

// executorImpl is an implementation of Executor.
type executorImpl struct{}

// Command implements the Executor interface.
func (e *executorImpl) Command(arg ...string) (string, error) {
	return e.CommandInDir("", arg...)
}

// CommandInDir implements the Executor interface.
func (e *executorImpl) CommandInDir(dir string, arg ...string) (string, error) {
	cmd := exec.Command("git", arg...)

	cmd.Dir = dir
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output

	if err := cmd.Run(); err != nil {
		return "", errors.New(output.String())
	}

	return output.String(), nil
}

// NewExecutor creates a new instance of Executor.
func NewExecutor() Executor {
	return &executorImpl{}
}
