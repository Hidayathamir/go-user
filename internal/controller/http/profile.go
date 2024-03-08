package http

import (
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/gin-gonic/gin"
)

// Profile -.
type Profile struct {
	usecaseProfile usecase.IProfile
}

// newProfile -.
func newProfile(usecaseProfile usecase.IProfile) *Profile {
	return &Profile{
		usecaseProfile: usecaseProfile,
	}
}

func (*Profile) getProfile(*gin.Context) { // TODO:
}

func (*Profile) updateProfile(*gin.Context) { // TODO:
}
