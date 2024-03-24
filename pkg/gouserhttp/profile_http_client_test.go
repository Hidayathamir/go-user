package gouserhttp

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/Hidayathamir/go-user/internal/controller/http"
	"github.com/Hidayathamir/go-user/internal/repo/db"
	"github.com/Hidayathamir/go-user/pkg/gouser"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test positive case http client.
// This is not integration test, for full test case scenario see
// http integration test in internal/controller/http package.

func TestHTTPClientUpdateProfile(t *testing.T) {
	t.Parallel()

	cfg := initTestIntegration(t)

	pg, err := db.NewPGPoolConn(cfg)
	require.NoError(t, err)

	go func() {
		gin.SetMode(gin.TestMode)
		err := http.RunServer(cfg, pg)
		assert.NoError(t, err)
	}()

	time.Sleep(time.Second * 1) // wait http server run.

	baseURL := "http://" + cfg.HTTP.Host + ":" + strconv.Itoa(cfg.HTTP.Port)
	gouserAuthClient := NewAuthClient(baseURL)

	username := uuid.NewString()
	password := uuid.NewString()

	reqRegister := gouser.ReqRegisterUser{
		Username: username,
		Password: password,
	}
	_, err = gouserAuthClient.RegisterUser(context.Background(), reqRegister)
	require.NoError(t, err)

	reqLogin := gouser.ReqLoginUser{
		Username: username,
		Password: password,
	}
	resLogin, err := gouserAuthClient.LoginUser(context.Background(), reqLogin)
	require.NoError(t, err)

	gouserProfileClient := NewProfileClient(baseURL)

	reqUpdateProfile := gouser.ReqUpdateProfileByUserID{
		UserJWT:  resLogin.UserJWT,
		Password: password,
	}
	err = gouserProfileClient.UpdateProfileByUserID(context.Background(), reqUpdateProfile)
	require.NoError(t, err)
}

func TestHTTPClientGetProfile(t *testing.T) {
	t.Parallel()

	cfg := initTestIntegration(t)

	pg, err := db.NewPGPoolConn(cfg)
	require.NoError(t, err)

	go func() {
		gin.SetMode(gin.TestMode)
		err := http.RunServer(cfg, pg)
		assert.NoError(t, err)
	}()

	time.Sleep(time.Second * 1) // wait http server run.

	baseURL := "http://" + cfg.HTTP.Host + ":" + strconv.Itoa(cfg.HTTP.Port)

	gouserAuthClient := NewAuthClient(baseURL)

	username := uuid.NewString()
	password := uuid.NewString()

	reqRegister := gouser.ReqRegisterUser{
		Username: username,
		Password: password,
	}
	_, err = gouserAuthClient.RegisterUser(context.Background(), reqRegister)
	require.NoError(t, err)

	gouserProfileClient := NewProfileClient(baseURL)

	reqGetProfile := gouser.ReqGetProfileByUsername{
		Username: username,
	}
	resGetProfile, err := gouserProfileClient.GetProfileByUsername(context.Background(), reqGetProfile)
	require.NoError(t, err)

	logrus.Info(resGetProfile)
}
