package config

import "github.com/cockroachdb/errors"

// App application config
type App struct {
	Env Env `env:"ENV"`
}

// Validate validate
func (a App) Validate() error {
	if a.Env == "" {
		return errors.New("APP_ENV is required")
	}

	return nil
}
