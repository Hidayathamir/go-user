package app

import (
	"path/filepath"

	"github.com/Hidayathamir/go-user/config"
	"github.com/sirupsen/logrus"
)

func initConfig(arg cliArg) config.Config {
	yamlPath := filepath.Join("config", "config.yml")

	var cfgLoader config.Loader
	if arg.isLoadEnv {
		cfgLoader = &config.EnvLoader{YAMLPath: yamlPath}
	} else {
		cfgLoader = &config.YamlLoader{Path: yamlPath}
	}

	cfg, err := config.Init(cfgLoader)
	if err != nil {
		logrus.Fatalf("config.Init: %v", err)
	}

	return cfg
}
