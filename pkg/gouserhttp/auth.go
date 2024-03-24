package gouserhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	controllerHTTP "github.com/Hidayathamir/go-user/internal/controller/http"
	"github.com/Hidayathamir/go-user/pkg/header"
	"github.com/sirupsen/logrus"
)

// API path list.
const (
	APIAuthRegister = "/api/v1/auth/register"
)

// IAuthClient -.
type IAuthClient interface {
	LoginUser(ctx context.Context, req ReqLoginUser) (ResLoginUser, error)
	RegisterUser(ctx context.Context, req ReqRegisterUser) (ResRegisterUser, error)
}

// AuthClient -.
type AuthClient struct {
	// BaseURL eg. http://localhost:8080.
	BaseURL string
}

var _ IAuthClient = &AuthClient{}

// NewAuthClient -.
func NewAuthClient(baseURL string) *AuthClient {
	return &AuthClient{
		BaseURL: baseURL,
	}
}

// LoginUser implements AuthClient.
func (a *AuthClient) LoginUser(context.Context, ReqLoginUser) (ResLoginUser, error) {
	panic("unimplemented") // TODO: IMPLEMENT
}

// RegisterUser implements AuthClient.
func (a *AuthClient) RegisterUser(ctx context.Context, req ReqRegisterUser) (ResRegisterUser, error) {
	url := a.BaseURL + APIAuthRegister

	fail := func(msg string, err error) (ResRegisterUser, error) {
		return ResRegisterUser{}, fmt.Errorf(msg+": %w", err)
	}

	reqJSONByte, err := json.Marshal(req)
	if err != nil {
		return fail("json.Marshal", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(reqJSONByte))
	if err != nil {
		return fail("http.NewRequestWithContext", err)
	}
	httpReq.Header.Add(header.ContentType, header.AppJSON)

	httpRes, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return fail("http.DefaultClient.Do", err)
	}
	defer func() {
		err := httpRes.Body.Close()
		if err != nil {
			logrus.Warnf("http.Response.Body.Close: %v", err)
		}
	}()

	httpResBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return fail("io.ReadAll", err)
	}

	if httpRes.StatusCode != http.StatusOK {
		resErr := controllerHTTP.ResError{}
		err := json.Unmarshal(httpResBody, &resErr)
		if err != nil {
			return fail("json.Unmarshal", err)
		}
		return ResRegisterUser{}, errors.New(resErr.Error)
	}

	resBody := controllerHTTP.ResRegisterUser{}

	err = json.Unmarshal(httpResBody, &resBody)
	if err != nil {
		return fail("json.Unmarshal", err)
	}

	res := ResRegisterUser{
		UserID: resBody.Data.UserID,
	}

	return res, nil
}
