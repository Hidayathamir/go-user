package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/dto"
	"github.com/Hidayathamir/go-user/internal/pkg/header"
	"github.com/Hidayathamir/go-user/internal/usecase/mockusecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUnitProfileGetProfileByUsername(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	t.Run("call usecase GetProfileByUsername success should return success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecaseProfile := mockusecase.NewMockIProfile(ctrl)

		p := &Profile{
			cfg:            config.Config{},
			usecaseProfile: usecaseProfile,
		}

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		ctx.Request = req
		ctx.Params = append(ctx.Params, gin.Param{
			Key:   "username",
			Value: "hidayat",
		})

		user := dto.ResGetProfileByUsername{
			ID:        23,
			Username:  "hidayat",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
		usecaseProfile.EXPECT().
			GetProfileByUsername(gomock.Any(), "hidayat").
			Return(user, nil)

		p.getProfileByUsername(ctx)

		assert.Equal(t, http.StatusOK, rr.Code)
		resBody := resGetProfileByUsernameSuccess{}
		require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resBody))
		assert.NotEmpty(t, resBody.Data)
		assert.Equal(t, user, resBody.Data)
		assert.Nil(t, resBody.Error)
	})
	t.Run("call usecase GetProfileByUsername error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecaseProfile := mockusecase.NewMockIProfile(ctrl)

		p := &Profile{
			cfg:            config.Config{},
			usecaseProfile: usecaseProfile,
		}

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		ctx.Request = req
		ctx.Params = append(ctx.Params, gin.Param{
			Key:   "username",
			Value: "hidayat",
		})

		usecaseProfile.EXPECT().
			GetProfileByUsername(gomock.Any(), "hidayat").
			Return(dto.ResGetProfileByUsername{}, assert.AnError)

		p.getProfileByUsername(ctx)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		resBody := resError{}
		require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resBody))
		assert.Nil(t, resBody.Data)
		assert.NotEmpty(t, resBody.Error)
		assert.Contains(t, resBody.Error, assert.AnError.Error())
	})
}

func TestProfileUpdateProfileByUserID(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	t.Run("call usecase UpdateProfileByUserID success should return success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecaseProfile := mockusecase.NewMockIProfile(ctrl)

		p := &Profile{
			cfg:            config.Config{},
			usecaseProfile: usecaseProfile,
		}

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		reqBody, _ := json.Marshal(dto.ReqUpdateProfileByUserID{
			Password: "newpassword",
		})
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(reqBody))
		req.Header.Set(header.Authorization, "Bearer dummyUserJWT")
		ctx.Request = req

		usecaseProfile.EXPECT().
			UpdateProfileByUserID(gomock.Any(), dto.ReqUpdateProfileByUserID{
				UserJWT:  "Bearer dummyUserJWT",
				Password: "newpassword",
			}).Return(nil)

		p.updateProfileByUserID(ctx)

		assert.Equal(t, http.StatusOK, rr.Code)
		resBody := resUpdatePofileSuccess{}
		require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resBody))
		assert.NotEmpty(t, resBody.Data)
		assert.Equal(t, "ok", resBody.Data)
		assert.Nil(t, resBody.Error)
	})
	t.Run("call usecase UpdateProfileByUserID error should return error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecaseProfile := mockusecase.NewMockIProfile(ctrl)

		p := &Profile{
			cfg:            config.Config{},
			usecaseProfile: usecaseProfile,
		}

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		reqBody, _ := json.Marshal(dto.ReqUpdateProfileByUserID{
			Password: "newpassword",
		})
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(reqBody))
		req.Header.Set(header.Authorization, "Bearer dummyUserJWT")
		ctx.Request = req

		usecaseProfile.EXPECT().
			UpdateProfileByUserID(gomock.Any(), dto.ReqUpdateProfileByUserID{
				UserJWT:  "Bearer dummyUserJWT",
				Password: "newpassword",
			}).Return(assert.AnError)

		p.updateProfileByUserID(ctx)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		resBody := resError{}
		require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resBody))
		assert.Nil(t, resBody.Data)
		assert.NotEmpty(t, resBody.Error)
		assert.Contains(t, resBody.Error, assert.AnError.Error())
	})
}
