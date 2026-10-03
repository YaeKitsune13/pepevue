package service

import "errors"

var (
	ErrProductNotFound     = errors.New("продукт не найден")
	ErrInvalidCount        = errors.New("количество должно быть больше 0")
	ErrNotEnoughStock      = errors.New("недостаточно товара на складе")
	ErrCartItemNotFound    = errors.New("товара нет в корзине")
	ErrInvalidDeliveryDate = errors.New("дата доставки должна быть не раньше завтрашнего дня")
)
