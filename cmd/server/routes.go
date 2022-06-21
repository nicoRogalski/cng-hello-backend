package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/config"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres"
	"github.com/rogalni/cng-hello-backend/internal/adapter/rest/v1/handler"
	"github.com/rogalni/cng-hello-backend/pkg/gin/health"
	"github.com/rogalni/cng-hello-backend/pkg/gin/log"
	"github.com/rogalni/cng-hello-backend/pkg/gin/metrics"
	"github.com/rogalni/cng-hello-backend/pkg/gin/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func setupRoutes(r *gin.Engine) {
	health.For(r).
		WithLiveness(serverStatus).
		WithReadiness(serverStatus)

	r.GET("/metrics", metrics.GetHandler)

	api := r.Group("/api")
	// Tracing and endpoint logging is attached to "/api" route
	api.Use(otelgin.Middleware(config.App.ServiceName))
	api.Use(log.Logger(config.App.ServiceName))
	api.Use(gzip.Gzip(gzip.DefaultCompression))

	// "v1" simulates a real world example of endpoints
	v1 := api.Group("/v1")
	{
		m := v1.Group("/messages")
		{
			mh := handler.NewMessage()
			m.GET("", mh.GetMessages)
			m.GET("/:id", mh.GetMessageById)
			m.POST("", mh.CreateMessage)
			m.DELETE("/:id", middleware.ValidateJWT, mh.DeleteMessage)
		}
	}
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
