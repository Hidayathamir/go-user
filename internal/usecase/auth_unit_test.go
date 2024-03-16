package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/dto"
	"github.com/Hidayathamir/go-user/internal/entity"
	"github.com/Hidayathamir/go-user/internal/pkg/auth"
	"github.com/Hidayathamir/go-user/internal/repo/mockrepo"
	"github.com/Hidayathamir/go-user/pkg/gouser"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUnitAuthLoginUser(t *testing.T) {
	t.Parallel()

	t.Run("login user with correct password should return success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoAuth := mockrepo.NewMockIAuth(ctrl)
		repoProfile := mockrepo.NewMockIProfile(ctrl)

		cfg := config.Config{
			JWT: config.JWT{ExpireHour: 24, SignedKey: "secretjwtkey"},
		}

		a := &Auth{
			cfg:         cfg,
			repoAuth:    repoAuth,
			repoProfile: repoProfile,
		}

		repoProfile.EXPECT().
			GetProfileByUsername(gomock.Any(), "hidayat").
			Return(entity.User{
				ID:        99,
				Username:  "hidayat",
				Password:  "$2a$10$KrDmeYfFUKWtTn9aS1ZrQ.L6WG0l0aQUStjxfOnm4U8gH9MqWrFKO", // hashed of "mypassword"
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			}, nil)

		userJWT, err := a.LoginUser(context.Background(), dto.ReqLoginUser{
			Username: "hidayat",
			Password: "mypassword",
		})

		require.NoError(t, err)
		assert.NotEmpty(t, userJWT)
		assert.Contains(t, userJWT, "Bearer ")
		userID, err := auth.GetUserIDFromJWTTokenString(cfg, userJWT)
		require.NoError(t, err)
		assert.Equal(t, int64(99), userID)
	})
	t.Run("login user with wrong password should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoAuth := mockrepo.NewMockIAuth(ctrl)
		repoProfile := mockrepo.NewMockIProfile(ctrl)

		cfg := config.Config{
			JWT: config.JWT{ExpireHour: 24, SignedKey: "secretjwtkey"},
		}

		a := &Auth{
			cfg:         cfg,
			repoAuth:    repoAuth,
			repoProfile: repoProfile,
		}

		repoProfile.EXPECT().
			GetProfileByUsername(gomock.Any(), "hidayat").
			Return(entity.User{
				ID:        99,
				Username:  "hidayat",
				Password:  "$2a$10$KrDmeYfFUKWtTn9aS1ZrQ.L6WG0l0aQUStjxfOnm4U8gH9MqWrFKO", // hashed of "mypassword"
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			}, nil)

		userJWT, err := a.LoginUser(context.Background(), dto.ReqLoginUser{
			Username: "hidayat",
			Password: "wrongpassword",
		})

		assert.Empty(t, userJWT)
		require.Error(t, err)
		require.ErrorIs(t, err, gouser.ErrWrongPassword)
	})
	t.Run("call repo GetProfileByUsername error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoAuth := mockrepo.NewMockIAuth(ctrl)
		repoProfile := mockrepo.NewMockIProfile(ctrl)

		cfg := config.Config{
			JWT: config.JWT{ExpireHour: 24, SignedKey: "secretjwtkey"},
		}

		a := &Auth{
			cfg:         cfg,
			repoAuth:    repoAuth,
			repoProfile: repoProfile,
		}

		repoProfile.EXPECT().
			GetProfileByUsername(gomock.Any(), "hidayat").
			Return(entity.User{}, assert.AnError)

		userJWT, err := a.LoginUser(context.Background(), dto.ReqLoginUser{
			Username: "hidayat",
			Password: "mypassword",
		})

		assert.Empty(t, userJWT)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("request validate error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		a := &Auth{
			cfg:         config.Config{},
			repoAuth:    mockrepo.NewMockIAuth(ctrl),
			repoProfile: mockrepo.NewMockIProfile(ctrl),
		}

		t.Run("username empty should return error", func(t *testing.T) {
			userJWT, err := a.LoginUser(context.Background(), dto.ReqLoginUser{
				Username: "",
				Password: "mypassword",
			})

			assert.Empty(t, userJWT)
			require.Error(t, err)
			require.ErrorIs(t, err, gouser.ErrRequestInvalid)
		})
		t.Run("password empty should return error", func(t *testing.T) {
			userJWT, err := a.LoginUser(context.Background(), dto.ReqLoginUser{
				Username: "hidayat",
				Password: "",
			})

			assert.Empty(t, userJWT)
			require.Error(t, err)
			require.ErrorIs(t, err, gouser.ErrRequestInvalid)
		})
	})
}

func TestUnitAuthRegisterUser(t *testing.T) {
	t.Parallel()

	t.Run("register user success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoAuth := mockrepo.NewMockIAuth(ctrl)
		repoProfile := mockrepo.NewMockIProfile(ctrl)

		a := &Auth{
			cfg:         config.Config{},
			repoAuth:    repoAuth,
			repoProfile: repoProfile,
		}

		repoAuth.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(int64(34), nil)

		userID, err := a.RegisterUser(context.Background(), dto.ReqRegisterUser{
			Username: "hidayat",
			Password: "mypassword",
		})

		assert.NotEmpty(t, userID)
		assert.Equal(t, int64(34), userID)
		require.NoError(t, err)
	})
	t.Run("call repo RegisterUser error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoAuth := mockrepo.NewMockIAuth(ctrl)
		repoProfile := mockrepo.NewMockIProfile(ctrl)

		a := &Auth{
			cfg:         config.Config{},
			repoAuth:    repoAuth,
			repoProfile: repoProfile,
		}

		repoAuth.EXPECT().
			RegisterUser(gomock.Any(), gomock.Any()).
			Return(int64(0), assert.AnError)

		userID, err := a.RegisterUser(context.Background(), dto.ReqRegisterUser{
			Username: "hidayat",
			Password: "mypassword",
		})

		assert.Empty(t, userID)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("generate hash password error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoAuth := mockrepo.NewMockIAuth(ctrl)
		repoProfile := mockrepo.NewMockIProfile(ctrl)

		a := &Auth{
			cfg:         config.Config{},
			repoAuth:    repoAuth,
			repoProfile: repoProfile,
		}

		userID, err := a.RegisterUser(context.Background(), dto.ReqRegisterUser{
			Username: "hidayat",
			Password: uuid.NewString() + uuid.NewString() + uuid.NewString(), // password > 72 bytes will error bcrypt
		})

		assert.Empty(t, userID)
		require.Error(t, err)
	})
	t.Run("validate error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoAuth := mockrepo.NewMockIAuth(ctrl)
		repoProfile := mockrepo.NewMockIProfile(ctrl)

		a := &Auth{
			cfg:         config.Config{},
			repoAuth:    repoAuth,
			repoProfile: repoProfile,
		}

		t.Run("empty username should return error", func(t *testing.T) {
			userID, err := a.RegisterUser(context.Background(), dto.ReqRegisterUser{
				Username: "",
				Password: "mypassword",
			})

			assert.Empty(t, userID)
			require.Error(t, err)
			require.ErrorIs(t, err, gouser.ErrRequestInvalid)
		})
		t.Run("empty password should return error", func(t *testing.T) {
			userID, err := a.RegisterUser(context.Background(), dto.ReqRegisterUser{
				Username: "",
				Password: "mypassword",
			})

			assert.Empty(t, userID)
			require.Error(t, err)
			require.ErrorIs(t, err, gouser.ErrRequestInvalid)
		})
	})
}
