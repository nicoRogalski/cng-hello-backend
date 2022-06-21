package dto

import "github.com/google/uuid"

type Message struct {
	Id      uuid.UUID `json:"id"`
	Code    string    `json:"code"`
	Text    string    `json:"text"`
	Version int       `json:"version"`
}

type CreateMessage struct {
	Code string `json:"code"`
	Text string `json:"text"`
}
