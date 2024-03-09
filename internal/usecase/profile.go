package usecase

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-user/internal/dto"
	"github.com/Hidayathamir/go-user/internal/usecase/repo"
)

// IProfile contains abstraction of usecase profile.
type IProfile interface {
	// GetProfileByUsername return user profile by username.
	GetProfileByUsername(ctx context.Context, username string) (dto.ResGetProfileByUsername, error)
	// UpdateProfileByUsername update user profile by username.
	UpdateProfileByUsername(ctx context.Context, req dto.ReqUpdateProfileByUsername) error
}

// Profile implement IProfile.
type Profile struct {
	repoProfile repo.IProfile
}

var _ IProfile = &Profile{}

func newProfile(repoProfile repo.IProfile) *Profile {
	return &Profile{
		repoProfile: repoProfile,
	}
}

// GetProfileByUsername return user profile by username.
func (p *Profile) GetProfileByUsername(ctx context.Context, username string) (dto.ResGetProfileByUsername, error) {
	user, err := p.repoProfile.GetProfileByUsername(ctx, username)
	if err != nil {
		return dto.ResGetProfileByUsername{}, fmt.Errorf("Profile.repoProfile.GetProfileByUsername: %w", err)
	}

	res := dto.ResGetProfileByUsername{}
	res = res.LoadEntityUser(user)

	return res, nil
}

// UpdateProfileByUsername update user profile by username.
func (p *Profile) UpdateProfileByUsername(ctx context.Context, req dto.ReqUpdateProfileByUsername) error {
	err := req.Validate()
	if err != nil {
		return fmt.Errorf("dto.ReqUpdateProfileByUsername.Validate: %w", err)
	}

	err = p.repoProfile.UpdateProfileByUsername(ctx, req.ToEntityUser())
	if err != nil {
		return fmt.Errorf("Profile.repoProfile.UpdateProfileByUsername: %w", err)
	}

	return nil
}
