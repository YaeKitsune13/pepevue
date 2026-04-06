package service

import (
	"apiservice/internal/model"
	"apiservice/internal/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

// 1. Исправлено название структуры
type OrderService struct {
	orderRepo      *repository.OrderRepository
	orderItemsRepo *repository.OrderItemRepository
	cartRepo       *repository.CartRepository
	db             *gorm.DB
}

func NewOrderService(ro *repository.OrderRepository, ri *repository.OrderItemRepository, rc *repository.CartRepository, db *gorm.DB) *OrderService {
	return &OrderService{
		orderRepo:      ro,
		orderItemsRepo: ri,
		cartRepo:       rc,
		db:             db,
	}
}

var ErrEmptyCart = errors.New("корзина пуста")

func (s *OrderService) PlaceOrder(userID uint, address string) (*model.Order, error) {
	var finalOrder *model.Order

	err := s.db.Transaction(func(tx *gorm.DB) error {
		cartItems, err := s.cartRepo.GetFullCart(userID)
		if err != nil {
			return err
		}
		if len(cartItems) == 0 {
			return ErrEmptyCart
		}

		var totalPrice float64
		orderItems := make([]model.OrderItem, 0, len(cartItems)) // Предварительно выделяем память

		for _, ci := range cartItems {
			totalPrice += ci.Product.Cost * float64(ci.Count)

			orderItems = append(orderItems, model.OrderItem{
				ProductID: ci.ProductID,
				Count:     ci.Count,
				Cost:      ci.Product.Cost, // Сохраняем цену на момент покупки! Это правильно.
			})
		}

		finalOrder = &model.Order{
			UserID:       userID,
			PriceAll:     totalPrice,
			Status:       model.OrderStatusNew,
			Address:      address,
			DateDelivery: time.Now().AddDate(0, 0, 3),
		}

		if err := s.orderRepo.Create(tx, finalOrder); err != nil {
			return err
		}

		// Устанавливаем OrderID для всех элементов
		for i := range orderItems {
			orderItems[i].OrderID = finalOrder.ID
		}

		if err := s.orderItemsRepo.CreateItems(tx, orderItems); err != nil {
			return err
		}

		if err := s.cartRepo.ClearCart(tx, userID); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return finalOrder, nil
}
