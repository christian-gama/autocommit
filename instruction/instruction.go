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

func Restore() error {
	if _, err := os.Stat(config.Dir()); !os.IsNotExist(err) {
		if err := os.Remove(filePath()); err != nil {
			return err
		}
	}

	return Create()
}

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
