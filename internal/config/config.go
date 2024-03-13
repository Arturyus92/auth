package config

import (
	"github.com/joho/godotenv"
)

// GRPCConfig - ...
type GRPCConfig interface {
	Address() string
}

// HTTPConfig - ...
type HTTPConfig interface {
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
