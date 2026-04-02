package repository

import (
	"apiservice/internal/model"

	"gorm.io/gorm"
)

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{
		db: db,
	}
}

func (r *OrderItemRepository) CreateItems(tx *gorm.DB, items []model.OrderItem) error {
	return tx.Create(&items).Error
}
