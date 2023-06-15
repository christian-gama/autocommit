package storage

import (
	"log"
	"os"
	"path"
)

// Storage is a struct that represents a file storage
type Storage struct {
	dir      string
	filename string
}

// Create creates a file with the given content
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

// Read reads the file content
func (s *Storage) Read() ([]byte, error) {
	content, err := os.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// Update updates the file content
func (s *Storage) Update(newContent []byte) error {
	if err := s.Delete(); err != nil {
		return err
	}

	return s.Create(newContent)
}

// Delete deletes the file
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

// NewStorage creates a new instance of Storage
func NewStorage(filename string) *Storage {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return &Storage{dir: path.Join(home, ".autocommit"), filename: path.Join(home, ".autocommit", filename)}
}
