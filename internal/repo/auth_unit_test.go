package repo

import (
	"context"
	"testing"
	"time"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/db"
	"github.com/Hidayathamir/go-user/internal/entity"
	"github.com/Hidayathamir/go-user/internal/entity/table"
	"github.com/Hidayathamir/go-user/pkg/gouser"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitAuthRegisterUser(t *testing.T) {
	t.Parallel()

	t.Run("register user success", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		a := &Auth{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		mockpool.
			ExpectQuery("INSERT").WithArgs("hidayat", "mypassword", anyTime{}, anyTime{}).
			WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(int64(334)))

		userID, err := a.RegisterUser(context.Background(), entity.User{
			ID:        0,
			Username:  "hidayat",
			Password:  "mypassword",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		})

		assert.Equal(t, int64(334), userID)
		require.NoError(t, err)
	})
	t.Run("QueryRow Scan error should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		a := &Auth{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		mockpool.
			ExpectQuery("INSERT").WithArgs("hidayat", "mypassword", anyTime{}, anyTime{}).
			WillReturnError(assert.AnError)

		userID, err := a.RegisterUser(context.Background(), entity.User{
			ID:        0,
			Username:  "hidayat",
			Password:  "mypassword",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		})

		assert.Equal(t, int64(0), userID)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("QueryRow Scan duplicate username error should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		a := &Auth{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		mockpool.
			ExpectQuery("INSERT").WithArgs("hidayat", "mypassword", anyTime{}, anyTime{}).
			WillReturnError(
				&pgconn.PgError{Code: pgerrcode.UniqueViolation, ConstraintName: table.User.Constraint.UserUn},
			)

		userID, err := a.RegisterUser(context.Background(), entity.User{
			ID:        0,
			Username:  "hidayat",
			Password:  "mypassword",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		})

		assert.Equal(t, int64(0), userID)
		require.Error(t, err)
		require.ErrorIs(t, err, gouser.ErrDuplicateUsername)
	})
}
