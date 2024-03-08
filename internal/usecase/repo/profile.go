package repo

// IProfile contains abstraction of repo profile.
type IProfile interface {
}

// Profile implement IProfile.
type Profile struct {
}

var _ IProfile = &Profile{}

// NewProfile -.
func NewProfile() *Profile {
	return &Profile{}
}
