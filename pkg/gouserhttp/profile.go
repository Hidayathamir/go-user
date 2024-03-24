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
	"github.com/Hidayathamir/go-user/internal/pkg/header"
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/sirupsen/logrus"
)

// API path list.
var (
	APIProfileGetProfileByUsername = func(username string) string {
		return "/api/v1/users/" + username
	}
	APIProfileUsers = "/api/v1/users"
)

// IProfileClient -.
type IProfileClient interface {
	GetProfileByUsername(ctx context.Context, req usecase.ReqGetProfileByUsername) (usecase.ResGetProfileByUsername, error)
	UpdateProfileByUserID(ctx context.Context, req usecase.ReqUpdateProfileByUserID) error
}

// ProfileClient -.
type ProfileClient struct {
	// BaseURL eg. http://localhost:8080.
	BaseURL string
}

var _ IProfileClient = &ProfileClient{}

// NewProfileClient -.
func NewProfileClient(baseURL string) *ProfileClient {
	return &ProfileClient{
		BaseURL: baseURL,
	}
}

// GetProfileByUsername implements IProfileClient.
func (p *ProfileClient) GetProfileByUsername(ctx context.Context, req usecase.ReqGetProfileByUsername) (usecase.ResGetProfileByUsername, error) {
	url := p.BaseURL + APIProfileGetProfileByUsername(req.Username)

	fail := func(msg string, err error) (usecase.ResGetProfileByUsername, error) {
		return usecase.ResGetProfileByUsername{}, fmt.Errorf(msg+": %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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
		return fail("http.Response.StatusCode != http.StatusOk", errors.New(resErr.Error))
	}

	res := controllerHTTP.ResGetProfileByUsername{}

	err = json.Unmarshal(httpResBody, &res)
	if err != nil {
		return fail("json.Unmarshal", err)
	}

	return res.Data, nil
}

// UpdateProfileByUserID implements IProfileClient.
func (p *ProfileClient) UpdateProfileByUserID(ctx context.Context, req usecase.ReqUpdateProfileByUserID) error {
	url := p.BaseURL + APIProfileUsers

	reqJSONByte, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(reqJSONByte))
	if err != nil {
		return fmt.Errorf("http.NewRequestWithContext: %w", err)
	}
	httpReq.Header.Add(header.ContentType, header.AppJSON)
	httpReq.Header.Add(header.Authorization, req.UserJWT)

	httpRes, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("http.DefaultClient.Do: %w", err)
	}
	defer func() {
		err := httpRes.Body.Close()
		if err != nil {
			logrus.Warnf("http.Response.Body.Close: %v", err)
		}
	}()

	httpResBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return fmt.Errorf("io.ReadAll: %w", err)
	}

	if httpRes.StatusCode != http.StatusOK {
		resErr := controllerHTTP.ResError{}
		err := json.Unmarshal(httpResBody, &resErr)
		if err != nil {
			return fmt.Errorf("json.Unmarshal: %w", err)
		}
		return fmt.Errorf("http.Response.StatusCode != http.StatusOk: %w", errors.New(resErr.Error))
	}

	return nil
}
