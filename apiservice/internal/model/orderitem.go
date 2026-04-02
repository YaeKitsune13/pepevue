package model

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primaryKey;column:id"`
	OrderID   uint    `json:"order_id" gorm:"column:order_id"`
	Order     Order   `json:"-" gorm:"foreignKey:OrderID"`
	ProductID uint    `json:"product_id" gorm:"column:product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Count     int16   `json:"count" gorm:"column:count"`
	Cost      float64 `json:"cost" gorm:"column:cost"`
}
