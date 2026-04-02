package service

import (
	"apiservice/internal/model"
	"apiservice/internal/repository"
	"errors"
)

type CartService struct {
	repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (s *CartService) AddProductToCart(userID uint, productID uint, count int16, prodService *ProductService) error {
	product, err := prodService.GetProduct(productID)
	if err != nil {
		return errors.New("продукт не найден")
	}

	if count <= 0 {
		return errors.New("количество должно быть больше 0")
	}

	if count > product.Count {
		return errors.New("недостаточно товара на складе")
	}

	// 4. Логика "Умного добавления" (Опционально, но круто для оценки)
	existingItem, err := s.repo.GetByUserAndProduct(userID, productID)

	if err == nil {
		newCount := existingItem.Count + count

		if newCount > product.Count {
			return errors.New("нельзя добавить больше, чем есть на складе")
		}

		return s.repo.UpdateProductCount(userID, existingItem.ID, newCount)
	}
	newItem := &model.Cart{
		UserID:    userID,
		ProductID: productID,
		Count:     count,
	}

	return s.repo.AddToCart(newItem)
}

func (serv *CartService) UpdateCountCart(productId uint, userId uint, count int16, prod_service *ProductService) error {
	product, err := prod_service.GetProduct(productId)

	if err != nil {
		return errors.New("Не нашли продукт в базе")
	}

	if count < 0 {
		return errors.New("Количество не может быть меньше 0")
	}

	if count >= int16(product.Count) {
		return errors.New("Кол-во товаров на складе закончилось")
	}

	return nil
}
