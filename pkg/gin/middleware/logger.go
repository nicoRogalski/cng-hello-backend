package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type ginHands struct {
	SerName    string
	Path       string
	Latency    time.Duration
	Method     string
	StatusCode int
	MsgStr     string
}

func Logger(s string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		t := time.Now()
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
			Latency:    time.Duration(time.Since(t).Milliseconds()),
			Method:     c.Request.Method,
			StatusCode: c.Writer.Status(),
			MsgStr:     msg,
		}

		otelzap.Ctx(c.Request.Context()).Debug(cData.MsgStr,
			zap.String("method", cData.Method),
			zap.String("path", cData.Path),
			zap.Duration("resp_time", cData.Latency),
			zap.Int("status", cData.StatusCode))

	}
}
