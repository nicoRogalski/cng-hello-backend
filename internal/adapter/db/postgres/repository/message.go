package repository

import (
	"context"
	"errors"

	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/model"
	"github.com/rogalni/cng-hello-backend/pkg/errs"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type MessageRepository struct{}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{}
}

func (mr MessageRepository) GetMessages(ctx context.Context) (m []*model.Message, err error) {
	log.Debug().Msg("Get messages")
	db := postgres.DBConn
	err = db.WithContext(ctx).Find(&m).Error
	return
}

func (mr MessageRepository) GetMessage(ctx context.Context, id string) (m *model.Message, err error) {
	log.Debug().Msg("Get message")
	db := postgres.DBConn
	err = db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errs.ErrNotFound
	} else {
		log.Err(err).Msg("Error getting message")
		err = errs.ErrInternalServer
	}
	return
}

func (mr MessageRepository) CreateMessage(ctx context.Context, m *model.Message) error {
	log.Debug().Msgf("Create message code: %s", m.Code)
	db := postgres.DBConn
	return db.WithContext(ctx).Create(&m).Error
}

func (mr MessageRepository) DeleteMessage(ctx context.Context, id string) error {
	log.Debug().Msgf("Delete message id: %s", id)
	db := postgres.DBConn
	return db.WithContext(ctx).Delete(&model.Message{}, id).Error
}
