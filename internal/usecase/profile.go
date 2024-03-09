package usecase

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-user/internal/entity"
	"github.com/Hidayathamir/go-user/internal/usecase/repo"
)

// IProfile contains abstraction of usecase profile.
type IProfile interface {
	// GetProfileByUsername return user profile by username.
	GetProfileByUsername(ctx context.Context, username string) (entity.User, error)
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
func (p *Profile) GetProfileByUsername(ctx context.Context, username string) (entity.User, error) {
	user, err := p.repoProfile.GetProfileByUsername(ctx, username)
	if err != nil {
		return entity.User{}, fmt.Errorf("Profile.repoProfile.GetProfileByUsername: %w", err)
	}
	return user, nil
}
