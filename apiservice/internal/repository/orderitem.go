package repository

import (
	"apiservice/internal/model"

	"gorm.io/gorm"
)

type OrderItemRespository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRespository {
	return &OrderItemRespository{db: db}
}

func (r *OrderItemRespository) GetAll() ([]model.OrderItem, error) {
	var OrderItems []model.OrderItem
	result := r.db.Find(&OrderItems)
	return OrderItems, result.Error
}

func (r *OrderItemRespository) AddAllNewOrders(orders *[]model.OrderItem) error {
	result := r.db.Create(orders)
	return result.Error
}
