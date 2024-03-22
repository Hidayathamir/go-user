package grpc

import (
	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/db"
	"github.com/Hidayathamir/go-user/pkg/gousergrpc"
	"google.golang.org/grpc"
)

// This file contains all available servers.

func registerServer(cfg config.Config, grpcServer *grpc.Server, db *db.Postgres) {
	gousergrpc.RegisterPingServer(grpcServer, &Ping{})

	cAuth := injectionAuth(cfg, db)
	cProfile := injectionProfile(cfg, db)

	gousergrpc.RegisterAuthServer(grpcServer, cAuth)
	gousergrpc.RegisterProfileServer(grpcServer, cProfile)
}
