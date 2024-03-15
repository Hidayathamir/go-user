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
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUnitAuthLoginUser(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	t.Run("call usecase LoginUser success should return success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecaseAuth := mockusecase.NewMockIAuth(ctrl)

		a := &Auth{
			cfg:         config.Config{},
			usecaseAuth: usecaseAuth,
		}

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		reqBody, err := json.Marshal(dto.ReqLoginUser{
			Username: "hidayat",
			Password: "mypassword",
		})
		require.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		ctx.Request = req

		usecaseAuth.EXPECT().LoginUser(gomock.Any(), dto.ReqLoginUser{
			Username: "hidayat",
			Password: "mypassword",
		}).Return("Bearer dummyUserJWT", nil)

		a.loginUser(ctx)

		assert.Equal(t, http.StatusOK, rr.Code)
		resBody := resLoginUserSuccess{}
		require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resBody))
		assert.NotEmpty(t, resBody.Data)
		assert.Contains(t, resBody.Data, "dummyUserJWT")
		assert.Nil(t, resBody.Error)
	})
	t.Run("call usecase LoginUser error should return error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecaseAuth := mockusecase.NewMockIAuth(ctrl)

		a := &Auth{
			cfg:         config.Config{},
			usecaseAuth: usecaseAuth,
		}

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		reqBody, err := json.Marshal(dto.ReqLoginUser{
			Username: "hidayat",
			Password: "mypassword",
		})
		require.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		ctx.Request = req

		usecaseAuth.EXPECT().LoginUser(gomock.Any(), dto.ReqLoginUser{
			Username: "hidayat",
			Password: "mypassword",
		}).Return("", assert.AnError)

		a.loginUser(ctx)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		resBody := resError{}
		require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resBody))
		assert.Nil(t, resBody.Data)
		assert.NotNil(t, resBody.Error)
		assert.Contains(t, resBody.Error, assert.AnError.Error())
	})
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
