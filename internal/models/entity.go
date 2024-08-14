package models

import "fmt"

// Backup struct
type Backup struct {
	Name       string `json:"name"`
	Driver     string `json:"driver"`
	BackupPath string `json:"backupPath"`
}

var (
	E_NOTHING_TO_PUSH = fmt.Errorf("nothing to push")
)
