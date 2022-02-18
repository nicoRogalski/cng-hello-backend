package logger

import (
	"os"

	"github.com/rogalni/cng-hello-backend/internal/utils/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// Setting up the default logger
	log.Logger = zerolog.New(os.Stderr).
		With().
		Str("server", config.Cfg.ServiceName).
		Timestamp().
		Caller().
		Logger()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if config.Cfg.IsLogLevelDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if !config.Cfg.IsJsonLogging {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

}
