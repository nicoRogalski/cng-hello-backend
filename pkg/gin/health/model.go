package health

import "encoding/json"

type Health struct {
	Status     Status      `json:"status"`
	Code       int         `json:"-"`
	Components []Component `json:"components,omitempty"`
}
type Status int

const (
	UP Status = iota
	DOWN
)

func (s Status) String() string {
	switch s {
	case UP:
		return "Up"
	case DOWN:
		return "Down"
	}
	return "unknown"
}

func (s Status) MarshalJSON() ([]byte, error) {
	var st string
	switch s {
	case UP:
		st = "Up"
	case DOWN:
		st = "Down"
	}

	return json.Marshal(st)
}

type Component struct {
	Name   string
	Status Status
}
