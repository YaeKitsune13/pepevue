package model

type Category struct {
	Id   int8   `json:"id" gorm:"primaryKey;column:id"`
	Name string `json:"name" gorm:"name"`
}
