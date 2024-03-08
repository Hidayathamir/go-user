package usecase

import (
	"github.com/Hidayathamir/go-user/internal/usecase/repo"
	"github.com/Hidayathamir/go-user/internal/usecase/repo/db"
)

// This file contains all dependency injection from layer usecase to repo.

// NewAuth return usecase *Auth which implement usecase.IAuth.
func NewAuth(db *db.Postgres) *Auth {
	repoAuth := repo.NewAuth(db)
	usecaseAuth := newAuth(repoAuth)
	return usecaseAuth
}

// NewProfile return usecase *Profile which implement usecase.IProfile.
func NewProfile(db *db.Postgres) *Profile {
	repoProfile := repo.NewProfile(db)
	usecaseProfile := newProfile(repoProfile)
	return usecaseProfile
}
