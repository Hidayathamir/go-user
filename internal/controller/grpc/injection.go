package grpc

import (
	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/repo"
	"github.com/Hidayathamir/go-user/internal/repo/db"
	"github.com/Hidayathamir/go-user/internal/usecase"
)

func injectionAuth(cfg config.Config, db *db.Postgres) *Auth {
	repoAuth := repo.NewAuth(cfg, db)
	repoProfile := repo.NewProfile(cfg, db)
	usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
	controllerAuth := newAuth(cfg, usecaseAuth)
	return controllerAuth
}

func injectionProfile(cfg config.Config, db *db.Postgres) *Profile {
	repoProfile := repo.NewProfile(cfg, db)
	usecaseProfile := usecase.NewProfile(cfg, repoProfile)
	controllerProfile := newProfile(cfg, usecaseProfile)
	return controllerProfile
}
