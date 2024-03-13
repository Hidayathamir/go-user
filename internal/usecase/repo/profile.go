package repo

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/entity"
	"github.com/Hidayathamir/go-user/internal/entity/table"
	"github.com/Hidayathamir/go-user/internal/usecase/repo/db"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

// IProfile contains abstraction of repo profile.
type IProfile interface {
	// GetProfileByUsername return user profile by username.
	GetProfileByUsername(ctx context.Context, username string) (entity.User, error)
	// UpdateProfileByUsername update user profile by username.
	UpdateProfileByUsername(ctx context.Context, user entity.User) error
}

// Profile implement IProfile.
type Profile struct {
	cfg config.Config
	db  *db.Postgres
}

var _ IProfile = &Profile{}

// NewProfile return *Profile which implement repo.IProfile.
func NewProfile(cfg config.Config, db *db.Postgres) *Profile {
	return &Profile{
		cfg: cfg,
		db:  db,
	}
}

// GetProfileByUsername return user profile by username.
func (p *Profile) GetProfileByUsername(ctx context.Context, username string) (entity.User, error) {
	sql, args, err := p.db.Builder.
		Select(
			table.User.ID, table.User.Username, table.User.Password,
			table.User.CreatedAt, table.User.UpdatedAt,
		).
		From(table.User.String()).
		Where(sq.Eq{
			table.User.Username: username,
		}).
		ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("Profile.db.Builder.ToSql: %w", err)
	}

	user := entity.User{}
	err = p.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.ID, &user.Username, &user.Password,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return entity.User{}, fmt.Errorf("pgx.Rows.Scan: %w", err)
	}

	return user, nil
}

// UpdateProfileByUsername update user profile by username.
func (p *Profile) UpdateProfileByUsername(ctx context.Context, user entity.User) error {
	set := sq.Eq{}

	if user.Password != "" {
		set[table.User.Password] = user.Password
	}

	sql, args, err := p.db.Builder.
		Update(table.User.String()).
		SetMap(set).
		Where(sq.Eq{
			table.User.Username: user.Username,
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("Profile.db.Builder.ToSql: %w", err)
	}

	commandTag, err := p.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("Profile.db.Pool.Exec: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("pgconn.CommandTag.RowsAffected == 0: %w", pgx.ErrNoRows)
	}

	return nil
}
