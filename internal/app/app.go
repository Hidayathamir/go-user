// Package app contains the application starter.
package app

import (
	"flag"
	"os"

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

	handleCommandLineArgsMigrate()

	db, err := db.NewPostgresPoolConnection()
	if err != nil {
		logrus.Fatalf("db.NewPostgresPoolConnection: %v", err)
	}

	err = http.RunServer(db)
	if err != nil {
		logrus.Fatalf("http.RunServer: %v", err)
	}
}

// handleCommandLineArgsMigrate do db migration then exit if args migrate exists.
func handleCommandLineArgsMigrate() {
	var isHasArgMigrate bool
	flag.BoolVar(&isHasArgMigrate, "migrate", false, "is do migrate, default false.")
	flag.Parse()

	if isHasArgMigrate {
		err := db.MigrateUp()
		if err != nil {
			logrus.Fatalf("db.MigrateUp: %v", err)
		}

		os.Exit(0)
	}
}
