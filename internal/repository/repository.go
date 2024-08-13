package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/olegtemek/t-backup/internal/config"
	"github.com/olegtemek/t-backup/internal/models"
)

// Repository struct
type Repository struct {
	cfg     *config.Config
	mutex   sync.Mutex
	file    string
	backups []models.Backup
}

// New return repository
func New(cfg *config.Config) *Repository {
	file := "./backups.json"
	repo := &Repository{
		cfg:  cfg,
		file: file,
	}

	// Initialize backups from the JSON file
	if _, err := os.Stat(file); err == nil {
		data, err := os.ReadFile(file)
		if err != nil {
			panic(fmt.Errorf("cannot read backups file, ERROR: %s", err))
		}
		err = json.Unmarshal(data, &repo.backups)
		if err != nil {
			panic(fmt.Errorf("cannot unmarshal backups data, ERROR: %s", err))
		}
	} else if os.IsNotExist(err) {
		repo.backups = []models.Backup{}
	} else {
		panic(fmt.Errorf("cannot stat backups file, ERROR: %s", err))
	}

	return repo
}

func (r *Repository) saveToFile() (err error) {
	data, err := json.MarshalIndent(r.backups, "", "  ")
	if err != nil {
		err = fmt.Errorf("cannot marshal backups data, ERROR: %s", err)
		return
	}
	err = os.WriteFile(r.file, data, 0644)
	if err != nil {
		err = fmt.Errorf("cannot write backups file, ERROR: %s", err)
		return
	}
	return
}


// Save for make backup in backups.json
func (r *Repository) Save(backupedPath string) (deletedPath string, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Count backups with the same name
	var count int
	for _, backup := range r.backups {
		if backup.Name == r.cfg.Name {
			count++
		}
	}

	// Delete the oldest backup if the count exceeds the limit
	if count >= r.cfg.NumberOfBackups {
		for i, backup := range r.backups {
			if backup.Name == r.cfg.Name {
				deletedPath = backup.BackupPath
				r.backups = append(r.backups[:i], r.backups[i+1:]...)
				break
			}
		}
	}

	// Add the new backup
	newBackup := models.Backup{
		Name:       r.cfg.Name,
		Driver:     r.cfg.Driver,
		BackupPath: backupedPath,
	}
	r.backups = append(r.backups, newBackup)

	err = r.saveToFile()
	if err != nil {
		return
	}

	return
}
