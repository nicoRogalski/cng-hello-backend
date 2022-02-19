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
	if !config.App.IsDevMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	logger.SetupGin(r)
	setupRoutes(r)

	http.ListenAndServe(":"+config.App.Port, r)
}

func setupRoutes(r *gin.Engine) {
	handler.SetupMetrics(r)
	handler.SetupHealth(r)

	api := r.Group("/api")

	tracer.ForGroup(api)
	{
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

}
