package dto

import (
	"apiservice/internal/model"
	"time"
)

type OrderResponse struct {
	PriceAll     float64           `json:"price_all"`
	UserID       uint              `json:"user_id"`
	Status       model.OrderStatus `json:"order_status"`
	Address      string            `json:"address"`
	DateDelivery time.Time         `json:"date_delivery"`
}

type OrderUpdateStatus struct {
	Status model.OrderStatus `json:"status" binding:"required"`
}
