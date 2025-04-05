package config

import (
	"errors"
)

// DB database config
type DB struct {
	Host string `env:"HOST"`
	Port string `env:"PORT"`
	User string `env:"USER"`
	Pass string `env:"PASS"`
	Name string `env:"NAME"`
}

// Validate validate
func (db *DB) Validate() error {
	if db.Host == "" {
		return errors.New("DB host required")
	}
	if db.Port == "" {
		return errors.New("DB port required")
	}
	if db.User == "" {
		return errors.New("DB user required")
	}
	if db.Pass == "" {
		return errors.New("DB password required")
	}

	return nil
}
