package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	writeResponse(c, http.StatusOK, "ping success", nil)
}
