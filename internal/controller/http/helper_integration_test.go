package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/pkg/auth"
	"github.com/Hidayathamir/go-user/internal/pkg/header"
	"github.com/Hidayathamir/go-user/internal/pkg/jutil"
	"github.com/Hidayathamir/go-user/internal/repo"
	"github.com/Hidayathamir/go-user/internal/repo/db"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// contains helper for integration test.

func configInit(t *testing.T) config.Config {
	t.Helper()

	require.NoError(t, os.Setenv("LOGGER_LOG_LEVEL", "fatal"))

	configYamlPath := filepath.Join("..", "..", "..", "config", "config.yml")
	cfg, err := config.Init(&config.EnvLoader{YAMLPath: configYamlPath})
	require.NoError(t, err)
	return cfg
}

type mute struct{}

func (n mute) Printf(string, ...interface{}) {}

// createPGContainer create pg container.
func createPGContainer(t *testing.T, cfg config.Config) *postgres.PostgresContainer {
	t.Helper()

	pgContainer, err := postgres.RunContainer(context.Background(),
		testcontainers.WithLogger(&mute{}),
		testcontainers.WithImage("postgres:16"),
		postgres.WithDatabase(cfg.PG.DBName),
		postgres.WithUsername(cfg.PG.Username),
		postgres.WithPassword(cfg.PG.Password),
		testcontainers.WithWaitStrategy(
			wait.
				ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	require.NoError(t, err)
	return pgContainer
}

// updateConfigPGPort update config port based on pg container, do db migrations.
func updateConfigPGPort(t *testing.T, cfg *config.Config, pgContainer *postgres.PostgresContainer) {
	t.Helper()

	dbURL, err := pgContainer.ConnectionString(context.Background())
	require.NoError(t, err)

	port, err := repo.GetPort(dbURL)
	require.NoError(t, err)

	cfg.PG.Port = port
}

// dbMigrateUp do db migrations.
func dbMigrateUp(t *testing.T, cfg config.Config) {
	t.Helper()

	schemaMigrationPath := filepath.Join("..", "..", "..", "internal", "repo", "db", "schema_migration")
	require.NoError(t, db.MigrateUp(cfg, schemaMigrationPath))
}

// initTestIntegration init config, create pg container, update config port
// based on pg container, do db migrations.
func initTestIntegration(t *testing.T) config.Config {
	t.Helper()

	cfg := configInit(t)

	pgContainer := createPGContainer(t, cfg)
	t.Cleanup(func() { require.NoError(t, pgContainer.Terminate(context.Background())) })

	updateConfigPGPort(t, &cfg, pgContainer)

	dbMigrateUp(t, cfg)

	return cfg
}

//////////////////////////////////////////

// registerUser registers user return raw response and http status code.
func registerUser(controllerAuth *Auth, username string, password string) (resBody []byte, httpStatusCode int) {
	rr := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rr)
	reqBody := bytes.NewReader([]byte(jutil.ToJSONString(map[string]string{
		"username": username,
		"password": password,
	})))
	ctx.Request = httptest.NewRequest(http.MethodPost, "/", reqBody)

	controllerAuth.registerUser(ctx)

	return rr.Body.Bytes(), rr.Code
}

// registerUserWithAssertSuccess registers user then validate, return response
// body register success.
func registerUserWithAssertSuccess(t *testing.T, controllerAuth *Auth, username string, password string) ResRegisterUser {
	t.Helper()

	resBodyByte, httpStatusCode := registerUser(controllerAuth, username, password)
	assert.Equal(t, http.StatusOK, httpStatusCode)
	resBody := ResRegisterUser{}
	require.NoError(t, json.Unmarshal(resBodyByte, &resBody))
	assert.NotEmpty(t, resBody.Data)
	assert.Nil(t, resBody.Error)
	return resBody
}

// loginUser login user return raw response and http status code.
func loginUser(controllerAuth *Auth, username string, password string) (resBody []byte, httpStatusCode int) {
	rr := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rr)
	reqBody := bytes.NewReader([]byte(jutil.ToJSONString(map[string]string{
		"username": username,
		"password": password,
	})))
	ctx.Request = httptest.NewRequest(http.MethodPost, "/", reqBody)

	controllerAuth.loginUser(ctx)

	return rr.Body.Bytes(), rr.Code
}

// loginUserWithAssertSuccess login user then validate, return response body
// login success.
func loginUserWithAssertSuccess(t *testing.T, cfg config.Config, controllerAuth *Auth, username string, password string) ResLoginUser {
	t.Helper()

	resBodyByte, httpStatusCode := loginUser(controllerAuth, username, password)
	assert.Equal(t, http.StatusOK, httpStatusCode)
	resBody := ResLoginUser{}
	require.NoError(t, json.Unmarshal(resBodyByte, &resBody))
	assert.NotEmpty(t, resBody.Data)
	assert.Contains(t, resBody.Data.UserJWT, "Bearer")
	_, err := auth.GetUserIDFromJWTTokenString(cfg, resBody.Data.UserJWT)
	require.NoError(t, err)
	return resBody
}

// updateProfileByUserID update user profile by id return raw response and http status code.
func updateProfileByUserID(controllerProfile *Profile, userJWT string, newPassword string) (resBody []byte, httpStatusCode int) {
	rr := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rr)
	reqBody := bytes.NewReader([]byte(jutil.ToJSONString(map[string]string{
		"password": newPassword,
	})))
	ctx.Request = httptest.NewRequest(http.MethodPut, "/", reqBody)
	ctx.Request.Header.Set(header.Authorization, userJWT)

	controllerProfile.updateProfileByUserID(ctx)

	return rr.Body.Bytes(), rr.Code
}

// getProfileByUsername get user profile by username return raw response and http status code.
func getProfileByUsername(controllerProfile *Profile, username string) (resBody []byte, httpStatusCode int) {
	rr := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rr)
	ctx.Request = httptest.NewRequest(http.MethodPut, "/", nil)
	ctx.Params = append(ctx.Params, gin.Param{Key: "username", Value: username})

	controllerProfile.getProfileByUsername(ctx)

	return rr.Body.Bytes(), rr.Code
}
