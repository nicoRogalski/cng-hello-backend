package otel

import (
	"context"
	"log"

	"github.com/rogalni/cng-hello-backend/config"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Otel struct {
	shutdownTracerFunc      func(context.Context) error
	shutdownMeterFunc       func(context.Context) error
	syncLoggerFunc          func() error
	closeGrpcConnectionFunc func() error
}

func (o *Otel) Shutdown(ctx context.Context) error {
	otelzap.Ctx(ctx).Debug("Shutting down Tracer")
	err := o.shutdownTracerFunc(ctx)
	if err != nil {
		return err
	}
	otelzap.Ctx(ctx).Debug("Shutting down Meter")
	err = o.shutdownMeterFunc(ctx)
	if err != nil {
		return err
	}

	if o.closeGrpcConnectionFunc != nil {
		otelzap.Ctx(ctx).Debug("Close otlp grpc connection")
		err = o.closeGrpcConnectionFunc()
		if err != nil {
			return err
		}
	}

	otelzap.Ctx(ctx).Debug("Sync otel logger")
	err = o.syncLoggerFunc()
	if err != nil {
		return err
	}
	return nil
}

func Setup(ctx context.Context, cfg *config.Config) func(context.Context) error {
	otel := &Otel{}
	resource := OtelResource(ctx, cfg.ServiceName)

	sync := setupLogger(cfg.LogLevel, cfg.IsJsonLogging, cfg.ServiceName)
	otel.syncLoggerFunc = sync

	if cfg.IsTracingEnabled || cfg.IsMetricsEnabled {
		con := otelGrpcCon(ctx)
		otel.closeGrpcConnectionFunc = con.Close

		if cfg.IsTracingEnabled {
			tp, err := setupTracing(ctx, cfg.ServiceName, con, resource)
			if err != nil {
				log.Fatal(err)
			}
			otel.shutdownTracerFunc = tp.Shutdown
		}

		if cfg.IsMetricsEnabled {
			mp, err := setupMetrics(ctx, cfg.ServiceName, con, resource)
			if err != nil {
				log.Fatal(err)
			}
			otel.shutdownMeterFunc = mp.Shutdown
		}
	}

	err := runtime.Start()
	if err != nil {
		log.Fatal(err)
	}
	return otel.Shutdown
}

func otelGrpcCon(ctx context.Context) *grpc.ClientConn {
	con, err := grpc.DialContext(ctx, "otel-collector:4317",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return con
}

// Returns a new OpenTelemetry resource describing this application.
func OtelResource(ctx context.Context, serviceName string) *resource.Resource {
	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
			attribute.String("service", serviceName),
		),
	)
	if err != nil {
		log.Fatalf("%s: %v", "Failed to create resource", err)
	}
	return res
}
