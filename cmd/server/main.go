package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/config"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres"
	"github.com/rogalni/cng-hello-backend/pkg/auth"
	"github.com/rogalni/cng-hello-backend/pkg/logger"
	"github.com/rogalni/cng-hello-backend/pkg/tracer"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

func main() {
	setup()
	err := run()
	if err != nil {
		otelzap.L().Error("Error serve http server")
	}
}

func setup() {
	config.Setup()
	logger.Setup(config.App.ServiceName, config.App.IsDevMode)
	tracer.Setup(config.App.JaegerEndpoint, config.App.ServiceName, config.App.IsTracingEnabled)
	go auth.Setup(config.App.JwkSetUri)
	go postgres.InitConnection()
}

func run() error {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	setupRoutes(r)

	port := config.App.Port
	return http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
