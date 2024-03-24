package usecase

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/pkg/auth"
	"github.com/Hidayathamir/go-user/internal/repo"
	"github.com/Hidayathamir/go-user/pkg/gouser"
)

//go:generate mockgen -source=auth.go -destination=mockusecase/auth.go -package=mockusecase

// IAuth contains abstraction of usecase authentication.
type IAuth interface {
	// RegisterUser register new user.
	RegisterUser(ctx context.Context, req gouser.ReqRegisterUser) (gouser.ResRegisterUser, error)
	// LoginUser validate username and password.
	LoginUser(ctx context.Context, req gouser.ReqLoginUser) (gouser.ResLoginUser, error)
}

// Auth implement IAuth.
type Auth struct {
	cfg         config.Config
	repoAuth    repo.IAuth
	repoProfile repo.IProfile
}

var _ IAuth = &Auth{}

// NewAuth return *Auth which implement IAuth.
func NewAuth(cfg config.Config, repoAuth repo.IAuth, repoProfile repo.IProfile) *Auth {
	return &Auth{
		cfg:         cfg,
		repoAuth:    repoAuth,
		repoProfile: repoProfile,
	}
}

// LoginUser validate username and password.
func (a *Auth) LoginUser(ctx context.Context, req gouser.ReqLoginUser) (gouser.ResLoginUser, error) {
	err := req.Validate()
	if err != nil {
		err := fmt.Errorf("ReqLoginUser.Validate: %w", err)
		return gouser.ResLoginUser{}, fmt.Errorf("%w: %w", gouser.ErrRequestInvalid, err)
	}

	user, err := a.repoProfile.GetProfileByUsername(ctx, req.Username)
	if err != nil {
		return gouser.ResLoginUser{}, fmt.Errorf("Auth.repoProfile.GetProfileByUsername: %w", err)
	}

	err = auth.CompareHashAndPassword(user.Password, req.Password)
	if err != nil {
		err := fmt.Errorf("auth.CompareHashAndPassword: %w", err)
		return gouser.ResLoginUser{}, fmt.Errorf("%w: %w", gouser.ErrWrongPassword, err)
	}

	userJWT := auth.GenerateUserJWTToken(user.ID, a.cfg)

	res := gouser.ResLoginUser{
		UserJWT: userJWT,
	}

	return res, nil
}

// RegisterUser register new user.
func (a *Auth) RegisterUser(ctx context.Context, req gouser.ReqRegisterUser) (gouser.ResRegisterUser, error) {
	err := req.Validate()
	if err != nil {
		err := fmt.Errorf("ReqRegisterUser.Validate: %w", err)
		return gouser.ResRegisterUser{}, fmt.Errorf("%w: %w", gouser.ErrRequestInvalid, err)
	}

	user := req.ToEntityUser()
	user.Password, err = auth.GenerateHashPassword(user.Password)
	if err != nil {
		return gouser.ResRegisterUser{}, fmt.Errorf("auth.GenerateHashPassword: %w", err)
	}

	userID, err := a.repoAuth.RegisterUser(ctx, user)
	if err != nil {
		return gouser.ResRegisterUser{}, fmt.Errorf("Auth.repoAuth.RegisterUser: %w", err)
	}

	res := gouser.ResRegisterUser{
		UserID: userID,
	}

	return res, nil
}
