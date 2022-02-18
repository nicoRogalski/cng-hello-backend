package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/internal/adapter/rest/handler"
	"github.com/rogalni/cng-hello-backend/internal/middleware"
	"github.com/rogalni/cng-hello-backend/internal/utils/config"
	"github.com/rogalni/cng-hello-backend/internal/utils/logger"
	"github.com/rogalni/cng-hello-backend/internal/utils/tracer"
)

func main() {
	if !config.Cfg.IsDevMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	logger.SetupGin(r)
	tracer.SetupGin(r)
	setupRoutes(r)

	http.ListenAndServe(":"+config.Cfg.Port, r)
}

func setupRoutes(r *gin.Engine) {
	r.GET("/metrics", handler.GetMetrics)

	health := r.Group("/health")
	{
		health.GET("/", handler.GetHealth)
		health.GET("/readiness", handler.GetReadiness)
		health.GET("/liveness", handler.GetLiveness)
	}

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.GET("/hello", handler.GetHello)
		// "/secure/hello" is a specific path in a group thats needs to be secured via jwt
		v1.GET("/secure/hello", middleware.ValidateJWT(), handler.GetHelloSecure)
	}
	// "/secure" simulates a path where all endpoints needs to be secured via jwt
	v2 := api.Group("/secure")
	v2.Use(middleware.ValidateJWT())
	{
		v2.GET("/hello", handler.GetHelloSecure)
	}
}
