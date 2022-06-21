package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/config"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres"
	"github.com/rogalni/cng-hello-backend/pkg/auth"
	"github.com/rogalni/cng-hello-backend/pkg/gin/middleware"
	"github.com/rogalni/cng-hello-backend/pkg/log"
	"github.com/rogalni/cng-hello-backend/pkg/tracer"
)

func main() {
	config.Setup()
	log.Setup(config.App.ServiceName, config.App.IsJsonLogging, config.App.IsLogLevelDebug, config.App.IsDevMode)
	tracer.Setup(config.App.JaegerEndpoint, config.App.ServiceName, config.App.IsTracingEnabled)

	go auth.Setup(config.App.JwkSetUri)
	go postgres.InitConnection()

	err := run()
	if err != nil {
		log.Ctx(context.Background()).Error().Err(err).Msg("Error serve http server")
	}

}

func run() error {
	if !config.App.IsDevMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.ErrorHandler)
	setupRoutes(r)

	port := config.App.Port
	return http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
