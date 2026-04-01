package repository

import (
	"apiservice/internal/model" // Импорт твоих моделей

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB // Используем указатель
}

// NewAccountRepository — это "конструктор".
// Мы передаем сюда соединение с БД, которое создадим в main.go
func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

// Пример метода для создания аккаунта
func (r *AccountRepository) Create(account *model.Account) error {
	// GORM сам сделает SQL запрос INSERT
	return r.db.Create(account).Error
}
