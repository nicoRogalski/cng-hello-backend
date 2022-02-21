package log

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

func Setup(serviceName string, isJson bool, isDebug bool) {
	// Setting up the default logger
	log.Logger = zerolog.New(os.Stderr).
		With().
		Str("service", serviceName).
		Timestamp().
		Caller().
		Logger()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if isDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if !isJson {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}

func InfoWithTrace(c context.Context) *zerolog.Event {
	sc := trace.SpanFromContext(c).SpanContext()
	return log.Info().Str("trace", sc.TraceID().String()).Str("span", sc.SpanID().String())

}
