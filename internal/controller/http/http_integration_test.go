package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/entity/table"
	"github.com/Hidayathamir/go-user/internal/usecase/repo/db"
	"github.com/Hidayathamir/go-user/pkg/auth"
	"github.com/Hidayathamir/go-user/pkg/jutil"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func configInit(t *testing.T) config.Config {
	configYamlPath := filepath.Join("..", "..", "..", "config", "config.yml")
	cfg, err := config.Init(&config.YamlLoader{Path: configYamlPath})
	if err != nil {
		t.Fatalf("config.Init: %v", err)
	}
	return cfg
}

func createPGContainer(cfg config.Config) (*postgres.PostgresContainer, error) {
	pgContainer, err := postgres.RunContainer(context.Background(),
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
	if err != nil {
		return nil, fmt.Errorf("postgres.RunContainer: %w", err)
	}

	return pgContainer, nil
}

func terminatePGContainer(t *testing.T, pgContainer *postgres.PostgresContainer) func() {
	return func() {
		err := pgContainer.Terminate(context.Background())
		if err != nil {
			t.Fatalf("postgres.PostgresContainer.Terminate: %v", err)
		}
	}
}

func updateConfigPGPort(t *testing.T, cfg *config.Config, pgContainer *postgres.PostgresContainer) {
	dbURL, err := pgContainer.ConnectionString(context.Background())
	if err != nil {
		t.Fatalf("postgres.PostgresContainer.ConnectionString: %v", err)
	}

	connConfig, err := pgx.ParseConfig(dbURL)
	if err != nil {
		t.Fatalf("pgx.ParseConfig: %v", err)
	}

	cfg.PG.Port = int(connConfig.Port)
}

func dbMigrateUp(t *testing.T, cfg config.Config) {
	schemaMigrationPath := filepath.Join("..", "..", "..", "internal", "usecase", "repo", "db", "schema_migration")
	err := db.MigrateUp(cfg, schemaMigrationPath)
	if err != nil {
		t.Fatalf("db.MigrateUp: %v", err)
	}
}

func initTestIntegration(t *testing.T) config.Config {
	cfg := configInit(t)

	table.Init()

	pgContainer, err := createPGContainer(cfg)
	if err != nil {
		t.Fatalf("createPGContainer: %v", err)
	}
	t.Cleanup(terminatePGContainer(t, pgContainer))

	updateConfigPGPort(t, &cfg, pgContainer)

	dbMigrateUp(t, cfg)

	return cfg
}

type ResRegisterUserSuccess struct {
	Data  int64 `json:"data"` // userID
	Error any   `json:"error"`
}

type ResRegisterUserError struct {
	Data  any    `json:"data"`
	Error string `json:"error"`
}

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
func registerUserWithAssertSuccess(t *testing.T, controllerAuth *Auth, username string, password string) ResRegisterUserSuccess {
	resBodyByte, httpStatusCode := registerUser(controllerAuth, username, password)
	assert.Equal(t, http.StatusOK, httpStatusCode)
	resBody := ResRegisterUserSuccess{}
	require.NoError(t, json.Unmarshal(resBodyByte, &resBody))
	assert.NotEmpty(t, resBody.Data)
	assert.IsType(t, int64(1), resBody.Data)
	assert.Nil(t, resBody.Error)
	return resBody
}

type ResLoginUserSuccess struct {
	Data  string `json:"data"` // userJWT
	Error any    `json:"error"`
}

type ResLoginUserError struct {
	Data  any    `json:"data"`
	Error string `json:"error"`
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
func loginUserWithAssertSuccess(t *testing.T, cfg config.Config, controllerAuth *Auth, username string, password string) ResLoginUserSuccess {
	resBodyByte, httpStatusCode := loginUser(controllerAuth, username, password)
	assert.Equal(t, http.StatusOK, httpStatusCode)
	resBody := ResLoginUserSuccess{}
	require.NoError(t, json.Unmarshal(resBodyByte, &resBody))
	assert.NotEmpty(t, resBody.Data)
	assert.Contains(t, resBody.Data, "Bearer")
	_, err := auth.ValidateUserJWTToken(cfg, resBody.Data)
	require.NoError(t, err)
	return resBody
}
