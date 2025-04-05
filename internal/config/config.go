package config

import (
	"github.com/caarlos0/env/v11"
)

var (
	conf Config
)

// Config system config
type Config struct {
	App App `envPrefix:"APP_"`
	Log Log `envPrefix:"LOG_"`
	DB  DB  `envPrefix:"DB_"`
}

// Get return conf instance
func Get() Config {
	return conf
}

// Load load env
func Load() (Config, error) {
	if err := env.Parse(&conf); err != nil {
		return conf, nil
	}
	if err := conf.Validate(); err != nil {
		return conf, nil
	}

	return conf, nil
}

// Validate validate
func (conf Config) Validate() error {
	if err := conf.DB.Validate(); err != nil {
		return err
	}
	if err := conf.App.Validate(); err != nil {
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
