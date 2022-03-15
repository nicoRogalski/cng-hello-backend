package dto

import "time"

type Error struct {
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}
