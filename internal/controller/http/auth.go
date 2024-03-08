package http

import (
	"github.com/Hidayathamir/go-user/internal/usecase"
	"github.com/gin-gonic/gin"
)

// Auth is controller HTTP for authentication related.
type Auth struct {
	usecaseAuth usecase.IAuth
}

func newAuth(usecaseAuth usecase.IAuth) *Auth {
	return &Auth{
		usecaseAuth: usecaseAuth,
	}
}

func (*Auth) login(*gin.Context) { // TODO:
}

func (*Auth) register(*gin.Context) { // TODO:
}
