package service

import (
	"errors"

	"apiservice/internal/model"
	"apiservice/internal/repository"

	"gorm.io/gorm"
)

type CartService struct {
	repo     *repository.CartRepository
	products *ProductService
}

func NewCartService(repo *repository.CartRepository, products *ProductService) *CartService {
	return &CartService{repo: repo, products: products}
}

func (s *CartService) GetCart(userID uint) ([]model.Cart, error) {
	return s.repo.GetFullCart(userID)
}

// AddProduct добавляет товар; если он уже в корзине — увеличивает количество.
func (s *CartService) AddProduct(userID, productID uint, count int16) error {
	if count <= 0 {
		return ErrInvalidCount
	}
	product, err := s.products.GetProduct(productID)
	if err != nil {
		return err
	}

	existing, err := s.repo.GetByUserAndProduct(userID, productID)
	switch {
	case err == nil:
		if int(existing.Count)+int(count) > int(product.Count) {
			return ErrNotEnoughStock
		}
		return s.repo.UpdateProductCount(userID, productID, existing.Count+count)
	case errors.Is(err, gorm.ErrRecordNotFound):
		if count > product.Count {
			return ErrNotEnoughStock
		}
		return s.repo.AddToCart(&model.Cart{UserID: userID, ProductID: productID, Count: count})
	default:
		return err
	}
}

// SetCount устанавливает точное количество товара, уже лежащего в корзине.
func (s *CartService) SetCount(userID, productID uint, count int16) error {
	if count <= 0 {
		return ErrInvalidCount
	}
	product, err := s.products.GetProduct(productID)
	if err != nil {
		return err
	}
	if count > product.Count {
		return ErrNotEnoughStock
	}
	if _, err := s.repo.GetByUserAndProduct(userID, productID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrCartItemNotFound
		}
		return err
	}
	return s.repo.UpdateProductCount(userID, productID, count)
}

// RemoveProduct идемпотентно удаляет позицию из корзины.
func (s *CartService) RemoveProduct(userID, productID uint) error {
	return s.repo.DeleteFromCart(userID, productID)
}
