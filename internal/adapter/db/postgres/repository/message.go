package repository

import (
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/model"
	"github.com/rs/zerolog/log"
)

type MessageRepository struct{}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{}
}

func (mr MessageRepository) GetMessages() (m []*model.Message) {
	log.Debug().Msg("Get messages")
	db := postgres.DBConn
	db.Find(&m)
	return
}

func (mr MessageRepository) GetMessage(id string) (m *model.Message) {
	log.Debug().Msg("Get message")
	db := postgres.DBConn
	db.Where("id = ?", id).First(&m)
	return
}

func (mr MessageRepository) CreateMessage(m *model.Message) {
	log.Debug().Msgf("Create message code: %s", m.Code)
	db := postgres.DBConn
	db.Create(&m)
}

func (mr MessageRepository) DeleteMessage(id string) {
	log.Debug().Msgf("Delete message id: %s", id)
	db := postgres.DBConn
	db.Delete(&model.Message{}, id)
}
