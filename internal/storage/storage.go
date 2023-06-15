package storage

import (
	"log"
	"os"
	"path"
)

type Storage struct {
	dir      string
	filename string
}

func (s *Storage) Create(content []byte) error {
	s.createDirIfNotExists()

	file, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(string(content))
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Read() ([]byte, error) {
	content, err := os.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (s *Storage) Update(newContent []byte) error {
	if err := s.Delete(); err != nil {
		return err
	}

	return s.Create(newContent)
}

func (s *Storage) Delete() error {
	return os.Remove(s.filename)
}

func (s *Storage) createDirIfNotExists() {
	if !s.dirExists() {
		os.MkdirAll(s.dir, os.ModePerm)
	}
}

func (s *Storage) dirExists() bool {
	if _, err := os.Stat(s.dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func NewStorage(filename string) *Storage {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return &Storage{dir: path.Join(home, ".autocommit"), filename: path.Join(home, ".autocommit", filename)}
}
