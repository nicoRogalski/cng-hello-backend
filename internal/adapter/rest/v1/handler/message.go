package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/model"
	"github.com/rogalni/cng-hello-backend/internal/adapter/rest/v1/dto"
	"github.com/rogalni/cng-hello-backend/internal/service"
)

type messageHandler struct {
	ms *service.MessageService
}

func RegisterMessages(g *gin.RouterGroup) {
	h := messageHandler{ms: service.NewMessageService()}
	messages := g.Group("/messages")
	messages.GET("", h.getMessages)
	messages.GET("/:id", h.getMessage)
	messages.POST("", h.createMessage)
	messages.DELETE("/:id", h.deleteMessage)

}

func (h messageHandler) getMessages(c *gin.Context) {
	m, err := h.ms.GetMessages(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}
	c.IndentedJSON(200, toDtos(m))
}

func (h messageHandler) getMessage(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))
	m, err := h.ms.GetMessage(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	c.IndentedJSON(200, &dto.Message{
		Id:   m.Id,
		Code: m.Code,
		Text: m.Text,
	})
}

func (h messageHandler) createMessage(c *gin.Context) {
	var m *dto.Message
	c.Bind(&m)
	if err := h.ms.CreateMessage(c.Request.Context(), toEntity(m)); err != nil {
		c.Error(err)
		return
	}
	c.Status(204)
}

func (h messageHandler) deleteMessage(c *gin.Context) {
	id := uuid.MustParse(c.Param("id"))
	if err := h.ms.DeleteMessage(c.Request.Context(), id); err != nil {
		c.Error(err)
		return
	}
	c.Status(204)
}

func toEntity(m *dto.Message) *model.Message {
	return &model.Message{
		Id:   m.Id,
		Code: m.Code,
		Text: m.Text,
	}
}

func toDto(m *model.Message) *dto.Message {
	return &dto.Message{
		Id:   m.Id,
		Code: m.Code,
		Text: m.Text,
	}
}

func toDtos(mm []*model.Message) (dtoM []dto.Message) {
	dtoM = make([]dto.Message, 0)
	for _, m := range mm {
		dtoM = append(dtoM, *toDto(m))
	}
	return
}
