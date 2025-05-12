package systemmsg

import (
	"embed"
	"errors"
	"os"
	"path"

	"github.com/christian-gama/autocommit/config"
)

const _systemMsgFileName = "system_msg.txt"

//go:embed system_msg.txt
var _defaultSystemMsg embed.FS

func Load() (string, error) {
	content, err := os.ReadFile(msgFilePath())
	if err != nil {
		return "", err
	}

	if len(content) == 0 {
		return "", errors.New("system message is empty")
	}

	return string(content), nil
}

func Restore() error {
	if _, err := os.Stat(config.Dir()); !os.IsNotExist(err) {
		if err := os.Remove(msgFilePath()); err != nil {
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

	content, err := _defaultSystemMsg.ReadFile(_systemMsgFileName)
	if err != nil {
		return err
	}

	return os.WriteFile(msgFilePath(), content, os.ModePerm)
}

func msgFilePath() string {
	return path.Join(config.Dir(), _systemMsgFileName)
}
