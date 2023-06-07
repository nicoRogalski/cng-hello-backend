package otel

import (
	"context"

	"github.com/rogalni/cng-hello-backend/config"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Otel struct {
	cleanupTracerFunc         func(context.Context) error
	cleanupMeterFunc          func(context.Context) error
	cleanupGrpcConnectionFunc func() error
	syncLoggerFunc            func() error
}

func Setup(ctx context.Context, cfg *config.Config) func(context.Context) error {
	otel := &Otel{}

	sync := setupLogger(cfg.LogLevel, cfg.IsJsonLogging, cfg.ServiceName)
	otel.syncLoggerFunc = sync

	if !cfg.IsTracingEnabled && !cfg.IsMetricsEnabled {
		return otel.cleanup
	}

	resource, err := otelResource(ctx, cfg.ServiceName)
	if err != nil {
		otelzap.S().Error(err)
	}
	conn := otelGrpcCon(ctx, cfg.OtelCollectorEndpoint)
	otel.cleanupGrpcConnectionFunc = conn.Close

	if cfg.IsTracingEnabled {
		tp, err := setupTracing(ctx, conn, resource)
		if err != nil {
			otelzap.S().Error(err)
		}
		otel.cleanupTracerFunc = tp.Shutdown
	}

	if cfg.IsMetricsEnabled {
		mp, err := setupMetrics(ctx, conn, resource)
		if err != nil {
			otelzap.S().Error(err)
		}
		otel.cleanupMeterFunc = mp.Shutdown
	}

	return otel.cleanup
}

func otelGrpcCon(ctx context.Context, endpoint string) *grpc.ClientConn {
	con, err := grpc.DialContext(ctx, endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		otelzap.S().Error(err)
	}
	return con
}

// Returns a new OpenTelemetry resource describing this application.
func otelResource(ctx context.Context, serviceName string) (*resource.Resource, error) {
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
		return nil, err
	}
	return res, nil
}

func (o *Otel) cleanup(ctx context.Context) error {
	if o.cleanupTracerFunc != nil {
		otelzap.Ctx(ctx).Debug("Shutting down Tracer")
		err := o.cleanupTracerFunc(ctx)
		if err != nil {
			return err
		}
	}

	if o.cleanupMeterFunc != nil {
		otelzap.Ctx(ctx).Debug("Shutting down Meter")
		err := o.cleanupMeterFunc(ctx)
		if err != nil {
			return err
		}
	}

	if o.cleanupGrpcConnectionFunc != nil {
		otelzap.Ctx(ctx).Debug("Close otlp grpc connection")
		err := o.cleanupGrpcConnectionFunc()
		if err != nil {
			return err
		}
	}

	if o.syncLoggerFunc != nil {
		otelzap.Ctx(ctx).Debug("Sync otel logger")
		err := o.syncLoggerFunc()
		if err != nil {
			return err
		}
	}
	return nil
}
