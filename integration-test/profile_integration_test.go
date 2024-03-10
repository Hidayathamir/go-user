package integration_test

import (
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
	"github.com/Hidayathamir/go-user/pkg/jutil"
	"github.com/google/uuid"
)

func TestHTTPGetProfileNotFoundProfile(t *testing.T) {
	Test(t,
		Description("get unknown profile should return error"),
		Get(urlUser+"/"+uuid.NewString()),
		Send().Headers(hContentType).Add(hAppJSON),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".error").NotEqual(nil),
	)
}

func TestHTTPGetProfileSuccess(t *testing.T) {
	username := uuid.NewString()
	password := uuid.NewString()
	Test(t,
		Description("register new user should return success"),
		Post(urlAuthRegister),
		Send().Headers(hContentType).Add(hAppJSON),
		Send().Body().String(jutil.ToJSONString(map[string]string{
			"username": username,
			"password": password,
		})),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".error").Equal(nil),
	)
	Test(t,
		Description("get new registered user profile should return success"),
		Get(urlUser+"/"+username),
		Send().Headers(hContentType).Add(hAppJSON),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".data.username").Equal(username),
		Expect().Body().JSON().JQ(".error").Equal(nil),
	)
}
