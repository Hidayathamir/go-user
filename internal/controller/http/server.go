// Package http contains func related to http as controller layer.
package http

import (
	"fmt"

	"github.com/Hidayathamir/go-user/config"
	"github.com/gin-gonic/gin"
)

// RunServer run http server.
func RunServer() error {
	ginEngine := gin.New()

	registerRouter(ginEngine)

	err := ginEngine.Run(fmt.Sprintf(":%d", config.HTTP.Port))
	if err != nil {
		return fmt.Errorf("gin.Engine.Run: %w", err)
	}

	return nil
}
