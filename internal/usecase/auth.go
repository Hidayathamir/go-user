package usecase

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-user/internal/dto"
	"github.com/Hidayathamir/go-user/internal/usecase/repo"
	"github.com/Hidayathamir/go-user/pkg/auth"
)

// IAuth contains abstraction of usecase authentication.
type IAuth interface {
	// RegisterUser register new user.
	RegisterUser(ctx context.Context, req dto.ReqRegisterUser) (int64, error)
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

// RegisterUser register new user.
func (a *Auth) RegisterUser(ctx context.Context, req dto.ReqRegisterUser) (int64, error) {
	err := req.Validate()
	if err != nil {
		return 0, fmt.Errorf("dto.ReqRegisterUser.Validate: %w", err)
	}

	user := req.ToEntityUser()
	user.Password, err = auth.GenerateHashPassword(user.Password)
	if err != nil {
		return 0, fmt.Errorf("auth.GenerateHashPassword: %w", err)
	}

	userID, err := a.repoAuth.RegisterUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("Auth.repoAuth.RegisterUser: %w", err)
	}

	return userID, nil
}
