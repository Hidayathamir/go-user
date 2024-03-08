package http

import (
	"github.com/gin-gonic/gin"
)

// This file contains all router available. It can be useful when you want to
// search API you want to debug.

func registerRouter(ginEngine *gin.Engine) {
	ginEngine.GET("ping", ping)

	registerRouterV1(ginEngine.Group("api/v1"))
}

func registerRouterV1(routerV1 *gin.RouterGroup) {
	cAuth := newControllerAuth()
	cProfile := newControllerProfile()

	authGroup := routerV1.Group("auth")
	{
		authGroup.POST("login", cAuth.login)
		authGroup.POST("register", cAuth.register)
	}

	userGroup := routerV1.Group("users")
	{
		userGroup.GET(":username", cProfile.getProfile)
		userGroup.PUT(":username", cProfile.updateProfile)
	}
}
