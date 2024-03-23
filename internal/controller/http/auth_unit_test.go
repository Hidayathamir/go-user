package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/dto"
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
		t.Parallel()

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
		}).Return(dto.ResLoginUser{UserJWT: "Bearer dummyUserJWT"}, nil)

		a.loginUser(ctx)

		assert.Equal(t, http.StatusOK, rr.Code)
		resBody := ResLoginUser{}
		require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resBody))
		assert.NotEmpty(t, resBody.Data)
		assert.Contains(t, resBody.Data.UserJWT, "dummyUserJWT")
		assert.Nil(t, resBody.Error)
	})
	t.Run("call usecase LoginUser error should return error", func(t *testing.T) {
		t.Parallel()

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
		}).Return(dto.ResLoginUser{}, assert.AnError)

		a.loginUser(ctx)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		resBody := ResError{}
		require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resBody))
		assert.Nil(t, resBody.Data)
		assert.NotEmpty(t, resBody.Error)
		assert.Contains(t, resBody.Error, assert.AnError.Error())
	})
}

func TestUnitAuthRegisterUser(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	t.Run("call usecase RegisterUser success should return success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		usecaseAuth := mockusecase.NewMockIAuth(ctrl)

		a := &Auth{
			cfg:         config.Config{},
			usecaseAuth: usecaseAuth,
		}

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		reqBody, err := json.Marshal(dto.ReqRegisterUser{
			Username: "hidayat",
			Password: "mypassword",
		})
		require.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
		ctx.Request = req

		usecaseAuth.EXPECT().RegisterUser(gomock.Any(), dto.ReqRegisterUser{
			Username: "hidayat",
			Password: "mypassword",
		}).Return(dto.ResRegisterUser{UserID: 323}, nil)

		a.registerUser(ctx)

		assert.Equal(t, http.StatusOK, rr.Code)
		resBody := ResRegisterUser{}
		require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resBody))
		assert.NotEmpty(t, resBody.Data)
		assert.Equal(t, int64(323), resBody.Data.UserID)
		assert.Nil(t, resBody.Error)
	})
	t.Run("call usecase RegisterUser error should return error", func(t *testing.T) {
		t.Parallel()

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

		usecaseAuth.EXPECT().RegisterUser(gomock.Any(), dto.ReqRegisterUser{
			Username: "hidayat",
			Password: "mypassword",
		}).Return(dto.ResRegisterUser{UserID: 0}, assert.AnError)

		a.registerUser(ctx)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		resBody := ResError{}
		require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resBody))
		assert.Nil(t, resBody.Data)
		assert.NotEmpty(t, resBody.Error)
		assert.Contains(t, resBody.Error, assert.AnError.Error())
	})
}
