package repo

// IAuth contains abstraction of repo authentication.
type IAuth interface {
}

// Auth implement IAuth.
type Auth struct {
}

var _ IAuth = &Auth{}

// NewAuth -.
func NewAuth() *Auth {
	return &Auth{}
}
