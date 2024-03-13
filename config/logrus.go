package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func initLogrusConfig(cfg Config) error {
	logrusLevel, err := logrus.ParseLevel(string(cfg.Logger.LogLevel))
	if err != nil {
		return fmt.Errorf("logrus.ParseLevel: %w", err)
	}

	logrus.SetLevel(logrusLevel)

	if cfg.App.Environment == envProd {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	return nil
}
