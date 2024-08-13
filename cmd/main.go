package main

import (
	"github.com/olegtemek/t-backup/internal/config"
	"github.com/olegtemek/t-backup/internal/repository"
	"github.com/olegtemek/t-backup/internal/service"
	"github.com/olegtemek/t-backup/internal/service/git"
	"github.com/olegtemek/t-backup/internal/service/local"
)

func main() {
	cfg := config.New()

	repo := repository.New(cfg)

	service := service.New(cfg, repo, nil)

	switch cfg.Driver {
	case "git":
		service.SetStrategy(git.New())
	case "local":
		service.SetStrategy(local.New())
	}

	err := service.RunBackup()
	if err != nil {
		panic(err)
	}
}
