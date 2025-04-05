package config

import (
	"github.com/caarlos0/env/v11"
)

var (
	cfg Config
)

// Config system config
type Config struct {
	App App `envPrefix:"APP_"`
	Log Log `envPrefix:"LOG_"`
	DB  DB  `envPrefix:"DB_"`
}

// Get return cfg instance
func Get() Config {
	return cfg
}

// Load load env
func Load() (Config, error) {
	if err := env.Parse(&cfg); err != nil {
		return cfg, nil
	}
	if err := cfg.Validate(); err != nil {
		return cfg, nil
	}

	return cfg, nil
}

// Validate validate
func (cfg Config) Validate() error {
	if err := cfg.DB.Validate(); err != nil {
		return err
	}
	if err := cfg.App.Validate(); err != nil {
		return err
	}

	return nil
}

// Env env
type Env string

const (
	EnvLocal Env = "local"
	EnvDev   Env = "dev"
	EnvStg   Env = "stg"
	EnvProd  Env = "prod"
)

// IsLocal local environment
func (e Env) IsLocal() bool {
	return e == EnvLocal
}

// IsProd production environment
func (e Env) IsProd() bool {
	return e == EnvProd
}
