package git

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

// baseCommand is a struct that holds the common attributes for all git commands.
type Git struct {
	cmd *exec.Cmd
}

func New() *Git {
	return &Git{}
}

func (c *Git) Commit(message string) error {
	_, err := c.command("commit", "-m", message)

	return err
}

func (d *Git) Diff() (string, error) {
	output, err := d.commandInDir(
		d.rootPath(),
		"diff", "--no-color", "--minimal", "--cached", "-U3",
	)

	if err != nil {
		return "", err
	}

	if len(output) == 0 {
		return "", errors.New(
			"there are no changes to commit - did you forget to stage your changes?",
		)
	}

	return output, nil
}

func (g *Git) Log(n int) (string, error) {
	output, err := g.commandInDir(
		g.rootPath(),
		"log", "--pretty=format:%s", "--no-color", "--no-merges", "--no-abbrev-commit", "-n", fmt.Sprint(n),
	)

	if err != nil {
		return "", err
	}

	if len(output) == 0 {
		return "", errors.New("there are no commits in this repository")
	}

	return output, nil
}

// ListProjectStructure returns a string representation of the project structure
func (g *Git) ListProjectStructure() (string, error) {
	output, err := g.command("ls-files")
	if err != nil {
		return "", err
	}

	dirs := make(map[string][]string)
	files := strings.SplitSeq(strings.TrimSpace(string(output)), "\n")

	for file := range files {
		dir := filepath.Dir(file)
		if dir == "." {
			dirs[dir] = append(dirs[dir], file)
		} else {
			dirs[dir] = append(dirs[dir], filepath.Base(file))
		}
	}

	// Build the tree structure
	var structure strings.Builder
	printDir(&structure, dirs, ".", 0)

	return structure.String(), nil
}

func (e *Git) command(arg ...string) (string, error) {
	return e.commandInDir("", arg...)
}

func (g *Git) commandInDir(dir string, arg ...string) (string, error) {
	g.cmd = exec.Command("git", arg...)

	g.cmd.Dir = dir
	var output bytes.Buffer
	g.cmd.Stdout = &output
	g.cmd.Stderr = &output

	if err := g.cmd.Run(); err != nil {
		return "", errors.New(output.String())
	}

	return output.String(), nil
}

func (g *Git) rootPath() string {
	output, err := g.command("rev-parse", "--show-toplevel")
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(output)
}

func printDir(
	sb *strings.Builder,
	dirs map[string][]string,
	currentDir string,
	level int,
) {
	indent := strings.Repeat("  ", level)

	if currentDir != "." {
		fmt.Fprintf(sb,
			"%s📁 %s\n", indent, filepath.Base(currentDir))
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
		printDir(sb, dirs, subdir, level+1)
	}
}
