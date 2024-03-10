package integration_test

import (
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
	. "github.com/Hidayathamir/go-user/integration-test/config"
	"github.com/Hidayathamir/go-user/pkg/jutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHTTPRegisterUserUniqueUsername(t *testing.T) {
	assert.Nil(t, Ping())

	registerUser(t, uuid.NewString(), uuid.NewString())
}

func TestHTTPRegisterUserDuplicateUsername(t *testing.T) {
	assert.Nil(t, Ping())

	duplicateUsername := uuid.NewString()

	registerUser(t, duplicateUsername, uuid.NewString())

	Test(t,
		Description("duplicate username should register fail, second is fail"),
		Post(URLAuthRegister),
		Send().Headers(HContentType).Add(HAppJSON),
		Send().Body().String(jutil.ToJSONString(map[string]string{
			"username": duplicateUsername,
			"password": uuid.NewString(),
		})),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".error").NotEqual(nil),
	)
}

func TestHTTPLoginUserNotRegister(t *testing.T) {
	assert.Nil(t, Ping())

	Test(t,
		Description("user not registered try to login should return error"),
		Post(URLAuthLogin),
		Send().Headers(HContentType).Add(HAppJSON),
		Send().Body().String(jutil.ToJSONString(map[string]string{
			"username": uuid.NewString(),
			"password": uuid.NewString(),
		})),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".error").NotEqual(nil),
	)
}

func TestHTTPLoginUserAlreadyRegister(t *testing.T) {
	assert.Nil(t, Ping())

	username := uuid.NewString()
	password := uuid.NewString()

	registerUser(t, username, password)

	loginUser(t, username, password)
}
