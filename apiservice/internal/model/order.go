package model

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderStatusNew       OrderStatus = "Новый"
	OrderStatusPending   OrderStatus = "Отправка"
	OrderStatusCompleted OrderStatus = "Завершен"
	OrderStatusCancelled OrderStatus = "Отменен"
)

type Order struct {
	gorm.Model
	PriceAll     float64     `json:"price_all" gorm:"column:price_all"`
	UserID       uint        `json:"user_id" gorm:"column:user_id"`
	User         Account     `json:"-" gorm:"foreignKey:UserID"`
	Status       OrderStatus `json:"status" gorm:"column:status;default:Новый"`
	Address      string      `json:"address" gorm:"column:address"`
	DateDelivery time.Time   `json:"date_delivery" gorm:"column:dateDelivery"`
}
