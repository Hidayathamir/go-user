package usecase

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/dto"
	"github.com/Hidayathamir/go-user/internal/pkg/auth"
	"github.com/Hidayathamir/go-user/internal/repo"
	"github.com/Hidayathamir/go-user/pkg/gouser"
)

//go:generate mockgen -source=profile.go -destination=mockusecase/profile.go -package=mockusecase

// IProfile contains abstraction of usecase profile.
type IProfile interface {
	// GetProfileByUsername return user profile by username.
	GetProfileByUsername(ctx context.Context, req dto.ReqGetProfileByUsername) (dto.ResGetProfileByUsername, error)
	// UpdateProfileByUserID update user profile by user id.
	UpdateProfileByUserID(ctx context.Context, req dto.ReqUpdateProfileByUserID) error
}

// Profile implement IProfile.
type Profile struct {
	cfg         config.Config
	repoProfile repo.IProfile
}

var _ IProfile = &Profile{}

// NewProfile return *Profile which implement IProfile.
func NewProfile(cfg config.Config, repoProfile repo.IProfile) *Profile {
	return &Profile{
		cfg:         cfg,
		repoProfile: repoProfile,
	}
}

// GetProfileByUsername return user profile by username.
func (p *Profile) GetProfileByUsername(ctx context.Context, req dto.ReqGetProfileByUsername) (dto.ResGetProfileByUsername, error) {
	err := req.Validate()
	if err != nil {
		err := fmt.Errorf("dto.ReqGetProfileByUsername.Validate: %w", err)
		return dto.ResGetProfileByUsername{}, fmt.Errorf("%w: %w", gouser.ErrRequestInvalid, err)
	}

	user, err := p.repoProfile.GetProfileByUsername(ctx, req.Username)
	if err != nil {
		return dto.ResGetProfileByUsername{}, fmt.Errorf("Profile.repoProfile.GetProfileByUsername: %w", err)
	}

	res := dto.ResGetProfileByUsername{}
	res = res.LoadEntityUser(user)

	return res, nil
}

// UpdateProfileByUserID update user profile by user id.
func (p *Profile) UpdateProfileByUserID(ctx context.Context, req dto.ReqUpdateProfileByUserID) error {
	err := req.Validate()
	if err != nil {
		err := fmt.Errorf("dto.ReqUpdateProfileByUserID.Validate: %w", err)
		return fmt.Errorf("%w: %w", gouser.ErrRequestInvalid, err)
	}

	userID, err := auth.GetUserIDFromJWTTokenString(p.cfg, req.UserJWT)
	if err != nil {
		return fmt.Errorf("auth.GetUserIDFromJWTTokenString: %w", err)
	}

	user := req.ToEntityUser()
	user.ID = userID

	if user.Password != "" {
		user.Password, err = auth.GenerateHashPassword(user.Password)
		if err != nil {
			return fmt.Errorf("auth.GenerateHashPassword: %w", err)
		}
	}

	err = p.repoProfile.UpdateProfileByUserID(ctx, user)
	if err != nil {
		return fmt.Errorf("Profile.repoProfile.UpdateProfileByUserID: %w", err)
	}

	return nil
}
