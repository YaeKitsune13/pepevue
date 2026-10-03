package repository

import (
	"apiservice/internal/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetProductById(id uint) (model.Product, error) {
	var product model.Product
	err := r.db.Preload("Category").First(&product, id).Error
	return product, err
}

// DecreaseStock атомарно уменьшает остаток. false — если остатка не хватило.
func (r *ProductRepository) DecreaseStock(tx *gorm.DB, productID uint, count int16) (bool, error) {
	res := tx.Model(&model.Product{}).
		Where("id = ? AND `count` >= ?", productID, count).
		Update("count", gorm.Expr("`count` - ?", count))
	return res.RowsAffected > 0, res.Error
}
