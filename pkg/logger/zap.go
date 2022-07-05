package logger

import (
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Setup(serviceName string, isDev bool) {
	var log *zap.Logger
	if isDev {
		log, _ = zap.NewDevelopment()
	} else {
		log, _ = zap.NewProduction(zap.Fields(defaultFields(serviceName)...))
	}
	otelzap.ReplaceGlobals(otelzap.New(log, otelzap.WithMinLevel(zap.DebugLevel), otelzap.WithTraceIDField(true)))
}

func defaultFields(serviceName string) []zapcore.Field {
	return []zapcore.Field{
		zap.String("service", serviceName),
	}
}
