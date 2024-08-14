package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/olegtemek/t-backup/internal/models"
)

// Service struct
type Service struct{}

// New return service
func New() *Service {
	return &Service{}
}

// CheckExist func
func (s *Service) CheckExist(originalPath string, backupPath string) (err error) {

	gitPath := filepath.Join(originalPath, backupPath)
	if _, err = os.Stat(gitPath); err != nil {
		err = fmt.Errorf("cannot get %s path, ERROR: %s", backupPath, err)
		return
	}

	return
}

// Save func
func (s *Service) Save(originalPath string, _ string) (backupedPath string, err error) {

	if count, errM := s.checkModifyedFiles(originalPath); errM != nil || count <= 0 {
		err = models.E_NOTHING_TO_PUSH
		return
	}

	if err = s.runGitCommand(originalPath, "pull"); err != nil {
		err = fmt.Errorf("error exec git pull: %s", err)
		return
	}

	if err = s.runGitCommand(originalPath, "add", "."); err != nil {
		err = fmt.Errorf("error exec git add: %s", err)
		return
	}

	backupedPath = fmt.Sprintf("t-backup: %s", time.Now().Format("2006-01-02--15:04:05"))
	if err = s.runGitCommand(originalPath, "commit", "-m", backupedPath); err != nil {
		err = fmt.Errorf("error exec git commit: %s", err)
		return
	}

	if err = s.runGitCommand(originalPath, "push"); err != nil {
		err = fmt.Errorf("error exec git push: %s", err)
		return
	}

	return
}

func (s *Service) runGitCommand(path string, args ...string) (err error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = path
	return cmd.Run()
}

func (s *Service) checkModifyedFiles(path string) (modifiedFilesCount int, err error) {
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = path

	output, err := cmd.Output()

	if err != nil {
		return
	}
	lines := strings.Split(string(output), "\n")

	for _, line := range lines {

		if line == "" {
			continue
		}

		status := line[:2]

		if status != "" {
			modifiedFilesCount++
		}
	}

	return
}

// DeletebackupPath func
func (s *Service) DeletebackupPath(deletedPath string) (err error) {
	return
}
