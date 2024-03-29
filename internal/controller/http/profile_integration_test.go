package http

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Hidayathamir/go-user/internal/repo"
	"github.com/Hidayathamir/go-user/internal/repo/db"
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/Hidayathamir/go-user/pkg/gouser"
	"github.com/gin-gonic/gin"
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

		gin.SetMode(gin.TestMode)

		username := uuid.NewString()
		oldPassword := uuid.NewString()
		registerUserWithAssertSuccess(t, controllerAuth, username, oldPassword)
		resBodyLogin := loginUserWithAssertSuccess(t, cfg, controllerAuth, username, oldPassword)

		newPassword := uuid.NewString()
		resBodyByte, httpStatusCode := updateProfileByUserID(controllerProfile, resBodyLogin.Data.UserJWT, newPassword)
		assert.Equal(t, http.StatusOK, httpStatusCode)
		resBody := ResUpdatePofile{}
		require.NoError(t, json.Unmarshal(resBodyByte, &resBody))
		assert.NotEmpty(t, resBody.Data)
		assert.Contains(t, resBody.Data, "ok")
		assert.Nil(t, resBody.Error)

		t.Run("after update profile should able to login with new password", func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			loginUserWithAssertSuccess(t, cfg, controllerAuth, username, newPassword)
		})

		t.Run("after update profile should not able to login with old password", func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			resBodyByte, httpStatusCode = loginUser(controllerAuth, username, oldPassword)
			assert.Equal(t, http.StatusBadRequest, httpStatusCode)
			resBodyLogin2 := ResError{}
			require.NoError(t, json.Unmarshal(resBodyByte, &resBodyLogin2))
			assert.Nil(t, resBodyLogin2.Data)
			assert.NotEmpty(t, resBodyLogin2.Error)
			assert.Contains(t, resBodyLogin2.Error, gouser.ErrWrongPassword.Error())
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

		gin.SetMode(gin.TestMode)

		t.Run("request header user jwt empty should error", func(t *testing.T) {
			resBodyByte, httpStatusCode := updateProfileByUserID(controllerProfile, "", uuid.NewString())
			assert.Equal(t, http.StatusBadRequest, httpStatusCode)
			resBodyUpdate := ResError{}
			require.NoError(t, json.Unmarshal(resBodyByte, &resBodyUpdate))
			assert.Nil(t, resBodyUpdate.Data)
			assert.NotEmpty(t, resBodyUpdate.Error)
			assert.Contains(t, resBodyUpdate.Error, gouser.ErrRequestInvalid.Error())
		})
		t.Run("request header user jwt wrong should error", func(t *testing.T) {
			resBodyByte, httpStatusCode := updateProfileByUserID(controllerProfile, "sdf", uuid.NewString())
			assert.Equal(t, http.StatusBadRequest, httpStatusCode)
			resBodyUpdate := ResError{}
			require.NoError(t, json.Unmarshal(resBodyByte, &resBodyUpdate))
			assert.Nil(t, resBodyUpdate.Data)
			assert.NotEmpty(t, resBodyUpdate.Error)
			assert.Contains(t, resBodyUpdate.Error, gouser.ErrJWTAuth.Error())
		})
		t.Run("request password empty should error", func(t *testing.T) {
			repoAuth := repo.NewAuth(cfg, pg)
			usecaseAuth := usecase.NewAuth(cfg, repoAuth, repoProfile)
			controllerAuth := newAuth(cfg, usecaseAuth)

			username := uuid.NewString()
			password := uuid.NewString()
			registerUserWithAssertSuccess(t, controllerAuth, username, password)
			resBodyLogin := loginUserWithAssertSuccess(t, cfg, controllerAuth, username, password)
			resBodyByte, httpStatusCode := updateProfileByUserID(controllerProfile, resBodyLogin.Data.UserJWT, "")
			assert.Equal(t, http.StatusBadRequest, httpStatusCode)
			resBodyUpdate := ResError{}
			require.NoError(t, json.Unmarshal(resBodyByte, &resBodyUpdate))
			assert.Nil(t, resBodyUpdate.Data)
			assert.NotEmpty(t, resBodyUpdate.Error)
			assert.Contains(t, resBodyUpdate.Error, gouser.ErrNothingToBeUpdate.Error())
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

		gin.SetMode(gin.TestMode)

		username := "hidayat"
		resBodyRegister := registerUserWithAssertSuccess(t, controllerAuth, username, "mypassword")

		resBodyByte, httpStatusCode := getProfileByUsername(controllerProfile, username)

		assert.Equal(t, http.StatusOK, httpStatusCode)
		resBodyGetProfile := ResGetProfileByUsername{}
		require.NoError(t, json.Unmarshal(resBodyByte, &resBodyGetProfile))
		assert.Equal(t, resBodyRegister.Data.UserID, resBodyGetProfile.Data.ID)
		assert.Equal(t, username, resBodyGetProfile.Data.Username)
		assert.Nil(t, resBodyGetProfile.Error)
	})
	t.Run("get profile unknown user should error", func(t *testing.T) {
		t.Parallel()

		cfg := initTestIntegration(t)

		pg, err := db.NewPGPoolConn(cfg)
		require.NoError(t, err)

		repoProfile := repo.NewProfile(cfg, pg)
		usecaseProfile := usecase.NewProfile(cfg, repoProfile)
		controllerProfile := newProfile(cfg, usecaseProfile)

		gin.SetMode(gin.TestMode)

		resBodyByte, httpStatusCode := getProfileByUsername(controllerProfile, uuid.NewString())

		assert.Equal(t, http.StatusBadRequest, httpStatusCode)
		resBody := ResError{}
		require.NoError(t, json.Unmarshal(resBodyByte, &resBody))
		assert.Nil(t, resBody.Data)
		assert.NotEmpty(t, resBody.Error)
		assert.Contains(t, resBody.Error, gouser.ErrUnknownUsername.Error())
	})
}
