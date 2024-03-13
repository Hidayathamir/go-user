package config

import "github.com/gin-gonic/gin"

func initGinConfig(cfg Config) {
	if cfg.App.Environment == envProd {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
