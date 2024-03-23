package grpc

import (
	"context"
	"testing"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/dto"
	"github.com/Hidayathamir/go-user/internal/usecase/mockusecase"
	"github.com/Hidayathamir/go-user/pkg/gousergrpc"
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

		usecaseAuth.EXPECT().LoginUser(gomock.Any(), dto.ReqLoginUser{
			Username: "hidayat",
			Password: "mypassword",
		}).Return(dto.ResLoginUser{UserJWT: "Bearer dummyUserJWT"}, nil)

		req := &gousergrpc.ReqLoginUser{
			Username: "hidayat",
			Password: "mypassword",
		}

		res, err := a.LoginUser(context.Background(), req)

		assert.NotNil(t, res.GetUserJwt())
		assert.Contains(t, res.GetUserJwt(), "dummyUserJWT")
		require.NoError(t, err)
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

		usecaseAuth.EXPECT().LoginUser(gomock.Any(), dto.ReqLoginUser{
			Username: "hidayat",
			Password: "mypassword",
		}).Return(dto.ResLoginUser{}, assert.AnError)

		req := &gousergrpc.ReqLoginUser{
			Username: "hidayat",
			Password: "mypassword",
		}

		res, err := a.LoginUser(context.Background(), req)

		assert.Nil(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
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

		usecaseAuth.EXPECT().RegisterUser(gomock.Any(), dto.ReqRegisterUser{
			Username: "hidayat",
			Password: "mypassword",
		}).Return(dto.ResRegisterUser{UserID: 323}, nil)

		req := &gousergrpc.ReqRegisterUser{
			Username: "hidayat",
			Password: "mypassword",
		}

		res, err := a.RegisterUser(context.Background(), req)

		assert.Equal(t, int64(323), res.GetUserId())
		require.NoError(t, err)
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

		usecaseAuth.EXPECT().RegisterUser(gomock.Any(), dto.ReqRegisterUser{
			Username: "hidayat",
			Password: "mypassword",
		}).Return(dto.ResRegisterUser{UserID: 0}, assert.AnError)

		req := &gousergrpc.ReqRegisterUser{
			Username: "hidayat",
			Password: "mypassword",
		}

		res, err := a.RegisterUser(context.Background(), req)

		assert.Nil(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
}
