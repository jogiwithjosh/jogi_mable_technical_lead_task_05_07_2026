package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"

	"api/internal/config"
)

func New(cfg *config.Config) zerolog.Logger {
	level := zerolog.InfoLevel

	switch strings.ToLower(cfg.LogLevel) {

	case "debug":
		level = zerolog.DebugLevel

	case "warn":
		level = zerolog.WarnLevel

	case "error":
		level = zerolog.ErrorLevel

	}

	zerolog.SetGlobalLevel(level)

	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("application", cfg.AppName).
		Logger()

	return logger
}
