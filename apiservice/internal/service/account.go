package service

import (
	"apiservice/internal/model"
	"apiservice/internal/repository"
	"errors"
	"unicode/utf8"
)

type AccountService struct {
	repo *repository.AccountRepository
}

func NewAccountService(repo *repository.AccountRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (serv *AccountService) AddNewAccount(account *model.Account) error {
	if utf8.RuneCountInString(account.Login) < 4 {
		return errors.New("Минимальная длинна логина 4 символа")
	}

	if utf8.RuneCountInString(account.Password) < 6 {
		return errors.New("Минимальная длина пароля 6 символов")
	}

	return serv.repo.Create(account)
}

func (serv *AccountService) GetAccountByUsernameAndPasword(account *model.Account) (*model.Account, error) {
	user, err := serv.repo.GetByLogin(account.Login)

	if err != nil {
		return nil, errors.New("Пользователь не найден")
	}

	if user.Password != account.Password {
		return nil, errors.New("Пароль не верный")
	}

	return user, nil
}
