package env

import (
	"errors"
)

var (
	ErrMissingEnvironmentFileEnvVar = errors.New("missing environment file environment variable")
)

// Environment defines the full environment for the application as a typed struct.
type Environment struct {
	Discord Discord `envconfig:"DISCORD"`
}

type Discord struct {
	APIKey        string `envconfig:"API_KEY"`
	ApplicationID string `envconfig:"APPLICATION_ID"`
}

type PostgreSQL struct {
	URL string `envconfig:"POSTGRES_DATABASE_URL"`
}
