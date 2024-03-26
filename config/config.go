package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		PG   `yaml:"postgres"`
	}
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}
	PG struct {
		PoolMax   int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		DB_HOST   string `env-required:"true" yaml:"DB_HOST" env:"DB_HOST"`
		DB_PORT   string `env-required:"true" yaml:"DB_PORT" env:"DB_PORT"`
		DB_USER   string `env-required:"true" yaml:"DB_USER" env:"DB_USER"`
		DB_NAME   string `env-required:"true" yaml:"DB_NAME" env:"DB_NAME"`
		DB_PASS   string `env-required:"true" yaml:"DB_PASS" env:"DB_PASS"`
		TIME_ZONE string `env-required:"true" yaml:"TIME_ZONE" env:"TIME_ZONE"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
