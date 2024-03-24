package gouserhttp

import (
	"context"

	"github.com/Hidayathamir/go-user/internal/usecase"
)

// IProfileClient -.
type IProfileClient interface {
	GetProfileByUsername(ctx context.Context, req usecase.ReqGetProfileByUsername) (usecase.ResGetProfileByUsername, error)
	UpdateProfileByUserID(ctx context.Context, req usecase.ReqUpdateProfileByUserID) error
}

// ProfileClient -.
type ProfileClient struct{}

var _ IProfileClient = &ProfileClient{}

// NewProfileClient -.
func NewProfileClient() *ProfileClient {
	return &ProfileClient{}
}

// GetProfileByUsername implements IProfileClient.
func (p *ProfileClient) GetProfileByUsername(context.Context, usecase.ReqGetProfileByUsername) (usecase.ResGetProfileByUsername, error) {
	panic("unimplemented") // TODO: IMPLEMENT
}

// UpdateProfileByUserID implements IProfileClient.
func (p *ProfileClient) UpdateProfileByUserID(context.Context, usecase.ReqUpdateProfileByUserID) error {
	panic("unimplemented") // TODO: IMPLEMENT
}
