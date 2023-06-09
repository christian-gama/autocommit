package git

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

// Commit commits the changes to the git repository.
func Commit(message string) error {
	_, err := executeCommand("git", "commit", "-m", message)
	return err
}

// Diff returns the diff of the changes in the git repository.
func Diff() (string, error) {
	gitRootPath, err := getGitRootPath()
	if err != nil {
		return "", err
	}

	output, err := executeCommandInDir(
		gitRootPath,
		"git",
		"diff",
		"--no-color",
		"--minimal",
		"--cached",
	)
	if err != nil {
		return "", err
	}

	if len(output) == 0 {
		return "", errors.New("There are no changes to commit")
	}

	return output, nil
}

func executeCommand(name string, arg ...string) (string, error) {
	return executeCommandInDir("", name, arg...)
}

func executeCommandInDir(dir, name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = dir
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output

	err := cmd.Run()
	if err != nil {
		return "", errors.New(output.String())
	}

	return output.String(), nil
}

func getGitRootPath() (string, error) {
	output, err := executeCommand("git", "rev-parse", "--show-toplevel")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(output), nil
}
