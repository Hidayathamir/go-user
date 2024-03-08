// Package http contains func related to http as controller layer.
package http

import (
	"fmt"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/usecase/repo/db"
	"github.com/gin-gonic/gin"
)

// RunServer run http server.
func RunServer(db *db.Postgres) error {
	ginEngine := gin.New()

	registerRouter(ginEngine, db)

	err := ginEngine.Run(fmt.Sprintf(":%d", config.HTTP.Port))
	if err != nil {
		return fmt.Errorf("gin.Engine.Run: %w", err)
	}

	return nil
}
