// Package app contains the application starter.
package app

import (
	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/controller/grpc"
	"github.com/Hidayathamir/go-user/internal/controller/http"
	"github.com/Hidayathamir/go-user/internal/repo/db"
	"github.com/sirupsen/logrus"
)

// Run application.
func Run() {
	arg := parseCLIArgs()

	cfg := initConfig(arg)

	handleCommandLineArgsMigrate(cfg, arg)

	db := newDBPostgres(cfg)

	go runGRPCServer(cfg, db)

	runHTTPServer(cfg, db)
}

func newDBPostgres(cfg config.Config) *db.Postgres {
	db, err := db.NewPGPoolConn(cfg)
	if err != nil {
		logrus.Fatalf("db.NewPostgresPoolConnection: %v", err)
	}
	return db
}

func runGRPCServer(cfg config.Config, db *db.Postgres) {
	err := grpc.RunServer(cfg, db)
	if err != nil {
		logrus.Fatalf("grpc.RunServer: %v", err)
	}
}

func runHTTPServer(cfg config.Config, db *db.Postgres) {
	err := http.RunServer(cfg, db)
	if err != nil {
		logrus.Fatalf("http.RunServer: %v", err)
	}
}
