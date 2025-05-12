package instruction

import (
	"embed"
	"errors"
	"os"
	"path"

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

func filePath() string {
	return path.Join(config.Dir(), _instructionFileName)
}
