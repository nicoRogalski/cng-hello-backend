package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nicoRogalski/cng-hello-backend/internal/adapter/rest"
	"github.com/nicoRogalski/cng-hello-backend/internal/utils/config"
	"github.com/nicoRogalski/cng-hello-backend/internal/utils/logger"
	"github.com/nicoRogalski/cng-hello-backend/internal/utils/trace"
)

func main() {
	if !config.Cfg.IsDevMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	logger.SetupGinLogger(r)
	trace.SetupGinTracer(r)
	setupRoutes(r)
	r.Run()
}

func setupRoutes(r *gin.Engine) {
	r.GET("/health", rest.HealthHandler)
	r.GET("/metrics", rest.MetricsHandler)

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.GET("/hello", rest.HelloHandler)
	}

}
