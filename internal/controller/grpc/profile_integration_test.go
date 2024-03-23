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

func TestIntegrationProfileUpdateProfileByUserID(t *testing.T) {
	t.Parallel()

	t.Run("registered user update profile should success", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPGPoolConn(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
		usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
		controllerAuth := newAuth(cfg, usecaseAuth)

		usecaseProfile := usecase.NewProfile(cfg, repoProfile)
		controllerProfile := newProfile(cfg, usecaseProfile)

		username := uuid.NewString()
		oldPassword := uuid.NewString()
		resRegister, err := controllerAuth.RegisterUser(context.Background(), &gousergrpc.ReqRegisterUser{
			Username: username,
			Password: oldPassword,
		})
		assert.NotNil(t, resRegister)
		require.NoError(t, err)

		resLogin, err := controllerAuth.LoginUser(context.Background(), &gousergrpc.ReqLoginUser{
			Username: username,
			Password: oldPassword,
		})
		assert.NotNil(t, resLogin)
		require.NoError(t, err)

		newPassword := uuid.NewString()
		resUpdate, err := controllerProfile.UpdateProfileByUserID(context.Background(), &gousergrpc.ReqUpdateProfileByUserID{
			UserJwt:  resLogin.GetUserJwt(),
			Password: newPassword,
		})
		assert.NotNil(t, resUpdate)
		require.NoError(t, err)

		t.Run("after update profile should able to login with new password", func(t *testing.T) {
			resLogin, err = controllerAuth.LoginUser(context.Background(), &gousergrpc.ReqLoginUser{
				Username: username,
				Password: newPassword,
			})
			assert.NotNil(t, resLogin)
			require.NoError(t, err)
		})
		t.Run("after update profile should not able to login with old password", func(t *testing.T) {
			resLogin, err = controllerAuth.LoginUser(context.Background(), &gousergrpc.ReqLoginUser{
				Username: username,
				Password: oldPassword,
			})
			assert.Nil(t, resLogin)
			require.Error(t, err)
			require.ErrorIs(t, err, gouser.ErrWrongPassword)
		})
	})
	t.Run("update profile but request invalid should error", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPGPoolConn(cfg)
		require.NoError(t, err)

		repoProfile := repo.NewProfile(cfg, pg)
		usecaseProfile := usecase.NewProfile(cfg, repoProfile)
		controllerProfile := newProfile(cfg, usecaseProfile)

		t.Run("request user jwt empty should error", func(t *testing.T) {
			res, err := controllerProfile.UpdateProfileByUserID(context.Background(), &gousergrpc.ReqUpdateProfileByUserID{
				Password: uuid.NewString(),
			})
			assert.Nil(t, res)
			require.Error(t, err)
			require.ErrorIs(t, err, gouser.ErrRequestInvalid)
		})
		t.Run("request header user jwt wrong should error", func(t *testing.T) {
			res, err := controllerProfile.UpdateProfileByUserID(context.Background(), &gousergrpc.ReqUpdateProfileByUserID{
				UserJwt:  "sdf",
				Password: uuid.NewString(),
			})
			assert.Nil(t, res)
			require.Error(t, err)
			require.ErrorIs(t, err, gouser.ErrJWTAuth)
		})
		t.Run("request password empty should error", func(t *testing.T) {
			repoAuth := repo.NewAuth(cfg, pg)
			usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
			controllerAuth := newAuth(cfg, usecaseAuth)

			username := uuid.NewString()
			password := uuid.NewString()

			resRegister, err := controllerAuth.RegisterUser(context.Background(), &gousergrpc.ReqRegisterUser{
				Username: username,
				Password: password,
			})
			assert.NotNil(t, resRegister)
			require.NoError(t, err)

			resLogin, err := controllerAuth.LoginUser(context.Background(), &gousergrpc.ReqLoginUser{
				Username: username,
				Password: password,
			})
			assert.NotNil(t, resLogin)
			require.NoError(t, err)

			resUpdate, err := controllerProfile.UpdateProfileByUserID(context.Background(), &gousergrpc.ReqUpdateProfileByUserID{
				UserJwt: resLogin.GetUserJwt(),
			})
			assert.Nil(t, resUpdate)
			require.Error(t, err)
			require.ErrorIs(t, err, gouser.ErrNothingToBeUpdate)
		})
	})
}

func TestIntegrationProfileGetProfileByUsername(t *testing.T) {
	t.Parallel()

	t.Run("get profile known user should success", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPGPoolConn(cfg)
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

		pg, err := db.NewPGPoolConn(cfg)
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
