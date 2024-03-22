package grpc

import (
	"context"

	"github.com/Hidayathamir/go-user/config"
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
func (a *Auth) LoginUser(context.Context, *pb.ReqLoginUser) (*pb.ResLoginUser, error) {
	panic("unimplemented")
}

// RegisterUser implements pb.AuthServer.
func (a *Auth) RegisterUser(context.Context, *pb.ReqRegisterUser) (*pb.ResRegisterUser, error) {
	panic("unimplemented")
}
