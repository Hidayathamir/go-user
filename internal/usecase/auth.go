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
	// LoginUser validate username and password, return jwt string and error.
	LoginUser(ctx context.Context, req dto.ReqLoginUser) (string, error)
}

// Auth implement IAuth.
type Auth struct {
	repoAuth    repo.IAuth
	repoProfile repo.IProfile
}

var _ IAuth = &Auth{}

// NewAuth return *Auth which implement IAuth.
func NewAuth(repoAuth repo.IAuth, repoProfile repo.IProfile) *Auth {
	return &Auth{
		repoAuth:    repoAuth,
		repoProfile: repoProfile,
	}
}

// LoginUser validate username and password, return jwt string and error.
func (a *Auth) LoginUser(ctx context.Context, req dto.ReqLoginUser) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", fmt.Errorf("dto.ReqLoginUser.Validate: %w", err)
	}

	user, err := a.repoProfile.GetProfileByUsername(ctx, req.Username)
	if err != nil {
		return "", fmt.Errorf("Auth.repoProfile.GetProfileByUsername: %w", err)
	}

	err = auth.CompareHashAndPassword(user.Password, req.Password)
	if err != nil {
		return "", fmt.Errorf("auth.CompareHashAndPassword: %w", err)
	}

	userJWT := auth.GenerateUserJWTToken(user.ID)

	return userJWT, nil
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
