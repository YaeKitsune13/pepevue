package repository

import (
	"apiservice/internal/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) CreateProduct(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) GetAllProducts() ([]model.Product, error) {
	var products []model.Product

	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductById(id uint) (model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	return product, err
}
