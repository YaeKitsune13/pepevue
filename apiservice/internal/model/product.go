package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string   `json:"title" gorm:"column:title"`
	Image       string   `json:"image_url" gorm:"column:image"`
	Description string   `json:"description" gorm:"column:description"`
	Cost        float64  `json:"cost" gorm:"column:cost"`
	Count       int16    `json:"count" gorm:"column:count"`
	CategoryID  int8     `json:"category_id" gorm:"column:category_id"`
	Category    Category `json:"-" gorm:"foreignKey:CategoryID"`
}
