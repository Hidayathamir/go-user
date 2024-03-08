package config

import "github.com/gin-gonic/gin"

func initGinConfig() {
	if App.Environment == envProd {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
