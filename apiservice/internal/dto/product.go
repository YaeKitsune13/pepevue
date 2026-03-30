package dto

type ProductResponse struct {
	Title        string  `json:"title"`
	Image        string  `json:"image_url"`
	Description  string  `json:"description"`
	Cost         float64 `json:"cost"`
	Count        int16   `json:"count"`
	CategoryName string  `json:"category_name"`
}

type ProductCreateRequest struct {
	Title       string  `json:"title" binding:"required"`
	Image       string  `json:"image_url"`
	Description string  `json:"description" binding:"required"`
	Cost        float64 `json:"cost" binding:"required,min=0,max=100000"`
	Count       int16   `json:"count" binding:"required,min=0,max=10000"`
	CategoryID  int8    `json:"category_id" binding:"required"`
}
