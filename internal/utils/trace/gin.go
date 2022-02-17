package trace

import (
	"github.com/gin-gonic/gin"
	"github.com/nicoRogalski/cng-hello-backend/internal/utils/config"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func SetupGinTracer(r *gin.Engine) {
	r.Use(otelgin.Middleware(config.Cfg.ServiceName))
}
