package repository

import (
	"fmt"

	"apiservice/internal/model"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
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

func (r *AccountRepository) GetByID(id uint) (*model.Account, error) {
	var user model.Account
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AccountRepository) Update(account *model.Account) error {
	return r.db.Save(account).Error
}

// SoftDelete помечает аккаунт удалённым (gorm DeletedAt) и освобождает логин,
// иначе unique-индекс не даст зарегистрироваться с тем же логином заново.
func (r *AccountRepository) SoftDelete(account *model.Account) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		freed := fmt.Sprintf("%s#deleted#%d", account.Login, account.ID)
		if err := tx.Model(account).Update("login", freed).Error; err != nil {
			return err
		}
		return tx.Delete(account).Error
	})
}
