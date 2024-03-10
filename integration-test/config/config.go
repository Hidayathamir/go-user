// Package config contains configuration for integration test.
package config

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Eun/go-hit"
	"github.com/sirupsen/logrus"
)

// Constant Header.
const (
	HContentType = "Content-Type"
	HAppJSON     = "application/json"
)

// URL for testing.
var (
	URLPing         = getBaseURL() + "/ping"
	URLAuthRegister = getBaseURL() + "/api/v1/auth/register"
	URLAuthLogin    = getBaseURL() + "/api/v1/auth/login"
	URLUser         = getBaseURL() + "/api/v1/users"
)

func getBaseURL() string {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	return baseURL
}

// Ping hit ping api.
func Ping() error {
	var err error
	for i := 0; i < 10; i++ {
		err = hit.Do(hit.Get(URLPing), hit.Expect().Status().Equal(http.StatusOK))
		if err == nil {
			return nil
		}

		logrus.
			WithFields(logrus.Fields{
				"attempt count": i + 1,
				"url ping":      URLPing,
			}).
			Warn("err ping")

		time.Sleep(time.Second)
	}

	return fmt.Errorf("err ping: %w", err)
}
