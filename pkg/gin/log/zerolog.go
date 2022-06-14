package log

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

type ginHands struct {
	SerName    string
	Path       string
	Latency    time.Duration
	Method     string
	StatusCode int
	MsgStr     string
}

func ForGroup(r *gin.RouterGroup, serviceName string) {
	r.Use(logger(serviceName))
}

func ForEngine(r *gin.Engine, serviceName string) {
	r.Use(logger(serviceName))
}

func logger(s string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// before request
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()
		// after request
		if raw != "" {
			path = path + "?" + raw
		}
		msg := c.Errors.String()
		if msg == "" {
			msg = "Request"
		}

		cData := ginHands{
			SerName:    s,
			Path:       path,
			Latency:    time.Since(t),
			Method:     c.Request.Method,
			StatusCode: c.Writer.Status(),
			MsgStr:     msg,
		}

		li := log.Debug()
		sc := trace.SpanFromContext(c.Request.Context()).SpanContext()
		if sc.IsValid() {
			li.Str("trace", sc.TraceID().String()).
				Str("span", sc.SpanID().String())
		}
		li.Str("server", cData.SerName).
			Str("method", cData.Method).
			Str("path", cData.Path).
			Dur("resp_time", cData.Latency).
			Int("status", cData.StatusCode).
			Msg(cData.MsgStr)
	}
}
