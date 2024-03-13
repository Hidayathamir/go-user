package http

import (
	"fmt"
	"net/http"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/dto"
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/gin-gonic/gin"
)

// Auth is controller HTTP for authentication related.
type Auth struct {
	cfg         config.Config
	usecaseAuth usecase.IAuth
}

func newAuth(cfg config.Config, usecaseAuth usecase.IAuth) *Auth {
	return &Auth{
		cfg:         cfg,
		usecaseAuth: usecaseAuth,
	}
}

func (a *Auth) loginUser(c *gin.Context) {
	req := dto.ReqLoginUser{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		err := fmt.Errorf("gin.Context.ShouldBindJSON: %w", err)
		writeResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	userJWT, err := a.usecaseAuth.LoginUser(c, req)
	if err != nil {
		err := fmt.Errorf("Auth.usecaseAuth.LoginUser: %w", err)
		writeResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	writeResponse(c, http.StatusOK, userJWT, nil)
}

func (a *Auth) registerUser(c *gin.Context) {
	req := dto.ReqRegisterUser{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		err := fmt.Errorf("gin.Context.ShouldBindJSON: %w", err)
		writeResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	userID, err := a.usecaseAuth.RegisterUser(c, req)
	if err != nil {
		err := fmt.Errorf("Auth.usecaseAuth.RegisterUser: %w", err)
		writeResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	writeResponse(c, http.StatusOK, userID, nil)
}
