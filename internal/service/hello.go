package service

import (
	"context"

	"github.com/rogalni/cng-hello-backend/internal/adapter/db/model"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/repository"
	"github.com/rogalni/cng-hello-backend/internal/pkg/logger"
)

type HelloService struct {
	helloRepository *repository.HelloRepository
}

func NewHelloService() *HelloService {
	return &HelloService{
		helloRepository: repository.NewHelloRepository(),
	}
}

func (h HelloService) GetMessage(ctx context.Context) *model.Message {
	logger.InfoWithTrace(ctx).Msg("Get message from service")
	return h.helloRepository.GetMessage()
}
