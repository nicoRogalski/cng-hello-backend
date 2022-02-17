package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nicoRogalski/cng-hello-backend/internal/utils/config"
	"github.com/rs/zerolog/log"
)

type ginHands struct {
	SerName    string
	Path       string
	Latency    time.Duration
	Method     string
	StatusCode int
	MsgStr     string
}

func SetupGinLogger(r *gin.Engine) {
	r.Use(logger())

}
func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// before request
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()
		// after request
		// latency := time.Since(t)
		// clientIP := c.ClientIP()
		// method := c.Request.Method
		// statusCode := c.Writer.Status()
		if raw != "" {
			path = path + "?" + raw
		}
		msg := c.Errors.String()
		if msg == "" {
			msg = "Request"
		}
		cData := &ginHands{
			SerName:   config.Cfg.ServiceName,
			Path:       path,
			Latency:    time.Since(t),
			Method:     c.Request.Method,
			StatusCode: c.Writer.Status(),
			MsgStr:     msg,
		}

		if path != "/metrics" {
			log.Info().
				Str("server", cData.SerName).
				Str("method", cData.Method).
				Str("path", cData.Path).
				Dur("resp_time", cData.Latency).
				Int("status", cData.StatusCode).
				Msg(cData.MsgStr)
		}
	}
}
