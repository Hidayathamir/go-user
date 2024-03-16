package repo

import (
	"context"
	"testing"
	"time"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/entity"
	"github.com/Hidayathamir/go-user/internal/usecase/repo/db"
	"github.com/Hidayathamir/go-user/pkg/gouser"
	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitProfileGetProfileByUsername(t *testing.T) {
	t.Parallel()

	t.Run("get profile by username success", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		p := &Profile{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		now := time.Now()
		mockpool.ExpectQuery("SELECT").WithArgs("hidayat").
			WillReturnRows(
				pgxmock.NewRows(
					[]string{"id", "username", "password", "created_at", "updated_at"},
				).AddRow(
					int64(441), "hidayat", "dummyhashedpassword", now, now,
				),
			)

		user, err := p.GetProfileByUsername(context.Background(), "hidayat")

		require.NoError(t, err)
		assert.NotEmpty(t, user)
		assert.Equal(t, "hidayat", user.Username)
		assert.Equal(t, "dummyhashedpassword", user.Password)
		assert.Equal(t, now, user.CreatedAt)
		assert.Equal(t, now, user.UpdatedAt)
	})
	t.Run("QueryRow Scan error should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		p := &Profile{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		mockpool.ExpectQuery("SELECT").WithArgs("hidayat").
			WillReturnError(assert.AnError)

		user, err := p.GetProfileByUsername(context.Background(), "hidayat")

		require.Error(t, err)
		assert.Empty(t, user)
	})
	t.Run("QueryRow Scan no row error should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		p := &Profile{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		mockpool.ExpectQuery("SELECT").WithArgs("hidayat").
			WillReturnError(pgx.ErrNoRows)

		user, err := p.GetProfileByUsername(context.Background(), "hidayat")

		require.Error(t, err)
		require.ErrorIs(t, err, gouser.ErrUnknownUsername)
		assert.Empty(t, user)
	})
}

func TestUnitProfileUpdateProfileByUserID(t *testing.T) {
	t.Parallel()

	t.Run("update profile success", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		p := &Profile{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		mockpool.ExpectExec("UPDATE").WithArgs("newpassword", int64(776)).
			WillReturnResult(pgxmock.NewResult("UPDATE", 1))

		err = p.UpdateProfileByUserID(context.Background(), entity.User{
			ID:       776,
			Password: "newpassword",
		})

		require.NoError(t, err)
	})
	t.Run("Exec error should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		p := &Profile{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		mockpool.ExpectExec("UPDATE").WithArgs("newpassword", int64(776)).
			WillReturnError(assert.AnError)

		err = p.UpdateProfileByUserID(context.Background(), entity.User{
			ID:       776,
			Password: "newpassword",
		})

		require.Error(t, err)
	})
	t.Run("RowsAffected 0 should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		p := &Profile{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		mockpool.ExpectExec("UPDATE").WithArgs("newpassword", int64(776)).
			WillReturnResult(pgxmock.NewResult("UPDATE", 0))

		err = p.UpdateProfileByUserID(context.Background(), entity.User{
			ID:       776,
			Password: "newpassword",
		})

		require.Error(t, err)
		require.ErrorContains(t, err, "RowsAffected == 0")
	})
	t.Run("ToSql nothing to be update error should return error", func(t *testing.T) {
		t.Parallel()

		mockpool, err := pgxmock.NewPool(pgxmock.QueryMatcherOption(pgxmock.QueryMatcherRegexp))
		require.NoError(t, err)

		p := &Profile{
			cfg: config.Config{},
			db: &db.Postgres{
				Builder: builder,
				Pool:    mockpool,
			},
		}

		err = p.UpdateProfileByUserID(context.Background(), entity.User{
			ID:       776,
			Password: "",
		})

		require.Error(t, err)
		require.ErrorIs(t, err, gouser.ErrNothingToBeUpdate)
	})
}
