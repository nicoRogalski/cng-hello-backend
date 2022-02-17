package logger

import (
	"os"

	"github.com/nicoRogalski/cng-hello-backend/internal/utils/config"
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

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.Cfg.DevMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}
