package usecase

import "github.com/Hidayathamir/go-user/internal/usecase/repo"

// IAuth contains abstraction of usecase authentication.
type IAuth interface {
}

// Auth implement IAuth.
type Auth struct {
	repoAuth repo.IAuth
}

var _ IAuth = &Auth{}

func newAuth(repoAuth repo.IAuth) *Auth {
	return &Auth{
		repoAuth: repoAuth,
	}
}
