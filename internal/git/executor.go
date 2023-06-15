package git

import (
	"bytes"
	"errors"
	"os/exec"
)

type Executor interface {
	Command(arg ...string) (string, error)
	CommandInDir(dir string, arg ...string) (string, error)
}

type executorImpl struct{}

func (e *executorImpl) Command(arg ...string) (string, error) {
	return e.CommandInDir("", arg...)
}

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

func NewExecutor() Executor {
	return &executorImpl{}
}
