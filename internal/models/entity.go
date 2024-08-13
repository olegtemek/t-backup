package models

// Backup struct
type Backup struct {
	Name       string `json:"name"`
	Driver     string `json:"driver"`
	BackupPath string `json:"backupPath"`
}
