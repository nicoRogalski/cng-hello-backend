package tracer

import (
	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/internal/utils/config"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func SetupGin(r *gin.Engine) {
	r.Use(otelgin.Middleware(config.Cfg.ServiceName))
}
