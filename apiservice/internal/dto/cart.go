package dto

type CartResponse struct {
	UserID    uint  `json:"user_id"`
	ProductID uint  `json:"product_id" `
	Count     int16 `json:"count"`
}

type CartUpdateRequest struct {
	ProductID uint  `json:"product_id" binding:"required"`
	Count     int16 `json:"count" binding:"required"`
}
