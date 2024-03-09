// Package app contains the application starter.
package app

import (
	"flag"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/controller/http"
	"github.com/Hidayathamir/go-user/internal/entity/table"
	"github.com/Hidayathamir/go-user/internal/usecase/repo/db"
	"github.com/sirupsen/logrus"
)

// Run application.
func Run() {
	arg := parseArgs()

	err := config.Init(arg.isLoadEnv)
	if err != nil {
		logrus.Fatalf("config.Init: %v", err)
	}

	handleCommandLineArgsMigrate(arg.isIncludeMigrate)

	table.Init()

	db, err := db.NewPostgresPoolConnection()
	if err != nil {
		logrus.Fatalf("db.NewPostgresPoolConnection: %v", err)
	}

	err = http.RunServer(db)
	if err != nil {
		logrus.Fatalf("http.RunServer: %v", err)
	}
}

type arg struct {
	isIncludeMigrate bool
	isLoadEnv        bool
}

func parseArgs() arg {
	a := arg{}

	flag.BoolVar(&a.isIncludeMigrate, "include-migrate", false, "is include migrate, if true will do migrate before run app, default false.")
	flag.BoolVar(&a.isLoadEnv, "load-env", false, "is load env var, if true load env var and override config, default false.")

	flag.Parse()

	return a
}

// handleCommandLineArgsMigrate do db migration then exit if args migrate exists.
func handleCommandLineArgsMigrate(isIncludeMigrate bool) {
	if isIncludeMigrate {
		err := db.MigrateUp()
		if err != nil {
			logrus.Fatalf("db.MigrateUp: %v", err)
		}
	}
}
