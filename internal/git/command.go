package git

import (
	"errors"
	"log"
	"strings"
)

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

type CommitCommand interface {
	Execute(message string) error
}

type commitCommandImpl struct {
	baseCommand
}

func (c *commitCommandImpl) Execute(message string) error {
	_, err := c.exec.Command("commit", "-m", message)

	return err
}

func NewCommitCommand(executor Executor) CommitCommand {
	g := new(commitCommandImpl)
	g.exec = executor

	if err := g.setRootPath(); err != nil {
		log.Fatal(err)
	}

	return g
}

type DiffCommand interface {
	Execute() (string, error)
}

type diffCommandImpl struct {
	baseCommand
}

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

func NewDiffCommand(executor Executor) DiffCommand {
	g := new(diffCommandImpl)
	g.exec = executor

	if err := g.setRootPath(); err != nil {
		log.Fatal(err)
	}

	return g
}
