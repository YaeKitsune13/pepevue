package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"apiservice/internal/dto"
	"apiservice/internal/service"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	serv *service.OrderService
}

func NewOrderHandler(serv *service.OrderService) *OrderHandler {
	return &OrderHandler{serv: serv}
}

type PlaceOrderInput struct {
	Address string `json:"address" binding:"required"`
	// DateDelivery в формате ГГГГ-ММ-ДД (как <input type="date">). Необязательно: по умолчанию +3 дня.
	DateDelivery string `json:"date_delivery"`
}

// PlaceOrder оформляет заказ из текущей корзины пользователя
// @Summary      Оформить заказ
// @Description  Создаёт заказ из корзины, списывает остатки со склада и очищает корзину
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        input  body      handler.PlaceOrderInput  true  "Адрес и дата доставки"
// @Success      201    {object}  dto.OrderResponse
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Failure      409    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /order [post]
// @Security     CookieAuth
func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	userID, ok := currentUserID(c)
	if !ok {
		return
	}

	var input PlaceOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат адреса"})
		return
	}

	var delivery time.Time
	if input.DateDelivery != "" {
		d, err := time.ParseInLocation("2006-01-02", input.DateDelivery, time.Local)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат даты, нужен ГГГГ-ММ-ДД"})
			return
		}
		delivery = d
	}

	order, err := h.serv.PlaceOrder(userID, input.Address, delivery)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrEmptyCart), errors.Is(err, service.ErrInvalidDeliveryDate):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errors.Is(err, service.ErrNotEnoughStock):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			internalError(c, err)
		}
		return
	}

	c.JSON(http.StatusCreated, dto.OrderResponse{
		ID:           order.ID,
		PriceAll:     order.PriceAll,
		UserID:       order.UserID,
		Status:       order.Status,
		Address:      order.Address,
		DateDelivery: order.DateDelivery,
	})
}

// GetByUserID получает список заказов пользователя по user_id
// @Summary      Получить заказы пользователя
// @Description  Возвращает список всех заказов указанного пользователя
// @Tags         order
// @Produce      json
// @Param        user_id  path      int  true  "ID пользователя"
// @Success      200      {array}   dto.OrderResponse
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /orders/get/{user_id} [get]
// @Security     CookieAuth
func (h *OrderHandler) GetByUserID(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil || userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID пользователя"})
		return
	}

	orders, err := h.serv.GetOrdersByUserID(uint(userID))
	if err != nil {
		internalError(c, err)
		return
	}

	resp := make([]dto.OrderResponse, 0, len(orders))
	for _, o := range orders {
		resp = append(resp, dto.OrderResponse{
			ID:           o.ID,
			PriceAll:     o.PriceAll,
			UserID:       o.UserID,
			Status:       o.Status,
			Address:      o.Address,
			DateDelivery: o.DateDelivery,
		})
	}

	c.JSON(http.StatusOK, orders)
}
