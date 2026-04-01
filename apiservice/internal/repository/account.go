package repository

import (
	"apiservice/internal/dto"
	"apiservice/internal/model"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) GetAll() ([]model.Account, error) {
	var accounts []model.Account
	result := r.db.Find(&accounts)
	return accounts, result.Error
}

func (r *AccountRepository) GetAccount(account *dto.AccountLoginRequest) (model.Account, error) {
	var UserAccount model.Account
	result := r.db.Where(model.Account{
		Login:    account.Login,
		Password: account.Password,
	}).Take(&UserAccount)
	return UserAccount, result.Error
}

func (r *AccountRepository) Create(account *model.Account) error {
	result := r.db.Create(account)
	return result.Error
}
