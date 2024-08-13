package local

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	cp "github.com/otiai10/copy"
)

// Service struct
type Service struct{}

// New return service
func New() *Service {
	return &Service{}
}

// CheckExist func
func (s *Service) CheckExist(originalPath string, backupPath string) (err error) {

	if _, err = os.Stat(originalPath); err != nil {
		err = fmt.Errorf("cannot get path, ERROR: %s", err)
		return
	}

	return
}

// Save fund
func (s *Service) Save(originalPath string, backupPath string) (backupedPath string, err error) {

	backupFolder := fmt.Sprintf("t-backup:%s", time.Now().Format("2006-01-02--15:04:05"))

	backupedPath = filepath.Join(backupPath, backupFolder)

	err = cp.Copy(originalPath, backupedPath)

	return
}

// DeletebackupPath func
func (s *Service) DeletebackupPath(deletedPath string) (err error) {
	err = os.RemoveAll(deletedPath)
	return
}
