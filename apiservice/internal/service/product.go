package service

import (
	"errors"
	"unicode/utf8"

	"apiservice/internal/model"
	"apiservice/internal/repository"

	"gorm.io/gorm"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (serv *ProductService) NewProduct(product *model.Product) error {
	if utf8.RuneCountInString(product.Title) <= 0 {
		return errors.New("Название продукта не может быть пустым")
	}
	if product.Count <= 0 {
		return errors.New("На складе должен быть хотя бы 1 продукт")
	}
	if product.Cost < 2 {
		return errors.New("Цена продукта не может быть меньше 2$")
	}
	return serv.repo.CreateProduct(product)
}

func (serv *ProductService) GetListProducts() ([]model.Product, error) {
	return serv.repo.GetAllProducts()
}

func (serv *ProductService) GetProduct(id uint) (*model.Product, error) {
	product, err := serv.repo.GetProductById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}
	return &product, nil
}
