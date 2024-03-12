package http

import (
	"github.com/Hidayathamir/go-user/internal/usecase/repo/db"
	"github.com/gin-gonic/gin"
)

// This file contains all available routers. It can be useful when you want to
// search for the API you want to debug. Think of it like an index in a dictionary.

func registerRouter(ginEngine *gin.Engine, db *db.Postgres) {
	ginEngine.GET("ping", ping)

	registerRouterV1(ginEngine.Group("api/v1"), db)
}

func registerRouterV1(routerV1 *gin.RouterGroup, db *db.Postgres) {
	cAuth := injectionAuth(db)
	cProfile := injectionProfile(db)

	authGroup := routerV1.Group("auth")
	{
		authGroup.POST("login", cAuth.loginUser)
		authGroup.POST("register", cAuth.registerUser)
	}

	userGroup := routerV1.Group("users")
	{
		userGroup.GET(":username", cProfile.getProfileByUsername)
		userGroup.PUT(":username", cProfile.updateProfileByUsername)
	}
}
