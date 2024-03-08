package http

import (
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/gin-gonic/gin"
)

// Profile is controller HTTP for profile related.
type Profile struct {
	usecaseProfile usecase.IProfile
}

func newProfile(usecaseProfile usecase.IProfile) *Profile {
	return &Profile{
		usecaseProfile: usecaseProfile,
	}
}

func (*Profile) getProfile(*gin.Context) { // TODO:
}

func (*Profile) updateProfile(*gin.Context) { // TODO:
}
