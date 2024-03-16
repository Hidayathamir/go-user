package http

import (
	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/db"
	"github.com/gin-gonic/gin"
)

// This file contains all available routers. It can be useful when you want to
// search for the API you want to debug. Think of it like an index in a dictionary.

func registerRouter(cfg config.Config, ginEngine *gin.Engine, db *db.Postgres) {
	ginEngine.GET("ping", ping)

	registerRouterV1(cfg, ginEngine.Group("api/v1"), db)
}

func registerRouterV1(cfg config.Config, routerV1 *gin.RouterGroup, db *db.Postgres) {
	cAuth := injectionAuth(cfg, db)
	cProfile := injectionProfile(cfg, db)

	authGroup := routerV1.Group("auth")
	{
		authGroup.POST("login", cAuth.loginUser)
		authGroup.POST("register", cAuth.registerUser)
	}

	userGroup := routerV1.Group("users")
	{
		userGroup.GET(":username", cProfile.getProfileByUsername)
		userGroup.PUT("", cProfile.updateProfileByUserID)
	}
}
