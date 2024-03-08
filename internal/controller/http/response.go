package http

import "github.com/gin-gonic/gin"

func writeResponse(c *gin.Context, code int, data any, err error) {
	if err != nil {
		c.JSON(code, baseResponse{Data: data, Error: err.Error()})
	} else {
		c.JSON(code, baseResponse{Data: data, Error: ""})
	}
}

type baseResponse struct {
	Data  any    `json:"data"`
	Error string `json:"error"`
}
