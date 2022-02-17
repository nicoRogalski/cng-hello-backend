package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nicoRogalski/cng-hello-backend/internal/adapter/rest/dto"
	"github.com/nicoRogalski/cng-hello-backend/internal/service"
	"github.com/nicoRogalski/cng-hello-backend/internal/utils/tracer"
)

func GetHello(c *gin.Context) {
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
