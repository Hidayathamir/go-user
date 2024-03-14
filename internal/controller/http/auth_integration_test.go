package http

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Hidayathamir/go-user/internal/pkg/auth"
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/Hidayathamir/go-user/internal/usecase/repo"
	"github.com/Hidayathamir/go-user/internal/usecase/repo/db"
	"github.com/Hidayathamir/go-user/pkg/gouser"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationAuthLoginUserRegistered(t *testing.T) {
	t.Parallel()

	cfg := initTestIntegration(t)

	t.Run("user registered try login should success", func(t *testing.T) {
		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
		usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
		controllerAuth := newAuth(cfg, usecaseAuth)

		gin.SetMode(gin.TestMode)

		username := uuid.NewString()
		password := uuid.NewString()
		resBodyRegister := registerUserWithAssertSuccess(t, controllerAuth, username, password)
		resBodyLogin := loginUserWithAssertSuccess(t, cfg, controllerAuth, username, password)
		t.Run("user id in user jwt should equal with user id when register", func(t *testing.T) {
			userID, err := auth.GetUserIDFromJWTTokenString(cfg, resBodyLogin.Data)
			require.NoError(t, err)
			assert.Equal(t, resBodyRegister.Data, userID)
			assert.Nil(t, resBodyLogin.Error)
		})
	})
}

func TestIntegrationAuthLoginUserRegisteredWrongPassword(t *testing.T) { //nolint:dupl
	t.Parallel()

	cfg := initTestIntegration(t)

	t.Run("user registered try login with wrong password should error", func(t *testing.T) {
		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
		usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
		controllerAuth := newAuth(cfg, usecaseAuth)

		gin.SetMode(gin.TestMode)

		username := uuid.NewString()
		registerUserWithAssertSuccess(t, controllerAuth, username, uuid.NewString())

		resBodyByte, httpStatusCode := loginUser(controllerAuth, username, uuid.NewString())
		assert.Equal(t, http.StatusBadRequest, httpStatusCode)
		resBodyLogin := resError{}
		require.NoError(t, json.Unmarshal(resBodyByte, &resBodyLogin))
		assert.Nil(t, resBodyLogin.Data)
		assert.NotEmpty(t, resBodyLogin.Error)
		assert.Contains(t, resBodyLogin.Error, gouser.ErrWrongPassword.Error())
	})
}

func TestIntegrationAuthLoginUserNotRegistered(t *testing.T) {
	t.Parallel()

	cfg := initTestIntegration(t)

	t.Run("user not registered try login should error", func(t *testing.T) {
		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
		usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
		controllerAuth := newAuth(cfg, usecaseAuth)

		gin.SetMode(gin.TestMode)

		resBodyByte, httpStatusCode := loginUser(controllerAuth, uuid.NewString(), uuid.NewString())
		assert.Equal(t, http.StatusBadRequest, httpStatusCode)
		resBodyLogin := resError{}
		require.NoError(t, json.Unmarshal(resBodyByte, &resBodyLogin))
		assert.Nil(t, resBodyLogin.Data)
		assert.NotEmpty(t, resBodyLogin.Error)
		assert.Contains(t, resBodyLogin.Error, gouser.ErrUnknownUsername.Error())
	})
}

func TestIntegrationAuthLoginUserRequestInvalid(t *testing.T) { //nolint:dupl
	t.Parallel()

	cfg := initTestIntegration(t)

	t.Run("login user but request invalid should error", func(t *testing.T) {
		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
		usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
		controllerAuth := newAuth(cfg, usecaseAuth)

		gin.SetMode(gin.TestMode)

		t.Run("request username empty should error", func(t *testing.T) {
			resBodyByte, httpStatusCode := loginUser(controllerAuth, "", uuid.NewString())
			assert.Equal(t, http.StatusBadRequest, httpStatusCode)
			resBodyLogin := resError{}
			require.NoError(t, json.Unmarshal(resBodyByte, &resBodyLogin))
			assert.Nil(t, resBodyLogin.Data)
			assert.NotEmpty(t, resBodyLogin.Error)
			assert.Contains(t, resBodyLogin.Error, gouser.ErrRequestInvalid.Error())
		})
		t.Run("request password empty should error", func(t *testing.T) {
			resBodyByte, httpStatusCode := loginUser(controllerAuth, uuid.NewString(), "")
			assert.Equal(t, http.StatusBadRequest, httpStatusCode)
			resBodyLogin := resError{}
			require.NoError(t, json.Unmarshal(resBodyByte, &resBodyLogin))
			assert.Nil(t, resBodyLogin.Data)
			assert.NotEmpty(t, resBodyLogin.Error)
			assert.Contains(t, resBodyLogin.Error, gouser.ErrRequestInvalid.Error())
		})
	})
}

func TestIntegrationAuthRegisterUserNew(t *testing.T) {
	t.Parallel()

	cfg := initTestIntegration(t)

	t.Run("register new user should success", func(t *testing.T) {
		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
		usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
		controllerAuth := newAuth(cfg, usecaseAuth)

		gin.SetMode(gin.TestMode)

		registerUserWithAssertSuccess(t, controllerAuth, uuid.NewString(), uuid.NewString())
	})
}

func TestIntegrationAuthRegisterUserDuplicate(t *testing.T) { //nolint:dupl
	t.Parallel()

	cfg := initTestIntegration(t)

	t.Run("register user but duplicate username should error", func(t *testing.T) {
		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
		usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
		controllerAuth := newAuth(cfg, usecaseAuth)

		gin.SetMode(gin.TestMode)

		username := uuid.NewString()
		registerUserWithAssertSuccess(t, controllerAuth, username, uuid.NewString())

		resBodyByte, httpStatusCode := registerUser(controllerAuth, username, uuid.NewString())
		assert.Equal(t, http.StatusBadRequest, httpStatusCode)
		resBodyRegister := resError{}
		require.NoError(t, json.Unmarshal(resBodyByte, &resBodyRegister))
		assert.Nil(t, resBodyRegister.Data)
		assert.NotEmpty(t, resBodyRegister.Error)
		assert.Contains(t, resBodyRegister.Error, gouser.ErrDuplicateUsername.Error())
	})
}

func TestIntegrationAuthRegisterUserRequestInvalid(t *testing.T) { //nolint:dupl
	t.Parallel()

	cfg := initTestIntegration(t)

	t.Run("register user but request invalid should error", func(t *testing.T) {
		pg, err := db.NewPostgresPoolConnection(cfg)
		require.NoError(t, err)

		repoAuth := repo.NewAuth(cfg, pg)
		repoProfile := repo.NewProfile(cfg, pg)
		usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
		controllerAuth := newAuth(cfg, usecaseAuth)

		gin.SetMode(gin.TestMode)

		t.Run("request username empty should error", func(t *testing.T) {
			resBodyByte, httpStatusCode := registerUser(controllerAuth, "", uuid.NewString())
			assert.Equal(t, http.StatusBadRequest, httpStatusCode)
			resBodyRegister := resError{}
			require.NoError(t, json.Unmarshal(resBodyByte, &resBodyRegister))
			assert.Nil(t, resBodyRegister.Data)
			assert.NotEmpty(t, resBodyRegister.Error)
			assert.Contains(t, resBodyRegister.Error, gouser.ErrRequestInvalid.Error())
		})
		t.Run("request password empty should error", func(t *testing.T) {
			resBodyByte, httpStatusCode := registerUser(controllerAuth, uuid.NewString(), "")
			assert.Equal(t, http.StatusBadRequest, httpStatusCode)
			resBodyRegister := resError{}
			require.NoError(t, json.Unmarshal(resBodyByte, &resBodyRegister))
			assert.Nil(t, resBodyRegister.Data)
			assert.NotEmpty(t, resBodyRegister.Error)
			assert.Contains(t, resBodyRegister.Error, gouser.ErrRequestInvalid.Error())
		})
	})
}
