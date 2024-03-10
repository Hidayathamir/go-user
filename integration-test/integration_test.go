package integration_test

import (
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
	"github.com/sirupsen/logrus"
)

var (
	urlPing         = "/ping"
	urlRegisterUser = "/api/v1/auth/register"
)

const (
	hContentType = "Content-Type"
	hAppJSON     = "application/json"
)

func TestMain(m *testing.M) {
	setupURL()

	err := ping()
	if err != nil {
		logrus.Fatalf("ping: %v", err)
	}

	code := m.Run()
	os.Exit(code)
}

func setupURL() {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		panic("BASE_URL not found in env var")
	}

	urlPing = baseURL + urlPing
	urlRegisterUser = baseURL + urlRegisterUser
}

func ping() error {
	var err error
	for i := 0; i < 10; i++ {
		err = Do(Get(urlPing), Expect().Status().Equal(http.StatusOK))
		if err == nil {
			logrus.Infof("ping success 🟢")
			return nil
		}

		logrus.
			WithField("attempt count", i+1).
			Warnf("err ping: %v", err)

		time.Sleep(time.Second)
	}

	return err
}
