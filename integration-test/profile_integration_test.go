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

func TestHTTPGetProfileNotFoundProfile(t *testing.T) {
	assert.Nil(t, Ping())

	Test(t,
		Description("get unknown profile should return error"),
		Get(URLUser+"/"+uuid.NewString()),
		Send().Headers(HContentType).Add(HAppJSON),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".error").NotEqual(nil),
	)
}

func TestHTTPGetProfileSuccess(t *testing.T) {
	assert.Nil(t, Ping())

	username := uuid.NewString()
	password := uuid.NewString()
	Test(t,
		Description("register new user should return success"),
		Post(URLAuthRegister),
		Send().Headers(HContentType).Add(HAppJSON),
		Send().Body().String(jutil.ToJSONString(map[string]string{
			"username": username,
			"password": password,
		})),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".error").Equal(nil),
	)
	Test(t,
		Description("get new registered user profile should return success"),
		Get(URLUser+"/"+username),
		Send().Headers(HContentType).Add(HAppJSON),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".data.username").Equal(username),
		Expect().Body().JSON().JQ(".error").Equal(nil),
	)
}
