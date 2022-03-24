package log

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

func Setup(serviceName string, isJson bool, isDebug bool, isDev bool) {
	// Setting up the default logger
	if isDev {
		log.Logger = zerolog.New(os.Stderr).
			With().
			Timestamp().
			Caller().
			Logger()
	} else {
		log.Logger = zerolog.New(os.Stderr).
			With().
			Str("service", serviceName).
			Timestamp().
			Caller().
			Logger()
	}
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

// Wrapper arround zerolog for adding traceing from context into log messages
type TracedZeroLog struct {
	sc trace.SpanContext
}

func Ctx(c context.Context) *TracedZeroLog {
	sc := trace.SpanFromContext(c).SpanContext()
	return &TracedZeroLog{sc: sc}
}

func Span(span trace.Span) *TracedZeroLog {
	sc := span.SpanContext()
	return &TracedZeroLog{sc: sc}
}

func (zl *TracedZeroLog) Trace() *zerolog.Event {
	return appendTrace(log.Trace(), zl)
}
func (zl *TracedZeroLog) Debug() *zerolog.Event {
	return appendTrace(log.Debug(), zl)
}

func (zl *TracedZeroLog) Info() *zerolog.Event {
	return appendTrace(log.Info(), zl)
}

func (zl *TracedZeroLog) Warn() *zerolog.Event {
	return appendTrace(log.Warn(), zl)
}

func (zl *TracedZeroLog) Error() *zerolog.Event {
	return appendTrace(log.Error(), zl)
}

func (zl *TracedZeroLog) Fatal() *zerolog.Event {
	return appendTrace(log.Fatal(), zl)
}

func appendTrace(e *zerolog.Event, zl *TracedZeroLog) *zerolog.Event {
	if zl.sc.IsValid() {
		e.Str("trace", zl.sc.TraceID().String()).
			Str("span", zl.sc.SpanID().String())
	}
	return e
}
