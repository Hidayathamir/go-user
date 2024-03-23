package grpc

import (
	"context"
	"testing"

	"github.com/Hidayathamir/go-user/internal/db"
	"github.com/Hidayathamir/go-user/internal/repo"
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/Hidayathamir/go-user/pkg/gouser"
	"github.com/Hidayathamir/go-user/pkg/gousergrpc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationProfileGetProfileByUsername(t *testing.T) {
	t.Parallel()

	t.Run("get profile known user should success", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
		usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
		controllerAuth := newAuth(cfg, usecaseAuth)

		usecaseProfile := usecase.NewProfile(cfg, repoProfile)
		controllerProfile := newProfile(cfg, usecaseProfile)

		username := uuid.NewString()
		resRegister, err := controllerAuth.RegisterUser(context.Background(), &gousergrpc.ReqRegisterUser{
			Username: username,
			Password: uuid.NewString(),
		})
		assert.NotNil(t, resRegister)
		require.NoError(t, err)

		resProfile, err := controllerProfile.GetProfileByUsername(context.Background(), &gousergrpc.ReqGetProfileByUsername{
			Username: username,
		})
		assert.NotNil(t, resProfile)
		require.NoError(t, err)

		assert.Equal(t, resRegister.GetUserId(), resProfile.GetId())
		assert.Equal(t, username, resProfile.GetUsername())
	})
	t.Run("get profile unknown user should error", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoProfile := repo.NewProfile(cfg, pg)
		usecaseProfile := usecase.NewProfile(cfg, repoProfile)
		controllerProfile := newProfile(cfg, usecaseProfile)

		res, err := controllerProfile.GetProfileByUsername(context.Background(), &gousergrpc.ReqGetProfileByUsername{
			Username: uuid.NewString(),
		})
		assert.Nil(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, gouser.ErrUnknownUsername)
	})
}
