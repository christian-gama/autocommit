package autocommit

import (
	"github.com/christian-gama/autocommit/internal/storage"
)

// SystemMsgRepo is the interface that wraps the basic operations with the system message file.
type SystemMsgRepo interface {
	// SaveSystemMsg saves the system message file.
	SaveSystemMsg() error

	// GetSystemMsg returns the system message file.
	GetSystemMsg() (string, error)

	// Exists returns true if the system message file exists.
	Exists() bool
}

// systemMsgRepoImpl is an implementation of SystemMsgRepo.
type systemMsgRepoImpl struct {
	storage *storage.Storage
}

// GetSystemMsg implements the SystemMsgRepo interface.
func (r *systemMsgRepoImpl) GetSystemMsg() (string, error) {
	content, err := r.storage.Read()
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// SaveSystemMsg implements the SystemMsgRepo interface.
func (r *systemMsgRepoImpl) SaveSystemMsg() error {
	return r.storage.Create([]byte(SystemMsg))
}

// Exists implements the SystemMsgRepo interface.
func (r *systemMsgRepoImpl) Exists() bool {
	content, err := r.storage.Read()
	if err != nil {
		return false
	}

	return len(content) > 0
}

// NewSystemMsgRepo creates a new instance of SystemMsgRepo.
func NewSystemMsgRepo(storage *storage.Storage) SystemMsgRepo {
	return &systemMsgRepoImpl{storage: storage}
}
