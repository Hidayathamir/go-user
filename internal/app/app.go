// Package app contains the application starter.
package app

import (
	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/controller/http"
	"github.com/Hidayathamir/go-user/internal/usecase/repo/db"
	"github.com/sirupsen/logrus"
)

// Run application.
func Run() {
	err := config.Init()
	if err != nil {
		logrus.Fatalf("config.Init: %v", err)
	}

	db, err := db.NewPostgresPoolConnection()
	if err != nil {
		logrus.Fatalf("db.NewPostgresPoolConnection: %v", err)
	}

	err = http.RunServer(db)
	if err != nil {
		logrus.Fatalf("http.RunServer: %v", err)
	}
}
