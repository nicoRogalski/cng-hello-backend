package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nicoRogalski/cng-hello-backend/internal/adapter/rest/dto"
	"github.com/nicoRogalski/cng-hello-backend/internal/service"
	"github.com/nicoRogalski/cng-hello-backend/internal/utils/trace"
)

func HelloHandler(c *gin.Context) {
	span := trace.Start(c.Request.Context(), "GetHello")
	defer span.End()
	service := service.NewHelloService()
	m := service.GetHelloMessage()
	c.IndentedJSON(200, &dto.Message{
		Id:   m.Id,
		Code: m.Code,
		Text: m.Text,
	})
}
