package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/model"
	"github.com/rogalni/cng-hello-backend/internal/adapter/rest/v1/dto"
	"github.com/rogalni/cng-hello-backend/internal/service"
	"github.com/rogalni/cng-hello-backend/pkg/errors"
	"github.com/rogalni/cng-hello-backend/pkg/gin/auth"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type Message struct {
	ms *service.Message
}

func NewMessage() *Message {
	return &Message{ms: service.NewMessage()}
}

func (h Message) GetMessages(c *gin.Context) {
	m, err := h.ms.GetMessages(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}
	c.IndentedJSON(200, toDtos(m))
}

func (h Message) GetMessageById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(errors.NewErrBadRequest("Missing id parameter"))
		return
	}
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

func (h Message) CreateMessage(c *gin.Context) {
	var m *dto.CreateMessage
	c.Bind(&m)
	err := h.ms.CreateMessage(c.Request.Context(), toEntity(m))
	if err != nil {
		c.Error(err)
		return
	}
	c.Status(204)
}

func (h Message) DeleteMessage(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(errors.NewErrBadRequest("Missing id parameter"))
		return
	}
	roles := auth.GetJWTRoles(c)
	otelzap.Ctx(c.Request.Context()).Sugar().Infof("Authorized with role: %s", roles)
	if err := h.ms.DeleteMessage(c.Request.Context(), id); err != nil {
		c.Error(err)
		return
	}
	c.Status(204)
}

func toEntity(m *dto.CreateMessage) *model.Message {
	return &model.Message{
		Code: m.Code,
		Text: m.Text,
	}
}

func toDto(m *model.Message) *dto.Message {
	return &dto.Message{
		Id:      m.Id,
		Code:    m.Code,
		Text:    m.Text,
		Version: int(m.Version.Int64),
	}
}

func toDtos(mm []*model.Message) (dtoM []dto.Message) {
	dtoM = make([]dto.Message, 0)
	for _, m := range mm {
		dtoM = append(dtoM, *toDto(m))
	}
	return
}
