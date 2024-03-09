// Package config contains configuration.
package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

// configuration list. Got panic? did you run Init?
var (
	App    *app
	HTTP   *http
	Logger *logger
	PG     *pg
	JWT    *jwt
)

// Init initiate configurations from `./config/config.yml` file.
func Init() error {
	if App != nil && HTTP != nil && Logger != nil && PG != nil && JWT != nil {
		logrus.Warn("config already initialized")
		return nil
	}

	cfg := config{}

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		return fmt.Errorf("cleanenv.ReadConfig: %w", err)
	}

	err = cfg.validate()
	if err != nil {
		return fmt.Errorf("config.Validate: %w", err)
	}

	App = &cfg.App
	HTTP = &cfg.HTTP
	Logger = &cfg.Logger
	PG = &cfg.PG
	JWT = &cfg.JWT

	err = initLogrusConfig()
	if err != nil {
		return fmt.Errorf("initLogrusConfig: %w", err)
	}

	initGinConfig()

	return nil
}

type config struct {
	App    app    `yaml:"app"      env-required:"true"`
	HTTP   http   `yaml:"http"     env-required:"true"`
	Logger logger `yaml:"logger"   env-required:"true"`
	PG     pg     `yaml:"postgres" env-required:"true"`
	JWT    jwt    `yaml:"jwt" env-required:"true"`
}

func (c *config) validate() error {
	err := c.App.Environment.validate()
	if err != nil {
		return fmt.Errorf("config.app.Environment.validate: %w", err)
	}

	err = c.Logger.LogLevel.validate()
	if err != nil {
		return fmt.Errorf("config.logger.LogLevel.validate: %w", err)
	}

	return nil
}

type env string

const (
	envDev  env = "dev"
	envProd env = "prod"
)

func (e env) validate() error {
	switch e {
	case envDev, envProd:
	default:
		return fmt.Errorf("unknown environment '%s'", e)
	}

	return nil
}

type app struct {
	Name        string `yaml:"name"        env-required:"true"`
	Version     string `yaml:"version"     env-required:"true"`
	Environment env    `yaml:"environment" env-required:"true"`
}

type http struct {
	Host string `yaml:"host" env-required:"true"`
	Port int    `yaml:"port" env-required:"true"`
}

type logLevel string

func (l logLevel) validate() error {
	switch l {
	case "panic", "fatal", "error", "warn", "warning", "info", "debug", "trace":
	default:
		return fmt.Errorf("unknown config logger log level '%s'", l)
	}

	return nil
}

type logger struct {
	LogLevel logLevel `yaml:"log_level" env-required:"true"`
}

type pg struct {
	PoolMax  int    `yaml:"pool_max" env-required:"true"`
	Username string `yaml:"username" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Host     string `yaml:"host"     env-required:"true"`
	Port     int    `yaml:"port"     env-required:"true"`
	DBName   string `yaml:"db_name"  env-required:"true"`
}

type jwt struct {
	ExpireHour int    `yaml:"expire_hour" env-required:"true"`
	SignedKey  string `yaml:"signed_key" env-required:"true"`
}
