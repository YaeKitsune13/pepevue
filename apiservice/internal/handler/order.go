package handler

import (
	"apiservice/internal/dto"
	"apiservice/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	serv *service.OrderService
}

func NewOrderHandler(serv *service.OrderService) *OrderHandler {
	return &OrderHandler{
		serv: serv,
	}
}

type PlaceOrderInput struct {
	Address string `json:"address" binding:"required"`
}

// PlaceOrder оформляет заказ из текущей корзины пользователя
// @Summary      Оформить заказ
// @Description  Создаёт заказ из содержимого корзины авторизованного пользователя и очищает корзину
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        input  body      handler.PlaceOrderInput  true  "Адрес доставки"
// @Success      201    {object}  dto.OrderResponse
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /order [post]
// @Security     CookieAuth
func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	var input PlaceOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный формат адреса",
		})
		return
	}

	order, err := h.serv.PlaceOrder(userID.(uint), input.Address)
	if err != nil {
		if err.Error() == "корзина пуста" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Не удалось создать заказ",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.OrderResponse{
		PriceAll:     order.PriceAll,
		UserID:       order.UserID,
		Status:       order.Status,
		Address:      order.Address,
		DateDelivery: order.DateDelivery,
	})
}
