package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

type Message struct {
	Id        uuid.UUID `gorm:"primary_key;type:uuid"`
	Code      string    `gorm:"column:code"`
	Text      string    `gorm:"column:text"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Version   optimisticlock.Version
}

func (m *Message) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.New()
	return
}
