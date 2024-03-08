package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func initLogrusConfig() error {
	logrusLevel, err := logrus.ParseLevel(string(Logger.LogLevel))
	if err != nil {
		return fmt.Errorf("logrus.ParseLevel: %w", err)
	}

	logrus.SetLevel(logrusLevel)

	if App.Environment == envProd {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	return nil
}
