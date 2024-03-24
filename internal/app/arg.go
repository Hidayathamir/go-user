package app

import (
	"flag"
	"path/filepath"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/internal/repo/db"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

type cliArg struct {
	isIncludeMigrate bool
	isLoadEnv        bool
}

func parseCLIArgs() cliArg {
	arg := cliArg{}

	flag.BoolVar(&arg.isIncludeMigrate, "include-migrate", false, "is include migrate, if true will do migrate before run app, default false.")
	flag.BoolVar(&arg.isLoadEnv, "load-env", false, "is load env var, if true load env var and override config, default false.")

	flag.Usage = cleanenv.FUsage(flag.CommandLine.Output(), &config.Config{}, nil, flag.Usage)

	flag.Parse()

	return arg
}

// handleCommandLineArgsMigrate do db migration then exit if args migrate exists.
func handleCommandLineArgsMigrate(cfg config.Config, arg cliArg) {
	if arg.isIncludeMigrate {
		schemaMigrationPath := filepath.Join("internal", "repo", "db", "schema_migration")
		err := db.MigrateUp(cfg, schemaMigrationPath)
		if err != nil {
			logrus.Fatalf("db.MigrateUp: %v", err)
		}
	}
}
