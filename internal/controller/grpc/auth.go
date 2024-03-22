package grpc

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/dto"
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/Hidayathamir/go-user/pkg/gouser/grpc/pb"
)

// Auth is controller GRPC for authentication related.
type Auth struct {
	pb.UnimplementedAuthServer

	cfg         config.Config
	usecaseAuth usecase.IAuth
}

var _ pb.AuthServer = &Auth{}

func newAuth(cfg config.Config, usecaseAuth usecase.IAuth) *Auth {
	return &Auth{
		cfg:         cfg,
		usecaseAuth: usecaseAuth,
	}
}

// LoginUser implements pb.AuthServer.
func (a *Auth) LoginUser(c context.Context, r *pb.ReqLoginUser) (*pb.ResLoginUser, error) {
	req := dto.ReqLoginUser{
		Username: r.GetUsername(),
		Password: r.GetPassword(),
	}

	userJWT, err := a.usecaseAuth.LoginUser(c, req)
	if err != nil {
		err := fmt.Errorf("Auth.usecaseAuth.LoginUser: %w", err)
		return nil, err
	}

	res := &pb.ResLoginUser{
		UserJwt: userJWT,
	}

	return res, nil
}

// RegisterUser implements pb.AuthServer.
func (a *Auth) RegisterUser(c context.Context, r *pb.ReqRegisterUser) (*pb.ResRegisterUser, error) {
	req := dto.ReqRegisterUser{
		Username: r.GetUsername(),
		Password: r.GetPassword(),
	}

	userID, err := a.usecaseAuth.RegisterUser(c, req)
	if err != nil {
		err := fmt.Errorf("Auth.usecaseAuth.RegisterUser: %w", err)
		return nil, err
	}

	res := pb.ResRegisterUser{
		UserId: userID,
	}

	return &res, nil
}
