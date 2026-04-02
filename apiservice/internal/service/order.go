package service

import (
	"apiservice/internal/model"
	"apiservice/internal/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type OrderService struct {
	orderRepo      *repository.OrderRepository
	orderItemsRepo *repository.OrderItemRepository
	cartRepo       *repository.CartRepository // Добавляем, чтобы забрать товары из корзины
	db             *gorm.DB                   // Нужно для транзакций
}

func NewOrderService(ro *repository.OrderRepository, ri *repository.OrderItemRepository, rc *repository.CartRepository, db *gorm.DB) *OrderService {
	return &OrderService{
		orderRepo:      ro,
		orderItemsRepo: ri,
		cartRepo:       rc,
		db:             db,
	}
}

func (s *OrderService) PlaceOrder(userID uint, address string) (*model.Order, error) {
	var finalOrder *model.Order

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// 1. Получаем корзину
		cartItems, err := s.cartRepo.GetFullCart(userID)
		if err != nil || len(cartItems) == 0 {
			return errors.New("корзина пуста")
		}

		// 2. Считаем сумму и готовим товары
		var totalPrice float64
		var orderItems []model.OrderItem

		for _, ci := range cartItems {
			totalPrice += ci.Product.Cost * float64(ci.Count)

			orderItems = append(orderItems, model.OrderItem{
				ProductID: ci.ProductID,
				Count:     ci.Count,
				Cost:      ci.Product.Cost, // Теперь поле есть в модели!
			})
		}

		// 3. Создаем заказ
		finalOrder = &model.Order{
			UserID:       userID,
			PriceAll:     totalPrice,
			Status:       model.OrderStatusNew,
			Address:      address,
			DateDelivery: time.Now().AddDate(0, 0, 3), // Доставка через 3 дня для примера
		}

		// Важно: передаем tx в репозиторий!
		if err := s.orderRepo.Create(tx, finalOrder); err != nil {
			return err
		}

		// 4. Сохраняем товары заказа
		for i := range orderItems {
			orderItems[i].OrderID = finalOrder.ID
		}
		if err := s.orderItemsRepo.CreateItems(tx, orderItems); err != nil {
			return err
		}

		// 5. Очищаем корзину (используя tx!)
		if err := s.cartRepo.ClearCart(tx, userID); err != nil {
			return err
		}

		return nil
	})

	return finalOrder, err
}
