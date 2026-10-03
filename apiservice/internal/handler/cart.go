package handler

import (
	"errors"
	"net/http"
	"strconv"

	"apiservice/internal/dto"
	"apiservice/internal/model"
	"apiservice/internal/service"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	serv *service.CartService
}

func NewCartHandler(serv *service.CartService) *CartHandler {
	return &CartHandler{serv: serv}
}

func buildCartResponse(items []model.Cart) dto.CartListResponse {
	resp := dto.CartListResponse{Items: make([]dto.CartItemResponse, 0, len(items))}
	for _, it := range items {
		total := it.Product.Cost * float64(it.Count)
		resp.Total += total
		resp.Items = append(resp.Items, dto.CartItemResponse{
			ProductID: it.ProductID,
			Title:     it.Product.Title,
			Image:     it.Product.Image,
			Cost:      it.Product.Cost,
			Count:     it.Count,
			Total:     total,
		})
	}
	return resp
}

func writeCartError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, service.ErrProductNotFound), errors.Is(err, service.ErrCartItemNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, service.ErrInvalidCount):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, service.ErrNotEnoughStock):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	default:
		internalError(c, err)
	}
}

// respondCart отдаёт актуальную корзину — Vue может просто заменить state.
func (h *CartHandler) respondCart(c *gin.Context, userID uint) {
	items, err := h.serv.GetCart(userID)
	if err != nil {
		internalError(c, err)
		return
	}
	c.JSON(http.StatusOK, buildCartResponse(items))
}

// Get возвращает корзину текущего пользователя
// @Summary      Корзина
// @Tags         cart
// @Produce      json
// @Success      200  {object}  dto.CartListResponse
// @Failure      401  {object}  map[string]string
// @Router       /cart [get]
// @Security     CookieAuth
func (h *CartHandler) Get(c *gin.Context) {
	userID, ok := currentUserID(c)
	if !ok {
		return
	}
	h.respondCart(c, userID)
}

// Add добавляет товар в корзину (или увеличивает количество)
// @Summary      Добавить в корзину
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        input  body      dto.CartAddRequest  true  "Товар и количество"
// @Success      200    {object}  dto.CartListResponse
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Failure      404    {object}  map[string]string
// @Failure      409    {object}  map[string]string
// @Router       /cart [post]
// @Security     CookieAuth
func (h *CartHandler) Add(c *gin.Context) {
	userID, ok := currentUserID(c)
	if !ok {
		return
	}
	var req dto.CartAddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: нужны product_id и count >= 1"})
		return
	}
	if err := h.serv.AddProduct(userID, req.ProductID, req.Count); err != nil {
		writeCartError(c, err)
		return
	}
	h.respondCart(c, userID)
}

// Update задаёт точное количество товара в корзине (кнопки − / +)
// @Summary      Изменить количество
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        input  body      dto.CartUpdateRequest  true  "Товар и новое количество"
// @Success      200    {object}  dto.CartListResponse
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Failure      404    {object}  map[string]string
// @Failure      409    {object}  map[string]string
// @Router       /cart [put]
// @Security     CookieAuth
func (h *CartHandler) Update(c *gin.Context) {
	userID, ok := currentUserID(c)
	if !ok {
		return
	}
	var req dto.CartUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: нужны product_id и count >= 1"})
		return
	}
	if err := h.serv.SetCount(userID, req.ProductID, req.Count); err != nil {
		writeCartError(c, err)
		return
	}
	h.respondCart(c, userID)
}

// Remove удаляет позицию из корзины
// @Summary      Удалить из корзины
// @Tags         cart
// @Produce      json
// @Param        product_id  path      int  true  "ID товара"
// @Success      200         {object}  dto.CartListResponse
// @Failure      400         {object}  map[string]string
// @Failure      401         {object}  map[string]string
// @Router       /cart/{product_id} [delete]
// @Security     CookieAuth
func (h *CartHandler) Remove(c *gin.Context) {
	userID, ok := currentUserID(c)
	if !ok {
		return
	}
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный id товара"})
		return
	}
	if err := h.serv.RemoveProduct(userID, uint(productID)); err != nil {
		writeCartError(c, err)
		return
	}
	h.respondCart(c, userID)
}
