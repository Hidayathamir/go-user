package grpc

import (
	"fmt"
	"net"
	"strconv"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/repo/db"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// RunServer run grpc server.
func RunServer(cfg config.Config, db *db.Postgres) error {
	grpcServer := grpc.NewServer()

	registerServer(cfg, grpcServer, db)

	addr := net.JoinHostPort(cfg.GRPC.Host, strconv.Itoa(cfg.GRPC.Port))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	logrus.WithField("address", addr).Info("run grpc server")
	err = grpcServer.Serve(lis)
	if err != nil {
		return fmt.Errorf("grpc.Server.Serve: %w", err)
	}

	return nil
}
