package handler

import (
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

func (h *OrderHandler) PlaceOrder(c *gin.Context)  {
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Пользователь не авторизован"})
		return
	}

	var input PlaceOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Неверный формат адреса",
		}
		return
	}
	order, err := h.serv.PlaceOrder(userID.(uint), input.Address)

	if err != nil {
		if err.Error() == "корзина пуста" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":err.Error()
			})
		}
	 	c.JSON(http.StatusInternalServerError, gin.H{
				"error":"Не удалось создать заказ"
			})
			return
	}
	c.JSON(http.StatusCreated,order)
}
