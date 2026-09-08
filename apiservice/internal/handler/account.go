package handler

import (
	"apiservice/internal/dto"
	"apiservice/internal/middleware"
	"apiservice/internal/model"
	"apiservice/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	serv      *service.AccountService
	jwtSecret string
}

func NewAccountHandler(serv *service.AccountService, jwtSecret string) *AccountHandler {
	return &AccountHandler{serv: serv, jwtSecret: jwtSecret}
}

// Register регистрирует нового пользователя
// @Summary      Регистрация аккаунта
// @Description  Создаёт новый аккаунт по логину и паролю
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        input  body      dto.AccountCreateRequest  true  "Данные регистрации"
// @Success      201    {object}  dto.AccountResponse
// @Failure      400    {object}  map[string]string
// @Router       /register [post]
func (h *AccountHandler) Register(c *gin.Context) {
	var req dto.AccountCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные регистрации"})
		return
	}

	account := &model.Account{
		Name:       req.Name,
		Patronymic: req.Patronymic,
		Surname:    req.Surname,
		Login:      req.Login,
		Password:   req.Password,
	}

	if err := h.serv.AddNewAccount(account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.AccountResponse{
		Name: account.Name, Patronymic: account.Patronymic,
		Surname: account.Surname, Login: account.Login,
	})
}

// Login аутентифицирует пользователя и выдаёт cookie-сессию
// @Summary      Вход в аккаунт
// @Description  Проверяет логин/пароль и устанавливает httpOnly cookie с JWT
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        input  body      dto.AccountLoginRequest  true  "Данные для входа"
// @Success      200    {object}  dto.AccountResponse
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Router       /login [post]
func (h *AccountHandler) Login(c *gin.Context) {
	var req dto.AccountLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	account, err := h.serv.GetAccountByUsernameAndPasword(&model.Account{
		Login:    req.Login,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := middleware.GenerateToken(account.ID, h.jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать сессию"})
		return
	}

	// secure=false для локальной разработки без https; в проде — true
	c.SetCookie("session_token", token, 60*60*24, "/", "", false, true)

	c.JSON(http.StatusOK, dto.AccountResponse{
		Name: account.Name, Patronymic: account.Patronymic,
		Surname: account.Surname, Login: account.Login,
	})
}

// Logout завершает сессию пользователя
// @Summary      Выход из аккаунта
// @Description  Удаляет cookie-сессию
// @Tags         account
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /logout [post]
// @Security     CookieAuth
func (h *AccountHandler) Logout(c *gin.Context) {
	c.SetCookie("session_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Вы вышли из аккаунта"})
}
