package ls

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

// ListProjectStructure returns a string representation of the project structure
func ListProjectStructure() (string, error) {
	cmd := exec.Command("git", "ls-files")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute git ls-files: %w", err)
	}

	dirs := make(map[string][]string)
	files := strings.Split(strings.TrimSpace(string(output)), "\n")

	for _, file := range files {
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
