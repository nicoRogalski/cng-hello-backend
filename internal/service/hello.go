package service

import (
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/model"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/repository"
	"github.com/rs/zerolog/log"
)

type HelloService struct {
	helloRepository *repository.HelloRepository
}

func NewHelloService() *HelloService {
	return &HelloService{
		helloRepository: repository.NewHelloRepository(),
	}
}

func (h HelloService) GetMessage() *model.Message {
	log.Info().Msg("Get message from service")
	return h.helloRepository.GetMessage()
}
