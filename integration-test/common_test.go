package integration_test

import (
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
	. "github.com/Hidayathamir/go-user/integration-test/config"
	"github.com/Hidayathamir/go-user/pkg/jutil"
)

// This file contains func related for testing.

// registerUser return userID that being registered.
func registerUser(t *testing.T, username string, password string) int64 {
	var userID int64

	Test(t,
		Description("unique username should register success"),
		Post(URLAuthRegister),
		Send().Headers(HContentType).Add(HAppJSON),
		Send().Body().String(jutil.ToJSONString(map[string]string{
			"username": username,
			"password": password,
		})),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".error").Equal(nil),
		Store().Response().Body().JSON().JQ(".data").In(&userID),
	)

	return userID
}

// loginUser return string user jwt.
func loginUser(t *testing.T, username string, password string) string {
	var userJWT string

	Test(t,
		Description("user already registered try to login should return success"),
		Post(URLAuthLogin),
		Send().Headers(HContentType).Add(HAppJSON),
		Send().Body().String(jutil.ToJSONString(map[string]string{
			"username": username,
			"password": password,
		})),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".error").Equal(nil),
		Store().Response().Body().JSON().JQ(".data").In(&userJWT),
	)

	return userJWT
}
