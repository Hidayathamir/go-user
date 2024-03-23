package grpc

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/dto"
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/Hidayathamir/go-user/pkg/gousergrpc"
)

// Auth is controller GRPC for authentication related.
type Auth struct {
	gousergrpc.UnimplementedAuthServer

	cfg         config.Config
	usecaseAuth usecase.IAuth
}

var _ gousergrpc.AuthServer = &Auth{}

func newAuth(cfg config.Config, usecaseAuth usecase.IAuth) *Auth {
	return &Auth{
		cfg:         cfg,
		usecaseAuth: usecaseAuth,
	}
}

// LoginUser implements gousergrpc.AuthServer.
func (a *Auth) LoginUser(c context.Context, r *gousergrpc.ReqLoginUser) (*gousergrpc.ResLoginUser, error) {
	req := dto.ReqLoginUser{
		Username: r.GetUsername(),
		Password: r.GetPassword(),
	}

	resLoginUser, err := a.usecaseAuth.LoginUser(c, req)
	if err != nil {
		err := fmt.Errorf("Auth.usecaseAuth.LoginUser: %w", err)
		return nil, err
	}

	res := &gousergrpc.ResLoginUser{
		UserJwt: resLoginUser.UserJWT,
	}

	return res, nil
}

// RegisterUser implements gousergrpc.AuthServer.
func (a *Auth) RegisterUser(c context.Context, r *gousergrpc.ReqRegisterUser) (*gousergrpc.ResRegisterUser, error) {
	req := dto.ReqRegisterUser{
		Username: r.GetUsername(),
		Password: r.GetPassword(),
	}

	resRegisterUser, err := a.usecaseAuth.RegisterUser(c, req)
	if err != nil {
		err := fmt.Errorf("Auth.usecaseAuth.RegisterUser: %w", err)
		return nil, err
	}

	res := gousergrpc.ResRegisterUser{
		UserId: resRegisterUser.UserID,
	}

	return &res, nil
}
