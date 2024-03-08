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
func RunServer(db *db.Postgres) error {
	ginEngine := gin.New()

	registerRouter(ginEngine, db)

	addr := net.JoinHostPort(config.HTTP.Host, strconv.Itoa(config.HTTP.Port))
	err := ginEngine.Run(addr)
	if err != nil {
		return fmt.Errorf("gin.Engine.Run: %w", err)
	}

	return nil
}
