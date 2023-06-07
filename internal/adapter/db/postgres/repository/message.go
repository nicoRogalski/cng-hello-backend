package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/model"
	perrors "github.com/rogalni/cng-hello-backend/pkg/errors"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"gorm.io/gorm"
)

type IMessageRepository interface {
	FindAll(ctx context.Context) ([]*model.Message, error)
	FindById(ctx context.Context, id uuid.UUID) (*model.Message, error)
	Create(ctx context.Context, m *model.Message) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(gdb *gorm.DB) *MessageRepository {
	return &MessageRepository{
		db: gdb,
	}
}

func (mr *MessageRepository) FindAll(ctx context.Context) (m []*model.Message, err error) {
	otelzap.Ctx(ctx).Debug("Get messages")
	err = mr.db.WithContext(ctx).Find(&m).Error
	return
}

func (mr *MessageRepository) FindById(ctx context.Context, id uuid.UUID) (m *model.Message, err error) {
	otelzap.Ctx(ctx).Debug("Get message")
	err = mr.db.WithContext(ctx).Where("id = ?", id).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = perrors.NewErrorNotFound("Message not found")
		} else {
			otelzap.Ctx(ctx).Error("Error getting message")
			err = perrors.NewErrInternalServer("Internal Server Error")
		}
	}
	return
}

func (mr *MessageRepository) Create(ctx context.Context, m *model.Message) error {
	err := mr.db.WithContext(ctx).Create(&m).Error
	if err != nil {
		otelzap.Ctx(ctx).Error("Error creating message")
	}
	return err
}

func (mr *MessageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := mr.db.WithContext(ctx).Delete(&model.Message{Id: id}).Error
	if err != nil {
		otelzap.Ctx(ctx).Error("Error deleting message")
	}
	return err
}
