package app

import (
	"flag"

	"github.com/Hidayathamir/go-user/config"
	"github.com/ilyakaznacheev/cleanenv"
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
