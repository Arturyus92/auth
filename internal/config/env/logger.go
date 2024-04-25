package env

import (
	"errors"
	"os"

	"github.com/Arturyus92/auth/internal/config"
)

var _ config.LoggerConfig = (*loggerConfig)(nil)

const (
	loggerLevel = "LOGGER_LEVEL"
)

type loggerConfig struct {
	loggerLevel string
}

// NewLoggerConfig - ...
func NewLoggerConfig() (*loggerConfig, error) {
	loggerLevel := os.Getenv(loggerLevel)
	if len(loggerLevel) == 0 {
		return nil, errors.New("loggerLevel not found")
	}

	return &loggerConfig{
		loggerLevel: loggerLevel,
	}, nil
}

func (cfg *loggerConfig) LoggerLevel() string {
	return cfg.loggerLevel
}
