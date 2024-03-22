// Package config contains the application configuration.
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
	App    App    `yaml:"app"      env-required:"true" env-prefix:"APP_"`
	HTTP   HTTP   `yaml:"http"     env-required:"true" env-prefix:"HTTP_"`
	GRPC   GRPC   `yaml:"grpc"     env-required:"true" env-prefix:"GRPC_"`
	Logger logger `yaml:"logger"   env-required:"true" env-prefix:"LOGGER_"`
	PG     PG     `yaml:"postgres" env-required:"true" env-prefix:"POSTGRES_"`
	JWT    JWT    `yaml:"jwt"      env-required:"true" env-prefix:"JWT_"`
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
	Name        string `yaml:"name"        env-required:"true" env:"NAME"        env-description:"app service name"`
	Version     string `yaml:"version"     env-required:"true" env:"VERSION"     env-description:"app service version"`
	Environment env    `yaml:"environment" env-required:"true" env:"ENVIRONMENT" env-description:"app env mode, \"dev\" or \"prod\""`
}

// HTTP hold HTTP configuration.
type HTTP struct {
	Host string `yaml:"host" env-required:"true" env:"HOST" env-description:"app http server host, e.g \"localhost\", \"0.0.0.0\""`
	Port int    `yaml:"port" env-required:"true" env:"PORT" env-description:"app http server port, e.g 8080"`
}

// GRPC hold GRPC configuration.
type GRPC struct {
	Host string `yaml:"host" env-required:"true" env:"HOST" env-description:"app grpc server host, e.g \"localhost\", \"0.0.0.0\""`
	Port int    `yaml:"port" env-required:"true" env:"PORT" env-description:"app grpc server port, e.g 9090"`
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
	LogLevel logLevel `yaml:"log_level" env-required:"true" env:"LOG_LEVEL" env-description:"log level minimum, \"panic\", \"fatal\", \"error\", \"warn\", \"warning\", \"info\", \"debug\", \"trace\""`
}

// PG hold postgres configuration.
type PG struct {
	PoolMax  int    `yaml:"pool_max" env-required:"true" env:"POOL_MAX" env-description:"maximum size of postgres pool connection, e.g 10"`
	Username string `yaml:"username" env-required:"true" env:"USERNAME" env-description:"postgres user"`
	Password string `yaml:"password" env-required:"true" env:"PASSWORD" env-description:"postgres password"`
	Host     string `yaml:"host"     env-required:"true" env:"HOST"     env-description:"postgres host"`
	Port     int    `yaml:"port"     env-required:"true" env:"PORT"     env-description:"postgres port"`
	DBName   string `yaml:"db_name"  env-required:"true" env:"DB_NAME"  env-description:"postgres database name"`
}

// JWT hold JWT configuration.
type JWT struct {
	ExpireHour int    `yaml:"expire_hour" env-required:"true" env:"EXPIRE_HOUR" env-description:"jwt expire in hour, e.g 24 for 1 day"`
	SignedKey  string `yaml:"signed_key"  env-required:"true" env:"SIGNED_KEY"  env-description:"jwt signed key"`
}
