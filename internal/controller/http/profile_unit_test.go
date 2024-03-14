package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/dto"
	"github.com/Hidayathamir/go-user/internal/pkg/jutil"
	"github.com/Hidayathamir/go-user/internal/usecase/mockusecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUnitProfileGetProfileByUsername(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	type fields struct {
		cfg            config.Config
		usecaseProfile *mockusecase.MockIProfile
	}
	type args struct {
		params    gin.Params
		reqHeader gin.H
		reqBody   gin.H
	}
	tests := []struct {
		name     string
		mock     func(f fields)
		args     args
		wantCode int
		wantBody baseResponse
	}{
		{
			name: "get profile success",
			mock: func(f fields) {
				f.usecaseProfile.
					EXPECT().
					GetProfileByUsername(gomock.Any(), "hidayat").
					Return(
						dto.ResGetProfileByUsername{
							ID:        23,
							Username:  "hidayat",
							CreatedAt: time.Time{},
							UpdatedAt: time.Time{},
						},
						nil,
					)
			},
			args: args{
				params: gin.Params{
					{
						Key:   "username",
						Value: "hidayat",
					},
				},
				reqHeader: gin.H{},
				reqBody:   gin.H{},
			},
			wantCode: http.StatusOK,
			wantBody: baseResponse{
				Data: dto.ResGetProfileByUsername{
					ID:        23,
					Username:  "hidayat",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				},
				Error: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				cfg:            config.Config{},
				usecaseProfile: mockusecase.NewMockIProfile(ctrl),
			}
			tt.mock(f)

			p := &Profile{
				cfg:            f.cfg,
				usecaseProfile: f.usecaseProfile,
			}

			rr := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rr)
			reqBody, _ := json.Marshal(tt.args.reqBody)
			req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(reqBody))
			for k, v := range tt.args.reqHeader {
				req.Header.Set(k, fmt.Sprintf("%v", v))
			}
			ctx.Request = req
			ctx.Params = append(ctx.Params, tt.args.params...)

			p.getProfileByUsername(ctx)

			assert.Equal(t, tt.wantCode, rr.Code)
			assert.Equal(t, jutil.ToJSONString(tt.wantBody), rr.Body.String())
		})
	}
}
