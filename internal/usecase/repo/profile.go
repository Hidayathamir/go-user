package repo

import "github.com/Hidayathamir/go-user/internal/usecase/repo/db"

// IProfile contains abstraction of repo profile.
type IProfile interface {
}

// Profile implement IProfile.
type Profile struct {
	db *db.Postgres
}

var _ IProfile = &Profile{}

// NewProfile return *Profile which implement repo.IProfile.
func NewProfile(db *db.Postgres) *Profile {
	return &Profile{
		db: db,
	}
}
