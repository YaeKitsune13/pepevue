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
	return tx.Create(order).Error
}
func (r *OrderRepository) GetByUserID(userID uint) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.
		Preload("Items.Product"). // Подгружаем позиции заказа и сами товары (если в OrderItem есть связь Product)
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders).Error
	return orders, err
}
