package http

import (
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/Hidayathamir/go-user/internal/usecase/repo"
)

// This file contains all dependency injection from layer controller to usecase
// to repo.

func newControllerAuth() *Auth {
	repoAuth := repo.NewAuth()
	usecaseAuth := usecase.NewAuth(repoAuth)
	controllerAuth := newAuth(usecaseAuth)
	return controllerAuth
}

func newControllerProfile() *Profile {
	repoProfile := repo.NewProfile()
	usecaseProfile := usecase.NewProfile(repoProfile)
	controllerProfile := newProfile(usecaseProfile)
	return controllerProfile
}
