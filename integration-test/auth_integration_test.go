package integration_test

import (
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
	"github.com/Hidayathamir/go-user/pkg/jutil"
	"github.com/google/uuid"
)

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
