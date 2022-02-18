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
	r.GET("/health", handler.GetHealth)
	r.GET("/metrics", handler.GetMetrics)

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.GET("/hello", handler.GetHello)
	}
	v2 := api.Group("/v2")
	v2.Use(middleware.AuthorizeJWT())
	{
		v2.GET("/hello", handler.GetHello)
	}
}
