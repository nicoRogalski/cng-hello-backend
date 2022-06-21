package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/model"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/repository"
	"github.com/rogalni/cng-hello-backend/pkg/log"
)

type Message struct {
	messageRepository repository.IMessage
}

func NewMessage() *Message {
	return &Message{
		messageRepository: repository.NewMessage(),
	}
}

func (ms Message) GetMessages(ctx context.Context) ([]*model.Message, error) {
	log.Ctx(ctx).Info().Msg("Get message from service with trace infos")
	return ms.messageRepository.FindAll(ctx)
}

func (ms Message) GetMessage(ctx context.Context, id uuid.UUID) (*model.Message, error) {
	log.Ctx(ctx).Info().Msg("Get message from service with trace infos")
	return ms.messageRepository.FindById(ctx, id)
}

func (ms Message) CreateMessage(ctx context.Context, m *model.Message) error {
	log.Ctx(ctx).Info().Msg("Create message from service with trace infos")
	return ms.messageRepository.Create(ctx, m)
}

func (ms Message) DeleteMessage(ctx context.Context, id uuid.UUID) error {
	log.Ctx(ctx).Info().Msg("Delete message from service with trace infos")
	return ms.messageRepository.Delete(ctx, id)
}
