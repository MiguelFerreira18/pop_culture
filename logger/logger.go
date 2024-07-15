package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func Logger(isDebugLevel bool) *zerolog.Logger {
	logLevel := zerolog.InfoLevel
	if isDebugLevel {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &logger
}
