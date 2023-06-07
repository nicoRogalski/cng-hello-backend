package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/model"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/repository"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"gorm.io/gorm"
)

type MessageService struct {
	messageRepository repository.IMessageRepository
}

func NewMessagService(db *gorm.DB) *MessageService {
	return &MessageService{
		messageRepository: repository.NewMessageRepository(db),
	}
}

func (ms *MessageService) GetMessages(ctx context.Context) ([]*model.Message, error) {
	otelzap.Ctx(ctx).Info("Get message from service with trace infos")
	return ms.messageRepository.FindAll(ctx)
}

func (ms *MessageService) GetMessage(ctx context.Context, id uuid.UUID) (*model.Message, error) {
	otelzap.Ctx(ctx).Info("Get message from service with trace infos")
	return ms.messageRepository.FindById(ctx, id)
}

func (ms *MessageService) CreateMessage(ctx context.Context, m *model.Message) error {
	otelzap.Ctx(ctx).Info("Create message from service with trace infos")
	return ms.messageRepository.Create(ctx, m)
}

func (ms *MessageService) DeleteMessage(ctx context.Context, id uuid.UUID) error {
	otelzap.Ctx(ctx).Info("Delete message from service with trace infos")
	return ms.messageRepository.Delete(ctx, id)
}
