package grpc

import (
	"context"
	"testing"

	"github.com/Hidayathamir/go-user/internal/db"
	"github.com/Hidayathamir/go-user/internal/pkg/auth"
	"github.com/Hidayathamir/go-user/internal/repo"
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/Hidayathamir/go-user/pkg/gouser"
	"github.com/Hidayathamir/go-user/pkg/gousergrpc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationAuthLoginUser(t *testing.T) {
	t.Parallel()

	t.Run("user registered try login should success", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
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
		t.Run("user id in user jwt should equal with user id when register", func(t *testing.T) {
			userID, err := auth.GetUserIDFromJWTTokenString(cfg, resLogin.GetUserJwt())
			require.NoError(t, err)
			assert.Equal(t, resRegister.GetUserId(), userID)
		})
	})
	t.Run("user registered try login with wrong password should error", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
		usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
		controllerAuth := newAuth(cfg, usecaseAuth)

		username := uuid.NewString()
		resRegister, err := controllerAuth.RegisterUser(context.Background(), &gousergrpc.ReqRegisterUser{
			Username: username,
			Password: uuid.NewString(),
		})
		assert.NotNil(t, resRegister)
		require.NoError(t, err)
		resLogin, err := controllerAuth.LoginUser(context.Background(), &gousergrpc.ReqLoginUser{
			Username: username,
			Password: uuid.NewString(),
		})
		assert.Nil(t, resLogin)
		require.Error(t, err)
		require.ErrorIs(t, err, gouser.ErrWrongPassword)
	})
	t.Run("user not registered try login should error", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
		usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
		controllerAuth := newAuth(cfg, usecaseAuth)

		resLogin, err := controllerAuth.LoginUser(context.Background(), &gousergrpc.ReqLoginUser{
			Username: uuid.NewString(),
			Password: uuid.NewString(),
		})

		assert.Nil(t, resLogin)
		require.Error(t, err)
		require.ErrorIs(t, err, gouser.ErrUnknownUsername)
	})
	t.Run("login try user but request invalid should error", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
		usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
		controllerAuth := newAuth(cfg, usecaseAuth)

		t.Run("request username empty should error", func(t *testing.T) {
			resLogin, err := controllerAuth.LoginUser(context.Background(), &gousergrpc.ReqLoginUser{
				Password: uuid.NewString(),
			})

			assert.Nil(t, resLogin)
			require.Error(t, err)
			require.ErrorIs(t, err, gouser.ErrRequestInvalid)
		})
		t.Run("request password empty should error", func(t *testing.T) {
			resLogin, err := controllerAuth.LoginUser(context.Background(), &gousergrpc.ReqLoginUser{
				Username: uuid.NewString(),
			})

			assert.Nil(t, resLogin)
			require.Error(t, err)
			require.ErrorIs(t, err, gouser.ErrRequestInvalid)
		})
	})
}

func TestIntegrationAuthRegisterUser(t *testing.T) {
	t.Parallel()

	t.Run("register new user should success", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
		usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
		controllerAuth := newAuth(cfg, usecaseAuth)

		res, err := controllerAuth.RegisterUser(context.Background(), &gousergrpc.ReqRegisterUser{
			Username: uuid.NewString(),
			Password: uuid.NewString(),
		})
		assert.NotNil(t, res)
		require.NoError(t, err)
	})
	t.Run("register user but duplicate username should error", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
		usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
		controllerAuth := newAuth(cfg, usecaseAuth)

		username := uuid.NewString()
		res, err := controllerAuth.RegisterUser(context.Background(), &gousergrpc.ReqRegisterUser{
			Username: username,
			Password: uuid.NewString(),
		})
		assert.NotNil(t, res)
		require.NoError(t, err)

		res, err = controllerAuth.RegisterUser(context.Background(), &gousergrpc.ReqRegisterUser{
			Username: username,
			Password: uuid.NewString(),
		})
		assert.Nil(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, gouser.ErrDuplicateUsername)
	})
	t.Run("register user but request invalid should error", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
		usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
		controllerAuth := newAuth(cfg, usecaseAuth)
		t.Run("request username empty should error", func(t *testing.T) {
			res, err := controllerAuth.RegisterUser(context.Background(), &gousergrpc.ReqRegisterUser{
				Password: uuid.NewString(),
			})
			assert.Nil(t, res)
			require.Error(t, err)
			require.ErrorIs(t, err, gouser.ErrRequestInvalid)
		})
		t.Run("request password empty should error", func(t *testing.T) {
			res, err := controllerAuth.RegisterUser(context.Background(), &gousergrpc.ReqRegisterUser{
				Username: uuid.NewString(),
			})
			assert.Nil(t, res)
			require.Error(t, err)
			require.ErrorIs(t, err, gouser.ErrRequestInvalid)
		})
	})
}
