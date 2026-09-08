package service

import (
	"apiservice/internal/model"
	"apiservice/internal/repository"
	"errors"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

type AccountService struct {
	repo *repository.AccountRepository
}

func NewAccountService(repo *repository.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (serv *AccountService) AddNewAccount(account *model.Account) error {
	if utf8.RuneCountInString(account.Login) < 4 {
		return errors.New("Минимальная длина логина 4 символа")
	}
	if utf8.RuneCountInString(account.Password) < 6 {
		return errors.New("Минимальная длина пароля 6 символов")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Не удалось обработать пароль")
	}
	account.Password = string(hashed)

	return serv.repo.Create(account)
}

func (serv *AccountService) GetAccountByUsernameAndPasword(input *model.Account) (*model.Account, error) {
	user, err := serv.repo.GetByLogin(input.Login)
	if err != nil {
		return nil, errors.New("Пользователь не найден")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, errors.New("Пароль не верный")
	}

	return user, nil
}
