package http

import (
	"fmt"
	"net/http"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/Hidayathamir/go-user/pkg/header"
	"github.com/gin-gonic/gin"
)

// Profile is controller HTTP for profile related.
type Profile struct {
	cfg            config.Config
	usecaseProfile usecase.IProfile
}

func newProfile(cfg config.Config, usecaseProfile usecase.IProfile) *Profile {
	return &Profile{
		cfg:            cfg,
		usecaseProfile: usecaseProfile,
	}
}

func (p *Profile) getProfileByUsername(c *gin.Context) {
	req := usecase.ReqGetProfileByUsername{Username: c.Param("username")}
	user, err := p.usecaseProfile.GetProfileByUsername(c, req)
	if err != nil {
		err := fmt.Errorf("Profile.usecaseProfile.GetProfileByUsername: %w", err)
		c.JSON(http.StatusBadRequest, ResError{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ResGetProfileByUsername{Data: user})
}

func (p *Profile) updateProfileByUserID(c *gin.Context) {
	req := usecase.ReqUpdateProfileByUserID{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		err := fmt.Errorf("gin.Context.ShouldBindJSON: %w", err)
		c.JSON(http.StatusBadRequest, ResError{Error: err.Error()})
		return
	}

	req.UserJWT = c.GetHeader(header.Authorization)

	err = p.usecaseProfile.UpdateProfileByUserID(c, req)
	if err != nil {
		err := fmt.Errorf("Profile.usecaseProfile.UpdateProfileByUserID: %w", err)
		c.JSON(http.StatusBadRequest, ResError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResString{Data: "ok"})
}
