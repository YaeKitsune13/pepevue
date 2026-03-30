package model

type Cart struct {
	ID        uint    `json:"id" gorm:"primaryKey;column:id"`
	UserID    uint    `json:"user_id" gorm:"column:user_id"`
	User      Account `json:"-" gorm:"foreignKey:UserID"`
	ProductID uint    `json:"product_id" gorm:"column:product_id"`
	Product   Product `json:"-" gorm:"foreignKey:ProductID"`
	Count     int16   `json:"count" gorm:"column:count"`
}
