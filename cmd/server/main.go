package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
		otelzap.S().Error("Error serve http server: ", err)
	}
}

func setup() {
	SetupConfig()
	logger.Setup(Config.ServiceName, Config.IsDevMode)
	tracer.Setup(Config.JaegerEndpoint, Config.ServiceName, Config.IsTracingEnabled)
	go auth.Setup(Config.JwkSetUri)
	go postgres.InitConnection(Config.PostgresHost, Config.PostgresUser, Config.PostresPassword, Config.PostgresDb, Config.PostgresPort)
}

func run() error {

	if Config.IsDevMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	setupRoutes(r)

	port := Config.Port
	return http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
