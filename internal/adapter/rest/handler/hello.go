package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/internal/adapter/rest/dto"
	"github.com/rogalni/cng-hello-backend/internal/service"
	"github.com/rogalni/cng-hello-backend/internal/utils"
	"github.com/rogalni/cng-hello-backend/internal/utils/tracer"
	"github.com/rs/zerolog/log"
)

type HelloController interface {
	GetHello(c *gin.Context)
}

func GetHello(c *gin.Context) {
	role := utils.GetJWTRole(c)
	log.Info().Msgf("Authorized with role: %s", role)
	span := tracer.Start(c.Request.Context(), "GetHello")
	defer span.End()
	hs := service.NewHelloService()
	m := hs.GetMessage()
	c.IndentedJSON(200, &dto.Message{
		Id:   m.Id,
		Code: m.Code,
		Text: m.Text,
	})
}
