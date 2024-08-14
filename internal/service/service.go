package service

import (
	"errors"

	"github.com/olegtemek/t-backup/internal/config"
	"github.com/olegtemek/t-backup/internal/models"
)

type strategy interface {
	CheckExist(originalPath string, backupPath string) (err error)
	Save(originalPath string, backupPath string) (backupedPath string, err error)
	DeletebackupPath(deletedPath string) (err error)
}

type repository interface {
	Save(backupedPath string) (deletedPath string, err error)
}

// Service struct
type Service struct {
	cfg  *config.Config
	st   strategy
	repo repository
}

// New return service
func New(cfg *config.Config, repo repository, st strategy) *Service {
	return &Service{
		cfg:  cfg,
		repo: repo,
		st:   st,
	}
}

// SetStrategy func
func (s *Service) SetStrategy(st strategy) {
	s.st = st
}

// RunBackup start backup with config
func (s *Service) RunBackup() (err error) {

	err = s.st.CheckExist(s.cfg.OriginalPath, ".git")
	if err != nil {
		return
	}

	backupedPath, err := s.st.Save(s.cfg.OriginalPath, s.cfg.BackupPath)
	if err != nil {
		if errors.Is(err, models.E_NOTHING_TO_PUSH) {
			err = nil
			return
		}
		return
	}

	deletedPath, err := s.repo.Save(backupedPath)
	if err != nil {
		return
	}

	if deletedPath != "" {
		err = s.st.DeletebackupPath(deletedPath)
	}

	return
}
