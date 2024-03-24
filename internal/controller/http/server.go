package http

import (
	"fmt"
	"net"
	"strconv"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/repo/db"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// RunServer run http server.
func RunServer(cfg config.Config, db *db.Postgres) error {
	ginEngine := gin.New()

	registerRouter(cfg, ginEngine, db)

	addr := net.JoinHostPort(cfg.HTTP.Host, strconv.Itoa(cfg.HTTP.Port))
	logrus.WithField("address", addr).Info("run http server")
	err := ginEngine.Run(addr)
	if err != nil {
		return fmt.Errorf("gin.Engine.Run: %w", err)
	}

	return nil
}
