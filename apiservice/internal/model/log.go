package model

import "time"

type Log struct {
	ID        uint      `json:"id" gorm:"primaryKey;column:id"`
	Code      uint16    `json:"code" gorm:"column:code"`
	Message   string    `json:"message" gorm:"column:message"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}
