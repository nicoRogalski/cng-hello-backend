package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/internal/utils"
	"github.com/rogalni/cng-hello-backend/pkg/gin/tracer"
	"github.com/rs/zerolog/log"
)

type HelloMessage struct {
	Text string `json:"text"`
}

func GetHello(c *gin.Context) {
	span := tracer.Start(c.Request.Context(), "handler.GetHello")
	defer span.End()
	log.Info().
		Str("trace", span.SpanContext().TraceID().String()).
		Str("span", span.SpanContext().SpanID().String()).
		Msg("Get Hello in new span")
	c.IndentedJSON(200, &HelloMessage{
		Text: "Welcome to cloud native go",
	})
}

func GetHelloSecure(c *gin.Context) {
	roles := utils.GetJWTRoles(c)
	log.Info().Msgf("Authorized with role: %s", roles)
	span := tracer.Start(c.Request.Context(), "GetHelloSecure")
	defer span.End()
	c.IndentedJSON(200, &HelloMessage{
		Text: "Welcome to cloud native go with jwt auth",
	})
}
