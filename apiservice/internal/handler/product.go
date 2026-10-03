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

type ProductHandler struct {
	serv *service.ProductService
}

func NewProductHandler(serv *service.ProductService) *ProductHandler {
	return &ProductHandler{serv: serv}
}

func toProductResponse(p *model.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID: p.ID, Title: p.Title, Image: p.Image, Description: p.Description,
		Cost: p.Cost, Count: p.Count, CategoryName: p.Category.Name,
	}
}

// List возвращает каталог товаров
// @Summary      Список товаров
// @Tags         product
// @Produce      json
// @Success      200  {array}   dto.ProductResponse
// @Failure      500  {object}  map[string]string
// @Router       /products [get]
func (h *ProductHandler) List(c *gin.Context) {
	products, err := h.serv.GetListProducts()
	if err != nil {
		internalError(c, err)
		return
	}
	resp := make([]dto.ProductResponse, 0, len(products))
	for i := range products {
		resp = append(resp, toProductResponse(&products[i]))
	}
	c.JSON(http.StatusOK, resp)
}

// Get возвращает один товар
// @Summary      Карточка товара
// @Tags         product
// @Produce      json
// @Param        id   path      int  true  "ID товара"
// @Success      200  {object}  dto.ProductResponse
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /products/{id} [get]
func (h *ProductHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный id товара"})
		return
	}
	product, err := h.serv.GetProduct(uint(id))
	if err != nil {
		if errors.Is(err, service.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
			return
		}
		internalError(c, err)
		return
	}
	c.JSON(http.StatusOK, toProductResponse(product))
}
