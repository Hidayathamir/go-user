package grpc

import (
	"context"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/Hidayathamir/go-user/pkg/gouser/grpc/pb"
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
func (p *Profile) GetProfileByUsername(context.Context, *pb.ReqGetProfileByUsername) (*pb.ResGetProfileByUsername, error) {
	panic("unimplemented")
}

// UpdateProfileByUserID implements pb.ProfileServer.
func (p *Profile) UpdateProfileByUserID(context.Context, *pb.ReqUpdateProfileByUserID) (*pb.ProfileEmpty, error) {
	panic("unimplemented")
}
