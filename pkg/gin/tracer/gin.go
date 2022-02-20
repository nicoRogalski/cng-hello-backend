package tracer

import (
	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/config"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func ForGroup(r *gin.RouterGroup) {
	r.Use(otelgin.Middleware(config.App.ServiceName))
}

func ForEngine(r *gin.Engine) {
	r.Use(otelgin.Middleware(config.App.ServiceName))
}
