package config

import (
	"flag"
)

// Config struct
type Config struct {
	Name            string
	OriginalPath    string
	BackupPath      string
	Driver          string
	NumberOfBackups int
}


// New return config 
func New() *Config {

	name := flag.String("name", "", "backup name")
	originalPath := flag.String("originalPath", "", "original folder path")
	backupPath := flag.String("backupPath", "", "backup path")
	driver := flag.String("driver", "", "driver, how to backup")
	numberOfBackups := flag.Int("numberOfBackups", 0, "number of backups")
	flag.Parse()

	if *name == "" || *originalPath == "" || *driver == "" {
		panic("cannot parse config")
	}

	return &Config{
		Name:            *name,
		OriginalPath:    *originalPath,
		BackupPath:      *backupPath,
		Driver:          *driver,
		NumberOfBackups: *numberOfBackups,
	}
}
