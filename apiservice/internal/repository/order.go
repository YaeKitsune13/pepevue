package repository

import (
	"apiservice/internal/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetAll() ([]model.Order, error) {
	var orders []model.Order
	result := r.db.Find(&orders)
	return orders, result.Error
}

func (r *OrderRepository) Create(order *model.Order) error {
	result := r.db.Create(order)
	return result.Error
}
