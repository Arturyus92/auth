package config

import (
	"github.com/joho/godotenv"
)

// LoggerConfig - ...
type LoggerConfig interface {
	LoggerLevel() string
}

// GRPCConfig - ...
type GRPCConfig interface {
	Address() string
}

// HTTPConfig - ...
type HTTPConfig interface {
	Address() string
}

// SwaggerConfig - ...
type SwaggerConfig interface {
	Address() string
}

// PGConfig - ...
type PGConfig interface {
	DSN() string
}

// Load - ...
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
