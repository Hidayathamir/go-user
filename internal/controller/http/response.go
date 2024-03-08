package http

import "github.com/gin-gonic/gin"

// writeResponse should be used to write the response JSON body, ensuring that
// the response body maintains consistent content.
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
