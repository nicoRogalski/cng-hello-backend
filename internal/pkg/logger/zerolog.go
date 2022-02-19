package logger

import (
	"os"

	"github.com/rogalni/cng-hello-backend/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Setup() {
	// Setting up the default logger
	log.Logger = zerolog.New(os.Stderr).
		With().
		Str("server", config.App.ServiceName).
		Timestamp().
		Caller().
		Logger()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if config.App.IsLogLevelDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if !config.App.IsJsonLogging {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

}
