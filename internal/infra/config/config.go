package config

import (
	"errors"
	"os"
	"strings"
)

type Settings struct {
	DatabaseURL string
	JwtKey      string
}

const (
	DATABASE_PATH = "DATABASE_PATH"
	JWT_KEY       = "JWT_KEY"
)

func Load() (*Settings, error) {
	// Load configurations, e.g., from environment variables or a config file
	var err error
	var dbURL string

	if dbURL, err = loadEnvVar(DATABASE_PATH); err != nil {
		return nil, err
	}

	var jwtKey string
	if jwtKey, err = loadEnvVar(JWT_KEY); err != nil {
		return nil, err
	}

	settings := &Settings{
		DatabaseURL: dbURL,
		JwtKey:      jwtKey,
	}

	return settings, nil
}

func loadEnvVar(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New(key + " environment variable is not set")
	}
	return strings.TrimSpace(value), nil

}
