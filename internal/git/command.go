package git

import (
	"errors"
	"strings"
)

// baseCommand is a struct that holds the common attributes for all git commands.
type baseCommand struct {
	exec     Executor
	rootPath string
}

func (b *baseCommand) setRootPath() error {
	output, err := b.exec.Command("rev-parse", "--show-toplevel")
	if err != nil {
		return err
	}

	b.rootPath = strings.TrimSpace(output)

	return nil
}

// CommitCommand is an interface that defines the contract for the git commit command.
type CommitCommand interface {
	Execute(message string) error
}

// commitCommandImpl is an implementation of CommitCommand.
type commitCommandImpl struct {
	baseCommand
}

// Execute implements the CommitCommand interface.
func (c *commitCommandImpl) Execute(message string) error {
	_, err := c.exec.Command("commit", "-m", message)

	return err
}

// NewCommitCommand creates a new instance of CommitCommand.
func NewCommitCommand(executor Executor) CommitCommand {
	g := new(commitCommandImpl)
	g.exec = executor

	if err := g.setRootPath(); err != nil {
		panic(err)
	}

	return g
}

// DiffCommand is an interface that defines the contract for the git diff command.
type DiffCommand interface {
	Execute() (string, error)
}

// diffCommandImpl is an implementation of DiffCommand.
type diffCommandImpl struct {
	baseCommand
}

// Execute implements the DiffCommand interface.
func (d *diffCommandImpl) Execute() (string, error) {
	output, err := d.exec.CommandInDir(
		d.rootPath,
		"diff", "--no-color", "--cached",
	)

	if err != nil {
		return "", err
	}

	if len(output) == 0 {
		return "", errors.New("There are no changes to commit - did you forget to stage your changes?")
	}

	return output, nil
}

// NewDiffCommand creates a new instance of DiffCommand.
func NewDiffCommand(executor Executor) DiffCommand {
	g := new(diffCommandImpl)
	g.exec = executor

	if err := g.setRootPath(); err != nil {
		panic(err)
	}

	return g
}
