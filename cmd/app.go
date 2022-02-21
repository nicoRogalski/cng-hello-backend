package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/config"
	"github.com/rogalni/cng-hello-backend/internal/adapter/rest/handler"
	"github.com/rogalni/cng-hello-backend/pkg/auth"
	"github.com/rogalni/cng-hello-backend/pkg/gin/auth/middleware"
	"github.com/rogalni/cng-hello-backend/pkg/gin/health"
	"github.com/rogalni/cng-hello-backend/pkg/gin/log"
	"github.com/rogalni/cng-hello-backend/pkg/gin/metrics"
	"github.com/rogalni/cng-hello-backend/pkg/gin/tracer"
	zlog "github.com/rogalni/cng-hello-backend/pkg/log"
)

func main() {
	config.Setup()
	isDev := config.App.IsDevMode
	zlog.Setup(config.App.ServiceName, config.App.IsJsonLogging, config.App.IsLogLevelDebug)
	tracer.Setup(config.App.JaegerEndpoint, config.App.ServiceName, isDev)

	auth.Setup(config.App.JwtCertUrl)

	if !isDev {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	setupRoutes(r)
	http.ListenAndServe(":"+config.App.Port, r)
}

func setupRoutes(r *gin.Engine) {
	health.For(r)
	metrics.For(r)

	api := r.Group("/api")
	tracer.ForGroup(api)
	log.ForGroup(api, config.App.ServiceName)
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/hello", handler.GetHello)
			// "/secure/hello" is a specific path in a group thats needs to be secured via jwt
			v1.GET("/secure/hello", middleware.ValidateJWT, handler.GetHelloSecure)
		}
		// "/secure" simulates a path where all endpoints needs to be secured via jwt
		v2 := api.Group("/secure")
		v2.Use(middleware.ValidateJWT)
		{
			v2.GET("/hello", handler.GetHelloSecure)
		}
	}

}
