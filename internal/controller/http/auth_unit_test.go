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

func TestUnitAuthRegisterUser(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	type fields struct {
		cfg         config.Config
		usecaseAuth *mockusecase.MockIAuth
	}
	type args struct {
		reqHeader gin.H
		reqBody   gin.H
	}
	tests := []struct {
		name     string
		mock     func(f fields)
		args     args
		wantCode int
		wantBody gin.H
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
				reqHeader: gin.H{},
				reqBody: gin.H{
					"username": "hidayat",
					"password": "mypassword",
				},
			},
			wantCode: http.StatusOK,
			wantBody: gin.H{
				"data":  int64(442),
				"error": nil,
			},
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
				reqHeader: gin.H{},
				reqBody: gin.H{
					"username": "hidayat",
					"password": "mypassword",
				},
			},
			wantCode: http.StatusBadRequest,
			wantBody: gin.H{
				"data":  nil,
				"error": fmt.Errorf("Auth.usecaseAuth.RegisterUser: %w", assert.AnError).Error(),
			},
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
			req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(reqBody))
			for k, v := range tt.args.reqHeader {
				req.Header.Set(k, fmt.Sprintf("%v", v))
			}
			ctx.Request = req

			a.registerUser(ctx)

			assert.Equal(t, tt.wantCode, rr.Code)
			assert.Equal(t, jutil.ToJSONString(tt.wantBody), rr.Body.String())
		})
	}
}
