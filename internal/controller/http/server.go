// Package http contains func related to http as controller layer.
package http

import (
	"fmt"
	"net"
	"strconv"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/usecase/repo/db"
	"github.com/gin-gonic/gin"
)

// RunServer run http server.
func RunServer(cfg config.Config, db *db.Postgres) error {
	ginEngine := gin.New()

	registerRouter(cfg, ginEngine, db)

	addr := net.JoinHostPort(cfg.HTTP.Host, strconv.Itoa(cfg.HTTP.Port))
	err := ginEngine.Run(addr)
	if err != nil {
		return fmt.Errorf("gin.Engine.Run: %w", err)
	}

	return nil
}
