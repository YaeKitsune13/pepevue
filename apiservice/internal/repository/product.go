package repository

import (
	"apiservice/internal/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductReposityory(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	result := r.db.Find(&products)
	return products, result.Error
}

func (r *ProductRepository) Create(product *model.Product) error {
	result := r.db.Create(product)
	return result.Error
}
