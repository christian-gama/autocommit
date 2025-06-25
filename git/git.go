// Package git provides Git operations for the autocommit tool.
package git

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

// Commit creates a Git commit with the provided message.
// It uses the staged changes in the Git repository.
func Commit(message string) error {
	cmd := git("commit", "-m", message)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// MinimalDiff returns the Git diff of the staged changes.
// It returns an error if there are no staged changes.
func MinimalDiff() (string, error) {
	output, err := git("diff", "--no-color", "--minimal", "--cached", "-U3").CombinedOutput()
	if err != nil {
		return "", err
	}

	if len(output) == 0 {
		return "", errors.New(
			"there are no changes to commit - did you forget to stage your changes?",
		)
	}

	return string(output), nil
}

// ListFiles returns a formatted string representation of all files in the Git repository.
// The output is organized as a directory tree structure.
func ListFiles() (string, error) {
	output, err := git("rev-parse", "--show-toplevel").CombinedOutput()
	if err != nil {
		return "", err
	}

	if len(output) == 0 {
		return "", errors.New("could not find the root path of the repository")
	}

	rootPath := strings.TrimSpace(string(output))

	output, err = git("-C", rootPath, "ls-files").CombinedOutput()
	if err != nil {
		return "", err
	}

	if len(output) == 0 {
		return "", errors.New("there are no files to list")
	}

	dirs := make(map[string][]string)

	files := strings.SplitSeq(string(output), "\n")
	for file := range files {
		dir := filepath.Dir(file)
		if dir == "." {
			dirs[dir] = append(dirs[dir], file)
		} else {
			dirs[dir] = append(dirs[dir], filepath.Base(file))
		}
	}

	return writeDir(new(strings.Builder), dirs, ".", 0), nil
}

func writeDir(
	sb *strings.Builder,
	dirs map[string][]string,
	currentDir string,
	level int,
) string {
	indent := strings.Repeat("  ", level)

	if currentDir != "." {
		fmt.Fprintf(
			sb,
			"%s📁 %s\n", indent, filepath.Base(currentDir),
		)
	}

	for _, file := range dirs[currentDir] {
		if !strings.Contains(
			file,
			"/",
		) {
			fmt.Fprintf(sb,
				"%s📄 %s\n", strings.Repeat("  ", level+1), file)
		}
	}

	subdirs := make([]string, 0)
	prefix := currentDir + "/"
	if currentDir == "." {
		prefix = ""
	}

	for dir := range dirs {
		if dir != "." && strings.HasPrefix(dir, prefix) &&
			!strings.Contains(dir[len(prefix):], "/") {
			subdirs = append(subdirs, dir)
		}
	}
	sort.Strings(subdirs)

	for _, subdir := range subdirs {
		writeDir(sb, dirs, subdir, level+1)
	}

	return sb.String()
}

func git(args ...string) *exec.Cmd {
	return exec.Command("git", args...)
}
