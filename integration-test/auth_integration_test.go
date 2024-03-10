package integration_test

import (
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
	"github.com/Hidayathamir/go-user/pkg/jutil"
	"github.com/google/uuid"
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

func TestHTTPRegisterUser(t *testing.T) {
	Test(t,
		Description("unique username should register success"),
		Post(urlRegisterUser),
		Send().Headers(hContentType).Add(hAppJSON),
		Send().Body().String(jutil.ToJSONString(map[string]string{
			"username": uuid.NewString(),
			"password": uuid.NewString(),
		})),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".error").Equal(nil),
	)

	duplicateUsername := uuid.NewString()
	Test(t,
		Description("duplicate username should register fail, first is success"),
		Post(urlRegisterUser),
		Send().Headers(hContentType).Add(hAppJSON),
		Send().Body().String(jutil.ToJSONString(map[string]string{
			"username": duplicateUsername,
			"password": uuid.NewString(),
		})),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".error").Equal(nil),
	)
	Test(t,
		Description("duplicate username should register fail, second is fail"),
		Post(urlRegisterUser),
		Send().Headers(hContentType).Add(hAppJSON),
		Send().Body().String(jutil.ToJSONString(map[string]string{
			"username": duplicateUsername,
			"password": uuid.NewString(),
		})),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".error").NotEqual(nil),
	)
}
