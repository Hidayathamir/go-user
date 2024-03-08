package usecase

import "github.com/Hidayathamir/go-user/internal/usecase/repo"

// IProfile contains abstraction of usecase profile.
type IProfile interface {
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
