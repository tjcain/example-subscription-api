package config

import "github.com/kelseyhightower/envconfig"

// Configuration represents a subapi config.
type Configuration struct {
	Database DatabaseConfig
	Port     string `envconfig:"PORT" default:"4000"`
}

// DatabaseConfig encapsulates a database configuration.
type DatabaseConfig struct {
	Name     string `envconfig:"DATABASE_NAME"`
	Username string `envconfig:"DATABASE_USERNAME"`
	Password string `envconfig:"DATABASE_PASSWORD"`
	Host     string `envconfig:"DATABASE_HOST"`
	Port     string `envconfig:"DATABASE_PORT"`
}

// AuthConfig encapsulates an (O)Auth configuration.
type AuthConfig struct {
	AuthDomain   string `envconfig:"AUTH_DOMAIN"`
	AuthAudience string `envconfig:"AUTH_AUDIENCE"`
}

// New returns a new configuration based on the required environment variables.
func New() (*Configuration, error) {
	var c Configuration
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
