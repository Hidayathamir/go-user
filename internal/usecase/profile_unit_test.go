package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/dto"
	"github.com/Hidayathamir/go-user/internal/entity"
	"github.com/Hidayathamir/go-user/internal/repo/mockrepo"
	"github.com/Hidayathamir/go-user/pkg/gouser"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUnitProfileGetProfileByUsername(t *testing.T) {
	t.Parallel()

	t.Run("get profile by username success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProfile := mockrepo.NewMockIProfile(ctrl)

		p := &Profile{
			cfg:         config.Config{},
			repoProfile: repoProfile,
		}

		repoProfile.EXPECT().
			GetProfileByUsername(gomock.Any(), "hidayat").
			Return(entity.User{
				ID:        124,
				Username:  "hidayat",
				Password:  "dummyhashedpassword",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			}, nil)

		profile, err := p.GetProfileByUsername(context.Background(), "hidayat")

		assert.NotEmpty(t, profile)
		assert.Equal(t, dto.ResGetProfileByUsername{
			ID:        124,
			Username:  "hidayat",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}, profile)
		require.NoError(t, err)
	})
	t.Run("call repo GetProfileByUsername error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProfile := mockrepo.NewMockIProfile(ctrl)

		p := &Profile{
			cfg:         config.Config{},
			repoProfile: repoProfile,
		}

		repoProfile.EXPECT().
			GetProfileByUsername(gomock.Any(), "hidayat").
			Return(entity.User{}, assert.AnError)

		profile, err := p.GetProfileByUsername(context.Background(), "hidayat")

		assert.Empty(t, profile)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("usernam empty should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProfile := mockrepo.NewMockIProfile(ctrl)

		p := &Profile{
			cfg:         config.Config{},
			repoProfile: repoProfile,
		}

		profile, err := p.GetProfileByUsername(context.Background(), "")

		assert.Empty(t, profile)
		require.Error(t, err)
		require.ErrorIs(t, err, gouser.ErrRequestInvalid)
	})
}

func TestUnitProfileUpdateProfileByUserID(t *testing.T) {
	t.Parallel()

	t.Run("update profile success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProfile := mockrepo.NewMockIProfile(ctrl)

		cfg := config.Config{
			JWT: config.JWT{ExpireHour: 24, SignedKey: "secretsignkey"},
		}

		p := &Profile{
			cfg:         cfg,
			repoProfile: repoProfile,
		}

		repoProfile.EXPECT().UpdateProfileByUserID(gomock.Any(), gomock.Any()).Return(nil)

		err := p.UpdateProfileByUserID(context.Background(), dto.ReqUpdateProfileByUserID{
			UserJWT:  "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTA1OTY1MDEsInVzZXJfaWQiOjIzfQ.ssdbnErvk_3aKxu3dUmhEFfPr36dhxR7Yx_nvMKzPMs",
			Password: "dummypassword",
		})

		require.NoError(t, err)
	})
	t.Run("call repo UpdateProfileByUserID error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProfile := mockrepo.NewMockIProfile(ctrl)

		cfg := config.Config{
			JWT: config.JWT{ExpireHour: 24, SignedKey: "secretsignkey"},
		}

		p := &Profile{
			cfg:         cfg,
			repoProfile: repoProfile,
		}

		repoProfile.EXPECT().
			UpdateProfileByUserID(gomock.Any(), gomock.Any()).
			Return(assert.AnError)

		err := p.UpdateProfileByUserID(context.Background(), dto.ReqUpdateProfileByUserID{
			UserJWT:  "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTA1OTY1MDEsInVzZXJfaWQiOjIzfQ.ssdbnErvk_3aKxu3dUmhEFfPr36dhxR7Yx_nvMKzPMs",
			Password: "dummypassword",
		})

		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("request validate error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProfile := mockrepo.NewMockIProfile(ctrl)

		cfg := config.Config{
			JWT: config.JWT{ExpireHour: 24, SignedKey: "secretsignkey"},
		}

		p := &Profile{
			cfg:         cfg,
			repoProfile: repoProfile,
		}

		err := p.UpdateProfileByUserID(context.Background(), dto.ReqUpdateProfileByUserID{
			UserJWT:  "",
			Password: "dummypassword",
		})

		require.Error(t, err)
		require.ErrorIs(t, err, gouser.ErrRequestInvalid)
	})
	t.Run("get user id from jwt token string error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProfile := mockrepo.NewMockIProfile(ctrl)

		cfg := config.Config{
			JWT: config.JWT{ExpireHour: 24, SignedKey: "secretsignkey"},
		}

		p := &Profile{
			cfg:         cfg,
			repoProfile: repoProfile,
		}

		err := p.UpdateProfileByUserID(context.Background(), dto.ReqUpdateProfileByUserID{
			UserJWT:  "Bearer dummyuserJWT",
			Password: "dummypassword",
		})

		require.Error(t, err)
		require.ErrorIs(t, err, gouser.ErrJWTAuth)
	})
	t.Run("generate hash password error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repoProfile := mockrepo.NewMockIProfile(ctrl)

		cfg := config.Config{
			JWT: config.JWT{ExpireHour: 24, SignedKey: "secretsignkey"},
		}

		p := &Profile{
			cfg:         cfg,
			repoProfile: repoProfile,
		}

		err := p.UpdateProfileByUserID(context.Background(), dto.ReqUpdateProfileByUserID{
			UserJWT:  "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTA1OTY1MDEsInVzZXJfaWQiOjIzfQ.ssdbnErvk_3aKxu3dUmhEFfPr36dhxR7Yx_nvMKzPMs",
			Password: uuid.NewString() + uuid.NewString() + uuid.NewString(),
		})

		require.Error(t, err)
	})
}
