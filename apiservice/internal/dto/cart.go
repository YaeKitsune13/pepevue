package dto

type CartAddRequest struct {
	ProductID uint  `json:"product_id" binding:"required"`
	Count     int16 `json:"count" binding:"required,min=1"`
}

type CartUpdateRequest struct {
	ProductID uint  `json:"product_id" binding:"required"`
	Count     int16 `json:"count" binding:"required,min=1"`
}

type CartItemResponse struct {
	ProductID uint    `json:"product_id"`
	Title     string  `json:"title"`
	Image     string  `json:"image_url"`
	Cost      float64 `json:"cost"`
	Count     int16   `json:"count"`
	Total     float64 `json:"total"`
}

type CartListResponse struct {
	Items []CartItemResponse `json:"items"`
	Total float64            `json:"total"`
}
