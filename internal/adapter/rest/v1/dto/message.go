package dto

import "github.com/google/uuid"

type Message struct {
	Id   uuid.UUID `json:"id"`
	Code string    `json:"code"`
	Text string    `json:"text"`
}
