package repository

import (
	"apiservice/internal/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Create(tx *gorm.DB, order *model.Order) error {
	// Мы передаем tx (*gorm.DB), чтобы метод мог работать внутри транзакции
	return tx.Create(order).Error
}
