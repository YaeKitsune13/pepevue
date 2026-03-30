package dto

import "time"

type LogResponse struct {
	Code      uint16    `json:"code"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
