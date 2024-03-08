package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRouter(ginEngine *gin.Engine) {
	ginEngine.GET("ping", ping)
}

func ping(c *gin.Context) {
	writeResponse(c, http.StatusOK, "ping success", nil)
}
