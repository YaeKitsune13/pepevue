package repository

import (
	"apiservice/internal/model"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) Create(account *model.Account) error {
	return r.db.Create(account).Error
}

func (r *AccountRepository) GetByLogin(login string) (*model.Account, error) {
	var user model.Account
	if err := r.db.Where("login = ?", login).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
