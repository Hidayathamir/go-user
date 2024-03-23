package gouserhttp

import (
	"context"
)

// IProfileClient -.
type IProfileClient interface {
	GetProfileByUsername(ctx context.Context, req ReqGetProfileByUsername) (ResGetProfileByUsername, error)
	UpdateProfileByUserID(ctx context.Context, req ReqUpdateProfileByUserID) error
}

// ProfileClient -.
type ProfileClient struct{}

var _ IProfileClient = &ProfileClient{}

// NewProfileClient -.
func NewProfileClient() *ProfileClient {
	return &ProfileClient{}
}

// GetProfileByUsername implements IProfileClient.
func (p *ProfileClient) GetProfileByUsername(context.Context, ReqGetProfileByUsername) (ResGetProfileByUsername, error) {
	panic("unimplemented") // TODO: IMPLEMENT
}

// UpdateProfileByUserID implements IProfileClient.
func (p *ProfileClient) UpdateProfileByUserID(context.Context, ReqUpdateProfileByUserID) error {
	panic("unimplemented") // TODO: IMPLEMENT
}
