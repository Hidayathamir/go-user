package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/dto"
	"github.com/Hidayathamir/go-user/internal/pkg/jutil"
	"github.com/Hidayathamir/go-user/internal/usecase/mockusecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUnitAuthLoginUser(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	type fields struct {
		cfg         config.Config
		usecaseAuth *mockusecase.MockIAuth
	}
	type args struct {
		reqBody dto.ReqLoginUser
	}
	tests := []struct {
		name     string
		mock     func(f fields)
		args     args
		wantCode int
		wantBody baseResponse
	}{
		{
			name: "login user success",
			mock: func(f fields) {
				f.usecaseAuth.
					EXPECT().
					LoginUser(gomock.Any(), dto.ReqLoginUser{
						Username: "hidayat",
						Password: "mypassword",
					}).
					Return("Bearer dummyUserJWT", nil)
			},
			args: args{
				reqBody: dto.ReqLoginUser{
					Username: "hidayat",
					Password: "mypassword",
				},
			},
			wantCode: http.StatusOK,
			wantBody: setResponseBody("Bearer dummyUserJWT", nil),
		},
		{
			name: "login user error",
			mock: func(f fields) {
				f.usecaseAuth.
					EXPECT().
					LoginUser(gomock.Any(), dto.ReqLoginUser{
						Username: "hidayat",
						Password: "mypassword",
					}).
					Return("", assert.AnError)
			},
			args: args{
				reqBody: dto.ReqLoginUser{
					Username: "hidayat",
					Password: "mypassword",
				},
			},
			wantCode: http.StatusBadRequest,
			wantBody: setResponseBody(nil, fmt.Errorf("Auth.usecaseAuth.LoginUser: %w", assert.AnError)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				cfg:         config.Config{},
				usecaseAuth: mockusecase.NewMockIAuth(ctrl),
			}
			tt.mock(f)

			a := &Auth{
				cfg:         f.cfg,
				usecaseAuth: f.usecaseAuth,
			}

			rr := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rr)
			reqBody, _ := json.Marshal(tt.args.reqBody)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
			ctx.Request = req

			a.loginUser(ctx)

			assert.Equal(t, tt.wantCode, rr.Code)
			assert.Equal(t, jutil.ToJSONString(tt.wantBody), rr.Body.String())
		})
	}
}

func TestUnitAuthRegisterUser(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	type fields struct {
		cfg         config.Config
		usecaseAuth *mockusecase.MockIAuth
	}
	type args struct {
		reqBody dto.ReqRegisterUser
	}
	tests := []struct {
		name     string
		mock     func(f fields)
		args     args
		wantCode int
		wantBody baseResponse
	}{
		{
			name: "register user success",
			mock: func(f fields) {
				f.usecaseAuth.
					EXPECT().
					RegisterUser(gomock.Any(), dto.ReqRegisterUser{
						Username: "hidayat",
						Password: "mypassword",
					}).
					Return(int64(442), nil)
			},
			args: args{
				reqBody: dto.ReqRegisterUser{
					Username: "hidayat",
					Password: "mypassword",
				},
			},
			wantCode: http.StatusOK,
			wantBody: setResponseBody(int64(442), nil),
		},
		{
			name: "register user error",
			mock: func(f fields) {
				f.usecaseAuth.
					EXPECT().
					RegisterUser(gomock.Any(), dto.ReqRegisterUser{
						Username: "hidayat",
						Password: "mypassword",
					}).
					Return(int64(0), assert.AnError)
			},
			args: args{
				reqBody: dto.ReqRegisterUser{
					Username: "hidayat",
					Password: "mypassword",
				},
			},
			wantCode: http.StatusBadRequest,
			wantBody: setResponseBody(nil, fmt.Errorf("Auth.usecaseAuth.RegisterUser: %w", assert.AnError)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				cfg:         config.Config{},
				usecaseAuth: mockusecase.NewMockIAuth(ctrl),
			}
			tt.mock(f)

			a := &Auth{
				cfg:         f.cfg,
				usecaseAuth: f.usecaseAuth,
			}

			rr := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rr)
			reqBody, _ := json.Marshal(tt.args.reqBody)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
			ctx.Request = req

			a.registerUser(ctx)

			assert.Equal(t, tt.wantCode, rr.Code)
			assert.Equal(t, jutil.ToJSONString(tt.wantBody), rr.Body.String())
		})
	}
}
