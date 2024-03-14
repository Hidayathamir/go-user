package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

// Loader load configuration.
type Loader interface {
	loadConfig(cfg *Config) error
}

var _ Loader = &YamlLoader{}

// YamlLoader load config from config yaml path.
type YamlLoader struct {
	Path string
}

func (y *YamlLoader) loadConfig(cfg *Config) error {
	err := cleanenv.ReadConfig(y.Path, cfg)
	if err != nil {
		return fmt.Errorf("cleanenv.ReadConfig: %w", err)
	}
	return nil
}

var _ Loader = &EnvLoader{}

// EnvLoader load config from env var.
type EnvLoader struct {
	// YAMLPath will read yaml config first then read env var. If you do not
	// specify then env var should have all required config.
	YAMLPath string
}

func (e *EnvLoader) loadConfig(cfg *Config) error {
	if e.YAMLPath != "" {
		err := cleanenv.ReadConfig(e.YAMLPath, cfg)
		if err != nil {
			return fmt.Errorf("cleanenv.ReadConfig: %w", err)
		}
	}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return fmt.Errorf("cleanenv.ReadEnv: %w", err)
	}
	return nil
}
