package repository

import (
	"apiservice/internal/model"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) GetByUserAndProduct(userID uint, productID uint) (*model.Cart, error) {
	var cart model.Cart
	if err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepository) AddToCart(item *model.Cart) error {
	return r.db.Create(item).Error
}

func (r *CartRepository) UpdateProductCount(userID uint, productID uint, newCount int16) error {
	return r.db.Model(&model.Cart{}).
		Where("user_id = ? AND product_id = ?", userID, productID).
		Update("count", newCount).Error
}

func (r *CartRepository) DeleteFromCart(userID uint, productID uint) error {
	return r.db.Where("user_id = ? AND product_id = ?", userID, productID).
		Delete(&model.Cart{}).Error
}

func (r *CartRepository) GetFullCart(userID uint) ([]model.Cart, error) {
	return r.GetFullCartTx(r.db, userID)
}

// GetFullCartTx читает корзину внутри переданной транзакции.
func (r *CartRepository) GetFullCartTx(tx *gorm.DB, userID uint) ([]model.Cart, error) {
	var items []model.Cart
	err := tx.Preload("Product").Where("user_id = ?", userID).Order("id").Find(&items).Error
	return items, err
}

func (r *CartRepository) ClearCart(tx *gorm.DB, userID uint) error {
	return tx.Where("user_id = ?", userID).Delete(&model.Cart{}).Error
}
