package tracer

import (
	"context"
	"os"

	"github.com/rogalni/cng-hello-backend/internal/utils/config"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

func init() {
	if config.Cfg.IsDevMode {
		return
	}

	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(os.Getenv("JAEGER_ENDPOINT"))))

	if err != nil {
		log.Fatal().Err(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewSchemaless(attribute.KeyValue{
			Key:   semconv.ServiceNameKey,
			Value: attribute.StringValue(config.Cfg.ServiceName),
		})),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
}

func Start(c context.Context, spanName string) trace.Span {
	_, span := otel.Tracer("").Start(c, spanName)
	return span
}
