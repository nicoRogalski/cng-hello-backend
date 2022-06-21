package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/model"
	perrors "github.com/rogalni/cng-hello-backend/pkg/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type IMessage interface {
	FindAll(ctx context.Context) ([]*model.Message, error)
	FindById(ctx context.Context, id uuid.UUID) (*model.Message, error)
	Create(ctx context.Context, m *model.Message) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Message struct{}

func NewMessage() *Message {
	return &Message{}
}

func (mr Message) FindAll(ctx context.Context) (m []*model.Message, err error) {
	log.Debug().Msg("Get messages")
	db := postgres.DBConn
	err = db.WithContext(ctx).Find(&m).Error
	return
}

func (mr Message) FindById(ctx context.Context, id uuid.UUID) (m *model.Message, err error) {
	log.Debug().Msg("Get message")
	db := postgres.DBConn
	err = db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = perrors.NewErrorNotFound("Message not found")
		} else {
			log.Err(err).Msg("Error getting message")
			err = perrors.NewErrInternalServer("Internal Server Error")
		}
	}
	return
}

func (mr Message) Create(ctx context.Context, m *model.Message) error {
	log.Debug().Msgf("Create message code: %s", m.Code)
	db := postgres.DBConn
	err := db.WithContext(ctx).Create(&m).Error
	if err != nil {
		log.Err(err).Msg("Error creating message")
	}
	return err
}

func (mr Message) Delete(ctx context.Context, id uuid.UUID) error {
	log.Debug().Msgf("Delete message id: %s", id)
	db := postgres.DBConn
	err := db.WithContext(ctx).Delete(&model.Message{Id: id}).Error
	if err != nil {
		log.Err(err).Msg("Error deleting message")
	}
	return err
}
