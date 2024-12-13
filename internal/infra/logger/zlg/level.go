package zlg

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const defaultLogLevel = zerolog.InfoLevel

func Level(cfg Config) zerolog.Level {
	level, err := zerolog.ParseLevel(cfg.LogLevel)
	if cfg.LogLevel == "" {
		level = defaultLogLevel
	}

	if err != nil {
		level = defaultLogLevel
		log.Warn().Msgf(fmt.Sprintf("invalid log level '%s', using default: %s", cfg.LogLevel, level))
	}

	return level
}
