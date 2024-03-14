package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/entity"
	"github.com/Hidayathamir/go-user/internal/entity/table"
	"github.com/Hidayathamir/go-user/internal/pkg/query"
	"github.com/Hidayathamir/go-user/internal/usecase/repo/db"
	"github.com/Hidayathamir/go-user/pkg/gouser"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

// IAuth contains abstraction of repo authentication.
type IAuth interface {
	// RegisterUser register new user.
	RegisterUser(ctx context.Context, user entity.User) (int64, error)
}

// Auth implement IAuth.
type Auth struct {
	cfg config.Config
	db  *db.Postgres
}

var _ IAuth = &Auth{}

// NewAuth return *Auth which implement repo.IAuth.
func NewAuth(cfg config.Config, db *db.Postgres) *Auth {
	return &Auth{
		cfg: cfg,
		db:  db,
	}
}

// RegisterUser register new user.
func (a *Auth) RegisterUser(ctx context.Context, user entity.User) (int64, error) {
	now := time.Now()

	sql, args, err := a.db.Builder.
		Insert(table.User.String()).
		Columns(
			table.User.Username, table.User.Password,
			table.User.CreatedAt, table.User.UpdatedAt,
		).
		Values(
			user.Username, user.Password,
			now, now,
		).
		Suffix(query.Returning(table.User.ID)).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("Auth.db.Builder.ToSql: %w", err)
	}

	var userID int64
	err = a.db.Pool.QueryRow(ctx, sql, args...).Scan(&userID)
	if err != nil {
		err := fmt.Errorf("Auth.db.Pool.QueryRow.Scan: %w", err)

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			isErrDuplicateUsername := pgErr.Code == pgerrcode.UniqueViolation &&
				pgErr.ConstraintName == table.User.Constraint.UserUn
			if isErrDuplicateUsername {
				return 0, fmt.Errorf("%w: %w", gouser.ErrDuplicateUsername, err)
			}
		}

		return 0, err
	}

	return userID, nil
}
