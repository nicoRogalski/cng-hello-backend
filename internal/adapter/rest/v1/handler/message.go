package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/model"
	"github.com/rogalni/cng-hello-backend/internal/adapter/rest/v1/dto"
	"github.com/rogalni/cng-hello-backend/internal/service"
)

func SetupMessages(g *gin.RouterGroup) {
	messages := g.Group("/messages")
	{
		messages.GET("/", getMessages)
		messages.GET("/:id", getMessage)
		messages.POST("/", createMessage)
		messages.DELETE("/:id", deleteMessage)
	}
}

func getMessages(c *gin.Context) {
	hs := service.NewMessageService()
	m := hs.GetMessages(c.Request.Context())
	c.IndentedJSON(200, toDtos(m))
}

func getMessage(c *gin.Context) {
	hs := service.NewMessageService()
	id := c.Param("id")
	m := hs.GetMessage(c.Request.Context(), id)
	c.IndentedJSON(200, &dto.Message{
		Id:   m.Id,
		Code: m.Code,
		Text: m.Text,
	})
}

func createMessage(c *gin.Context) {
	hs := service.NewMessageService()
	var m *model.Message
	c.Bind(&m)
	hs.CreateMessage(c.Request.Context(), m)
	c.Status(204)
}

func deleteMessage(c *gin.Context) {
	hs := service.NewMessageService()

	id := c.Param("id")
	hs.DeleteMessage(c.Request.Context(), id)
	c.Status(204)
}

func toDto(m *model.Message) *dto.Message {
	return &dto.Message{
		Id:   m.Id,
		Code: m.Code,
		Text: m.Text,
	}
}

func toDtos(mm []*model.Message) (dtoM []dto.Message) {
	for _, m := range mm {
		dtoM = append(dtoM, *toDto(m))
	}
	return
}
