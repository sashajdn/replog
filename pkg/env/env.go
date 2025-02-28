package env

import (
	"errors"
)

var (
	ErrMissingEnvironmentFileEnvVar = errors.New("missing environment file environment variable")
)

// Environment defines the full environment for the application as a typed struct.
type Environment struct {
	Discord    Discord    `envconfig:"DISCORD"`
	PostgreSQL PostgreSQL `envconfig:"POSTGRESQL"`
}

type Discord struct {
	APIToken      string `envconfig:"API_TOKEN"`
	ApplicationID string `envconfig:"APPLICATION_ID"`
}

type PostgreSQL struct {
	URL string `envconfig:"URL"`
}
