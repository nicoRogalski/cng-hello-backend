package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	Id   uuid.UUID `gorm:"primary_key;type:uuid"`
	Code string    `gorm:"column:code"`
	Text string    `gorm:"column:text"`
}

func (m *Message) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.New()
	return
}
