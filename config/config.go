// Package config contains configuration.
package config

import (
	"fmt"
)

// Init initiate configurations either from config yml or env var.
func Init(cfgLoader Loader) (Config, error) {
	cfg := Config{}

	err := cfgLoader.loadConfig(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("ConfigLoader.loadConfig: %w", err)
	}

	err = cfg.validate()
	if err != nil {
		return Config{}, fmt.Errorf("config.Validate: %w", err)
	}

	err = initLogrusConfig(cfg)
	if err != nil {
		return Config{}, fmt.Errorf("initLogrusConfig: %w", err)
	}

	initGinConfig(cfg)

	return cfg, nil
}

// Config holds all config.
type Config struct {
	App    App    `yaml:"app"      env-required:"true"`
	HTTP   HTTP   `yaml:"http"     env-required:"true"`
	Logger logger `yaml:"logger"   env-required:"true"`
	PG     PG     `yaml:"postgres" env-required:"true"`
	JWT    JWT    `yaml:"jwt"      env-required:"true"`
}

func (c *Config) validate() error {
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

// App hold app configuration.
type App struct {
	Name        string `yaml:"name"        env-required:"true" env:"APP_NAME"`
	Version     string `yaml:"version"     env-required:"true" env:"APP_VERSION"`
	Environment env    `yaml:"environment" env-required:"true" env:"APP_ENVIRONMENT"`
}

// HTTP hold HTTP configuration.
type HTTP struct {
	Host string `yaml:"host" env-required:"true" env:"HTTP_HOST"`
	Port int    `yaml:"port" env-required:"true" env:"HTTP_PORT"`
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
	LogLevel logLevel `yaml:"log_level" env-required:"true" env:"LOGGER_LOG_LEVEL"`
}

// PG hold postgres configuration.
type PG struct {
	PoolMax  int    `yaml:"pool_max" env-required:"true" env:"POSTGRES_POOL_MAX"`
	Username string `yaml:"username" env-required:"true" env:"POSTGRES_USERNAME"`
	Password string `yaml:"password" env-required:"true" env:"POSTGRES_PASSWORD"`
	Host     string `yaml:"host"     env-required:"true" env:"POSTGRES_HOST"`
	Port     int    `yaml:"port"     env-required:"true" env:"POSTGRES_PORT"`
	DBName   string `yaml:"db_name"  env-required:"true" env:"POSTGRES_DB_NAME"`
}

// JWT hold JWT configuration.
type JWT struct {
	ExpireHour int    `yaml:"expire_hour" env-required:"true" env:"JWT_EXPIRE_HOUR"`
	SignedKey  string `yaml:"signed_key"  env-required:"true" env:"JWT_SIGNED_KEY"`
}
