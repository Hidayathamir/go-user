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

	Test(t,
		Description("unique username should register success"),
		Post(URLAuthRegister),
		Send().Headers(HContentType).Add(HAppJSON),
		Send().Body().String(jutil.ToJSONString(map[string]string{
			"username": uuid.NewString(),
			"password": uuid.NewString(),
		})),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".error").Equal(nil),
	)
}

func TestHTTPRegisterUserDuplicateUsername(t *testing.T) {
	assert.Nil(t, Ping())

	duplicateUsername := uuid.NewString()
	Test(t,
		Description("duplicate username should register fail, first is success"),
		Post(URLAuthRegister),
		Send().Headers(HContentType).Add(HAppJSON),
		Send().Body().String(jutil.ToJSONString(map[string]string{
			"username": duplicateUsername,
			"password": uuid.NewString(),
		})),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".error").Equal(nil),
	)
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
