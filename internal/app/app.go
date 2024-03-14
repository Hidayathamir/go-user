// Package app contains the application starter.
package app

import (
	"flag"
	"path/filepath"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/controller/http"
	"github.com/Hidayathamir/go-user/internal/usecase/repo/db"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

// Run application.
func Run() {
	arg := parseArgs()

	cfg := initConfig(arg.isLoadEnv)

	handleCommandLineArgsMigrate(cfg, arg.isIncludeMigrate)

	db, err := db.NewPostgresPoolConnection(cfg)
	if err != nil {
		logrus.Fatalf("db.NewPostgresPoolConnection: %v", err)
	}

	err = http.RunServer(cfg, db)
	if err != nil {
		logrus.Fatalf("http.RunServer: %v", err)
	}
}

func initConfig(isLoadEnv bool) config.Config {
	var cfgLoader config.Loader
	if isLoadEnv {
		cfgLoader = &config.EnvLoader{YAMLPath: "./config/config.yml"}
	} else {
		cfgLoader = &config.YamlLoader{Path: "./config/config.yml"}
	}

	cfg, err := config.Init(cfgLoader)
	if err != nil {
		logrus.Fatalf("config.Init: %v", err)
	}

	return cfg
}

type arg struct {
	isIncludeMigrate bool
	isLoadEnv        bool
}

func parseArgs() arg {
	a := arg{}

	flag.BoolVar(&a.isIncludeMigrate, "include-migrate", false, "is include migrate, if true will do migrate before run app, default false.")
	flag.BoolVar(&a.isLoadEnv, "load-env", false, "is load env var, if true load env var and override config, default false.")

	flag.Usage = cleanenv.FUsage(flag.CommandLine.Output(), &config.Config{}, nil, flag.Usage)

	flag.Parse()

	return a
}

// handleCommandLineArgsMigrate do db migration then exit if args migrate exists.
func handleCommandLineArgsMigrate(cfg config.Config, isIncludeMigrate bool) {
	if isIncludeMigrate {
		schemaMigrationPath := filepath.Join("internal", "usecase", "repo", "db", "schema_migration")
		err := db.MigrateUp(cfg, schemaMigrationPath)
		if err != nil {
			logrus.Fatalf("db.MigrateUp: %v", err)
		}
	}
}
