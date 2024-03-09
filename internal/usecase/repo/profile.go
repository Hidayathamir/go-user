package repo

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-user/internal/entity"
	"github.com/Hidayathamir/go-user/internal/entity/table"
	"github.com/Hidayathamir/go-user/internal/usecase/repo/db"
	sq "github.com/Masterminds/squirrel"
)

// IProfile contains abstraction of repo profile.
type IProfile interface {
	// GetProfileByUsername return user profile by username.
	GetProfileByUsername(ctx context.Context, username string) (entity.User, error)
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

// GetProfileByUsername return user profile by username.
func (p *Profile) GetProfileByUsername(ctx context.Context, username string) (entity.User, error) {
	sql, _, err := p.db.Builder.
		Select(
			table.User.ID, table.User.Username, table.User.Password,
			table.User.CreatedAt, table.User.UpdatedAt,
		).
		From(table.User.String()).
		Where(sq.Eq{
			table.User.Username: "?",
		}).
		ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("Profile.db.Builder.ToSql: %w", err)
	}

	rows, err := p.db.Pool.Query(ctx, sql, username)
	if err != nil {
		return entity.User{}, fmt.Errorf("Profile.db.Pool.Query: %w", err)
	}
	defer rows.Close()

	user := entity.User{}
	for rows.Next() {
		err := rows.Scan(
			&user.ID, &user.Username, &user.Password,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return entity.User{}, fmt.Errorf("pgx.Rows.Scan: %w", err)
		}
	}

	return user, nil
}
