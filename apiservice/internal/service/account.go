package service

import (
	"errors"
	"unicode/utf8"

	"apiservice/internal/model"
	"apiservice/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrLoginTooShort      = errors.New("минимальная длина логина 4 символа")
	ErrPasswordTooShort   = errors.New("минимальная длина пароля 6 символов")
	ErrLoginTaken         = errors.New("логин уже занят")
	ErrInvalidCredentials = errors.New("неверный логин или пароль")
	ErrAccountNotFound    = errors.New("аккаунт не найден")
)

type AccountService struct {
	repo *repository.AccountRepository
}

func NewAccountService(repo *repository.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

// ProfileUpdate — данные для изменения профиля. Пустой Password = не менять.
type ProfileUpdate struct {
	Surname, Name, Patronymic, Login, Password string
}

func hashPassword(password string) (string, error) {
	if utf8.RuneCountInString(password) < 6 {
		return "", ErrPasswordTooShort
	}
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(h), nil
}

func (s *AccountService) loginExists(login string) (bool, error) {
	_, err := s.repo.GetByLogin(login)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return false, err
}

func (s *AccountService) AddNewAccount(account *model.Account) error {
	if utf8.RuneCountInString(account.Login) < 4 {
		return ErrLoginTooShort
	}
	exists, err := s.loginExists(account.Login)
	if err != nil {
		return err
	}
	if exists {
		return ErrLoginTaken
	}
	hashed, err := hashPassword(account.Password)
	if err != nil {
		return err
	}
	account.Password = hashed
	return s.repo.Create(account)
}

// Authenticate проверяет логин/пароль. Причину отказа наружу не раскрываем.
func (s *AccountService) Authenticate(login, password string) (*model.Account, error) {
	user, err := s.repo.GetByLogin(login)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}
	return user, nil
}

func (s *AccountService) GetByID(id uint) (*model.Account, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}
	return user, nil
}

func (s *AccountService) UpdateProfile(id uint, in ProfileUpdate) (*model.Account, error) {
	acc, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	if utf8.RuneCountInString(in.Login) < 4 {
		return nil, ErrLoginTooShort
	}
	if in.Login != acc.Login {
		exists, err := s.loginExists(in.Login)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, ErrLoginTaken
		}
		acc.Login = in.Login
	}
	acc.Surname, acc.Name, acc.Patronymic = in.Surname, in.Name, in.Patronymic

	if in.Password != "" {
		hashed, err := hashPassword(in.Password)
		if err != nil {
			return nil, err
		}
		acc.Password = hashed
	}
	if err := s.repo.Update(acc); err != nil {
		return nil, err
	}
	return acc, nil
}

func (s *AccountService) DeleteAccount(id uint) error {
	acc, err := s.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.SoftDelete(acc)
}
