package repo

import "github.com/Hidayathamir/go-user/internal/usecase/repo/db"

// IAuth contains abstraction of repo authentication.
type IAuth interface {
}

// Auth implement IAuth.
type Auth struct {
	db *db.Postgres
}

var _ IAuth = &Auth{}

// NewAuth return *Auth which implement repo.IAuth.
func NewAuth(db *db.Postgres) *Auth {
	return &Auth{
		db: db,
	}
}
