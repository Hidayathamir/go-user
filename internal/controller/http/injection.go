package http

import (
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/Hidayathamir/go-user/internal/usecase/repo"
	"github.com/Hidayathamir/go-user/internal/usecase/repo/db"
)

func injectionAuth(db *db.Postgres) *Auth {
	repoAuth := repo.NewAuth(db)
	repoProfile := repo.NewProfile(db)
	usecaseAuth := usecase.NewAuth(repoAuth, repoProfile)
	controllerAuth := newAuth(usecaseAuth)
	return controllerAuth
}

func injectionProfile(db *db.Postgres) *Profile {
	repoProfile := repo.NewProfile(db)
	usecaseProfile := usecase.NewProfile(repoProfile)
	controllerProfile := newProfile(usecaseProfile)
	return controllerProfile
}
