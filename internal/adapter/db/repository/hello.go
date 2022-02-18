package repository

import (
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/model"
	"github.com/rs/zerolog/log"
)

type HelloRepository struct{}

func NewHelloRepository() *HelloRepository {
	return &HelloRepository{}
}

func (h HelloRepository) GetMessage() *model.Message {
	log.Info().Msg("Get message from repository")
	return &model.Message{
		Id:   "1234",
		Code: "hello",
		Text: "Welcome to Cloud Native Go",
	}
}
