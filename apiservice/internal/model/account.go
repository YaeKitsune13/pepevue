package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Surname    string `json:"surname" gorm:"column:surname"`
	Name       string `json:"name" gorm:"column:name"`
	Patronymic string `json:"patronymic" gorm:"column:patronymic"`
	Role       string `json:"role"  gorm:"column:role;default:user"`
	Login      string `json:"login" gorm:"column:login;unique"`
	Password   string `json:"-" gorm:"column:password"`
}
