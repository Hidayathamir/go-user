package grpc

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/dto"
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/Hidayathamir/go-user/pkg/gouser/grpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Profile is controller GRPC for profile related.
type Profile struct {
	pb.UnimplementedProfileServer

	cfg            config.Config
	usecaseProfile usecase.IProfile
}

var _ pb.ProfileServer = &Profile{}

func newProfile(cfg config.Config, usecaseProfile usecase.IProfile) *Profile {
	return &Profile{
		cfg:            cfg,
		usecaseProfile: usecaseProfile,
	}
}

// GetProfileByUsername implements pb.ProfileServer.
func (p *Profile) GetProfileByUsername(c context.Context, r *pb.ReqGetProfileByUsername) (*pb.ResGetProfileByUsername, error) {
	user, err := p.usecaseProfile.GetProfileByUsername(c, r.GetUsername())
	if err != nil {
		err := fmt.Errorf("Profile.usecaseProfile.GetProfileByUsername: %w", err)
		return nil, err
	}

	res := &pb.ResGetProfileByUsername{
		Id:        user.ID,
		Username:  user.Username,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}

	return res, nil
}

// UpdateProfileByUserID implements pb.ProfileServer.
func (p *Profile) UpdateProfileByUserID(c context.Context, r *pb.ReqUpdateProfileByUserID) (*pb.ProfileEmpty, error) {
	req := dto.ReqUpdateProfileByUserID{
		UserJWT:  r.GetUserJwt(),
		Password: r.GetPassword(),
	}

	err := p.usecaseProfile.UpdateProfileByUserID(c, req)
	if err != nil {
		err := fmt.Errorf("Profile.usecaseProfile.UpdateProfileByUserID: %w", err)
		return nil, err
	}

	res := &pb.ProfileEmpty{}

	return res, nil
}
