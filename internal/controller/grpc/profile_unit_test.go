package grpc

import (
	"context"
	"testing"
	"time"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/Hidayathamir/go-user/internal/usecase/mockusecase"
	"github.com/Hidayathamir/go-user/pkg/gousergrpc"
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

		user := usecase.ResGetProfileByUsername{
			ID:        23,
			Username:  "hidayat",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}

		usecaseProfile.EXPECT().
			GetProfileByUsername(gomock.Any(), usecase.ReqGetProfileByUsername{Username: "hidayat"}).
			Return(user, nil)

		req := &gousergrpc.ReqGetProfileByUsername{
			Username: "hidayat",
		}

		res, err := p.GetProfileByUsername(context.Background(), req)

		assert.NotNil(t, res)
		assert.Equal(t, user.ID, res.GetId())
		assert.Equal(t, user.Username, res.GetUsername())
		assert.Equal(t, user.CreatedAt, res.GetCreatedAt().AsTime())
		assert.Equal(t, user.UpdatedAt, res.GetUpdatedAt().AsTime())
		require.NoError(t, err)
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

		user := usecase.ResGetProfileByUsername{}

		usecaseProfile.EXPECT().
			GetProfileByUsername(gomock.Any(), usecase.ReqGetProfileByUsername{Username: "hidayat"}).
			Return(user, assert.AnError)

		req := &gousergrpc.ReqGetProfileByUsername{
			Username: "hidayat",
		}

		res, err := p.GetProfileByUsername(context.Background(), req)

		assert.Nil(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
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

		usecaseProfile.EXPECT().
			UpdateProfileByUserID(gomock.Any(), usecase.ReqUpdateProfileByUserID{
				UserJWT:  "Bearer dummyUserJWT",
				Password: "newpassword",
			}).Return(nil)

		req := &gousergrpc.ReqUpdateProfileByUserID{
			UserJwt:  "Bearer dummyUserJWT",
			Password: "newpassword",
		}

		res, err := p.UpdateProfileByUserID(context.Background(), req)

		assert.NotNil(t, res)
		require.NoError(t, err)
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

		usecaseProfile.EXPECT().
			UpdateProfileByUserID(gomock.Any(), usecase.ReqUpdateProfileByUserID{
				UserJWT:  "Bearer dummyUserJWT",
				Password: "newpassword",
			}).Return(assert.AnError)

		req := &gousergrpc.ReqUpdateProfileByUserID{
			UserJwt:  "Bearer dummyUserJWT",
			Password: "newpassword",
		}

		res, err := p.UpdateProfileByUserID(context.Background(), req)

		assert.Nil(t, res)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
}
