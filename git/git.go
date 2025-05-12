package git

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

func Commit(message string) error {
	if err := git("commit", "-m", message).Run(); err != nil {
		return err
	}

	return nil
}

func MinimalDiff() (string, error) {
	output, err := git("diff", "--no-color", "--minimal", "--cached", "-U3").CombinedOutput()
	if err != nil {
		return "", err
	}

	if len(output) == 0 {
		return "", errors.New("there are no changes to commit - did you forget to stage your changes?")
	}

	return string(output), nil
}

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
