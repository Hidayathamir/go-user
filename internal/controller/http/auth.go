package http

import (
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/gin-gonic/gin"
)

// Auth -.
type Auth struct {
	usecaseAuth usecase.IAuth
}

// newAuth -.
func newAuth(usecaseAuth usecase.IAuth) *Auth {
	return &Auth{
		usecaseAuth: usecaseAuth,
	}
}

func (*Auth) login(*gin.Context) { // TODO:
}

func (*Auth) register(*gin.Context) { // TODO:
}
