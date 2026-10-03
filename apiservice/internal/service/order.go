package service

import (
	"errors"
	"fmt"
	"time"

	"apiservice/internal/model"
	"apiservice/internal/repository"

	"gorm.io/gorm"
)

type OrderService struct {
	orderRepo      *repository.OrderRepository
	orderItemsRepo *repository.OrderItemRepository
	cartRepo       *repository.CartRepository
	productRepo    *repository.ProductRepository
	db             *gorm.DB
}

func NewOrderService(
	ro *repository.OrderRepository,
	ri *repository.OrderItemRepository,
	rc *repository.CartRepository,
	rp *repository.ProductRepository,
	db *gorm.DB,
) *OrderService {
	return &OrderService{orderRepo: ro, orderItemsRepo: ri, cartRepo: rc, productRepo: rp, db: db}
}

var ErrEmptyCart = errors.New("корзина пуста")

// PlaceOrder оформляет заказ из корзины, списывает остатки и чистит корзину —
// всё в одной транзакции. Нулевая deliveryDate = через 3 дня.
func (s *OrderService) PlaceOrder(userID uint, address string, deliveryDate time.Time) (*model.Order, error) {
	if deliveryDate.IsZero() {
		deliveryDate = time.Now().AddDate(0, 0, 3)
	} else {
		y, m, d := time.Now().Date()
		if deliveryDate.Before(time.Date(y, m, d+1, 0, 0, 0, 0, time.Local)) {
			return nil, ErrInvalidDeliveryDate
		}
	}

	var finalOrder *model.Order

	err := s.db.Transaction(func(tx *gorm.DB) error {
		cartItems, err := s.cartRepo.GetFullCartTx(tx, userID)
		if err != nil {
			return err
		}
		if len(cartItems) == 0 {
			return ErrEmptyCart
		}

		var totalPrice float64
		orderItems := make([]model.OrderItem, 0, len(cartItems))

		for _, ci := range cartItems {
			ok, err := s.productRepo.DecreaseStock(tx, ci.ProductID, ci.Count)
			if err != nil {
				return err
			}
			if !ok {
				return fmt.Errorf("%w: %s", ErrNotEnoughStock, ci.Product.Title)
			}

			totalPrice += ci.Product.Cost * float64(ci.Count)
			orderItems = append(orderItems, model.OrderItem{
				ProductID: ci.ProductID,
				Count:     ci.Count,
				Cost:      ci.Product.Cost,
			})
		}

		finalOrder = &model.Order{
			UserID:       userID,
			PriceAll:     totalPrice,
			Status:       model.OrderStatusNew,
			Address:      address,
			DateDelivery: deliveryDate,
		}
		if err := s.orderRepo.Create(tx, finalOrder); err != nil {
			return err
		}

		for i := range orderItems {
			orderItems[i].OrderID = finalOrder.ID
		}
		if err := s.orderItemsRepo.CreateItems(tx, orderItems); err != nil {
			return err
		}

		return s.cartRepo.ClearCart(tx, userID)
	})
	if err != nil {
		return nil, err
	}
	return finalOrder, nil
}
