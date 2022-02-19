package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/internal/adapter/rest/dto"
	"github.com/rogalni/cng-hello-backend/internal/pkg/tracer"
	"github.com/rogalni/cng-hello-backend/internal/service"
	"github.com/rogalni/cng-hello-backend/internal/utils"
	"github.com/rs/zerolog/log"
)

func GetHello(c *gin.Context) {
	span := tracer.Start(c.Request.Context(), "handler.GetHello")
	defer span.End()
	log.Info().
		Str("trace", span.SpanContext().TraceID().String()).
		Str("span", span.SpanContext().SpanID().String()).
		Msg("Get Hello in new span")

	hs := service.NewHelloService()
	m := hs.GetMessage(c.Request.Context())
	c.IndentedJSON(200, &dto.Message{
		Id:   m.Id,
		Code: m.Code,
		Text: m.Text,
	})
}

func GetHelloSecure(c *gin.Context) {
	role := utils.GetJWTRole(c)
	log.Info().Msgf("Authorized with role: %s", role)
	span := tracer.Start(c.Request.Context(), "GetHelloSecure")
	defer span.End()
	hs := service.NewHelloService()
	m := hs.GetMessage(c.Request.Context())
	c.IndentedJSON(200, &dto.Message{
		Id:   m.Id,
		Code: m.Code,
		Text: m.Text,
	})
}
