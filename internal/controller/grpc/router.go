package grpc

import (
	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/db"
	"github.com/Hidayathamir/go-user/pkg/gouser/grpc/pb"
	"google.golang.org/grpc"
)

// This file contains all available servers.

func registerServer(cfg config.Config, grpcServer *grpc.Server, db *db.Postgres) {
	pb.RegisterPingServer(grpcServer, &Ping{})

	cAuth := injectionAuth(cfg, db)
	cProfile := injectionProfile(cfg, db)

	pb.RegisterAuthServer(grpcServer, cAuth)
	pb.RegisterProfileServer(grpcServer, cProfile)
}
