package http

import (
	"fmt"
	"net/http"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/dto"
	"github.com/Hidayathamir/go-user/internal/pkg/header"
	"github.com/Hidayathamir/go-user/internal/usecase"
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
	user, err := p.usecaseProfile.GetProfileByUsername(c, c.Param("username"))
	if err != nil {
		err := fmt.Errorf("Profile.usecaseProfile.GetProfileByUsername: %w", err)
		writeResponse(c, http.StatusBadRequest, nil, err)
		return
	}
	writeResponse(c, http.StatusOK, user, nil)
}

func (p *Profile) updateProfileByUserID(c *gin.Context) {
	req := dto.ReqUpdateProfileByUserID{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		err := fmt.Errorf("gin.Context.ShouldBindJSON: %w", err)
		writeResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	req.UserJWT = c.GetHeader(header.Authorization)

	err = p.usecaseProfile.UpdateProfileByUserID(c, req)
	if err != nil {
		err := fmt.Errorf("Profile.usecaseProfile.UpdateProfileByUserID: %w", err)
		writeResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	writeResponse(c, http.StatusOK, "ok", nil)
}
