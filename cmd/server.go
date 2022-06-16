package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/rogalni/cng-hello-backend/config"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres"
	"github.com/rogalni/cng-hello-backend/internal/adapter/rest/v1/handler"
	"github.com/rogalni/cng-hello-backend/pkg/auth"
	"github.com/rogalni/cng-hello-backend/pkg/gin/health"
	"github.com/rogalni/cng-hello-backend/pkg/gin/log"
	"github.com/rogalni/cng-hello-backend/pkg/gin/metrics"
	"github.com/rogalni/cng-hello-backend/pkg/gin/middleware"
	zlog "github.com/rogalni/cng-hello-backend/pkg/log"
	"github.com/rogalni/cng-hello-backend/pkg/tracer"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	config.Setup()
	zlog.Setup(config.App.ServiceName, config.App.IsJsonLogging, config.App.IsLogLevelDebug, config.App.IsDevMode)
	tracer.Setup(config.App.JaegerEndpoint, config.App.ServiceName, config.App.IsTracingEnabled)

	go auth.SetupAuth(config.App.JwkSetUri)
	go postgres.InitConnection()

	if !config.App.IsDevMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(middleware.ErrorHandler)
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	setupRoutes(r)

	port := config.App.Port
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

func setupRoutes(r *gin.Engine) {
	health.For(r).
		WithLiveness(serverStatus).
		WithReadiness(serverStatus)

	metrics.For(r)

	api := r.Group("/api")
	// Tracing and endpoint logging is attached to "/api" route
	api.Use(otelgin.Middleware(config.App.ServiceName))
	log.ForGroup(api, config.App.ServiceName)

	// "v1" simulates a real world example of endpoints
	v1 := api.Group("/v1")
	handler.RegisterMessages(v1)
}

func serverStatus() (h health.Health) {

	phc := health.Component{Name: "postgres"}
	if db, err := postgres.DBConn.DB(); err != nil {
		phc.Status = health.DOWN
	} else if err := db.Ping(); err != nil {
		phc.Status = health.DOWN
	} else {
		phc.Status = health.UP
	}

	h.Components = []health.Component{phc}
	return
}
