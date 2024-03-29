package grpc

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/Hidayathamir/go-user/pkg/gouser"
	"github.com/Hidayathamir/go-user/pkg/gousergrpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Profile is controller GRPC for profile related.
type Profile struct {
	gousergrpc.UnimplementedProfileServer

	cfg            config.Config
	usecaseProfile usecase.IProfile
}

var _ gousergrpc.ProfileServer = &Profile{}

func newProfile(cfg config.Config, usecaseProfile usecase.IProfile) *Profile {
	return &Profile{
		cfg:            cfg,
		usecaseProfile: usecaseProfile,
	}
}

// GetProfileByUsername implements gousergrpc.ProfileServer.
func (p *Profile) GetProfileByUsername(c context.Context, r *gousergrpc.ReqGetProfileByUsername) (*gousergrpc.ResGetProfileByUsername, error) {
	req := gouser.ReqGetProfileByUsername{Username: r.GetUsername()}

	user, err := p.usecaseProfile.GetProfileByUsername(c, req)
	if err != nil {
		err := fmt.Errorf("Profile.usecaseProfile.GetProfileByUsername: %w", err)
		return nil, err
	}

	res := &gousergrpc.ResGetProfileByUsername{
		Id:        user.ID,
		Username:  user.Username,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}

	return res, nil
}

// UpdateProfileByUserID implements gousergrpc.ProfileServer.
func (p *Profile) UpdateProfileByUserID(c context.Context, r *gousergrpc.ReqUpdateProfileByUserID) (*gousergrpc.ProfileEmpty, error) {
	req := gouser.ReqUpdateProfileByUserID{
		UserJWT:  r.GetUserJwt(),
		Password: r.GetPassword(),
	}

	err := p.usecaseProfile.UpdateProfileByUserID(c, req)
	if err != nil {
		err := fmt.Errorf("Profile.usecaseProfile.UpdateProfileByUserID: %w", err)
		return nil, err
	}

	res := &gousergrpc.ProfileEmpty{}

	return res, nil
}
