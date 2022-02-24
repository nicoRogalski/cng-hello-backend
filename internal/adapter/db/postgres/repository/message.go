package repository

import (
	"context"

	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/model"
	"github.com/rs/zerolog/log"
)

type MessageRepository struct{}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{}
}

func (mr MessageRepository) GetMessages(ctx context.Context) (m []*model.Message) {
	log.Debug().Msg("Get messages")
	db := postgres.DBConn
	db.WithContext(ctx).Find(&m)
	return
}

func (mr MessageRepository) GetMessage(ctx context.Context, id string) (m *model.Message) {
	log.Debug().Msg("Get message")
	db := postgres.DBConn
	db.WithContext(ctx).Where("id = ?", id).First(&m)
	return
}

func (mr MessageRepository) CreateMessage(ctx context.Context, m *model.Message) {
	log.Debug().Msgf("Create message code: %s", m.Code)
	db := postgres.DBConn
	db.WithContext(ctx).Create(&m)
}

func (mr MessageRepository) DeleteMessage(ctx context.Context, id string) {
	log.Debug().Msgf("Delete message id: %s", id)
	db := postgres.DBConn
	db.WithContext(ctx).Delete(&model.Message{}, id)
}
