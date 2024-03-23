package gouserhttp

import (
	"context"
)

// IAuthClient -.
type IAuthClient interface {
	LoginUser(ctx context.Context, req ReqLoginUser) (ResLoginUser, error)
	RegisterUser(ctx context.Context, req ReqRegisterUser) (ResRegisterUser, error)
}

// AuthClient -.
type AuthClient struct{}

var _ IAuthClient = &AuthClient{}

// NewAuthClient -.
func NewAuthClient() *AuthClient {
	return &AuthClient{}
}

// LoginUser implements AuthClient.
func (a *AuthClient) LoginUser(context.Context, ReqLoginUser) (ResLoginUser, error) {
	panic("unimplemented") // TODO: IMPLEMENT
}

// RegisterUser implements AuthClient.
func (a *AuthClient) RegisterUser(context.Context, ReqRegisterUser) (ResRegisterUser, error) {
	panic("unimplemented") // TODO: IMPLEMENT
}
