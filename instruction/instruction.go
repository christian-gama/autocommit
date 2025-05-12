// Package instruction provides functionality to manage the instruction template used by the LLM.
package instruction

import (
	"embed"
	"errors"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/christian-gama/autocommit/config"
)

const _instructionFileName = "instruction.txt"

//go:embed instruction.txt
var _defaultInstruction embed.FS

// Load reads the instruction template from disk. If the file doesn't exist,
// it creates a new instruction file with default content first.
func Load() (string, error) {
	content, err := os.ReadFile(filePath())
	if err != nil {
		if os.IsNotExist(err) {
			if err := Create(); err != nil {
				return "", err
			}

			content, err = os.ReadFile(filePath())
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}

	if len(content) == 0 {
		return "", errors.New("empty instruction file")
	}

	return string(content), nil
}

// Restore resets the instruction file back to its default content,
// removing any customizations.
func Restore() error {
	if _, err := os.Stat(config.Dir()); !os.IsNotExist(err) {
		if err := os.Remove(filePath()); err != nil {
			return err
		}
	}

	return Create()
}

// Create writes the default instruction template to disk.
// Creates the configuration directory if it doesn't exist.
func Create() error {
	if _, err := os.Stat(config.Dir()); os.IsNotExist(err) {
		if err := os.MkdirAll(config.Dir(), os.ModePerm); err != nil {
			return err
		}
	}

	content, err := _defaultInstruction.ReadFile(_instructionFileName)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath(), content, os.ModePerm)
}

// Open launches the instruction file in the default text editor
// for the current operating system.
func Open() error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", filePath())
	case "linux", "darwin":
		cmd = exec.Command("open", filePath())
	default:
		return errors.New("unsupported platform")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func filePath() string {
	return path.Join(config.Dir(), _instructionFileName)
}
