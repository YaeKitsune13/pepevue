package handler

import (
	"errors"
	"net/http"

	"apiservice/internal/dto"
	"apiservice/internal/middleware"
	"apiservice/internal/model"
	"apiservice/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	sessionCookie = "session_token"
	sessionMaxAge = 60 * 60 * 24
)

type AccountHandler struct {
	serv      *service.AccountService
	jwtSecret string
}

func NewAccountHandler(serv *service.AccountService, jwtSecret string) *AccountHandler {
	return &AccountHandler{serv: serv, jwtSecret: jwtSecret}
}

func toAccountResponse(a *model.Account) dto.AccountResponse {
	return dto.AccountResponse{
		Name: a.Name, Patronymic: a.Patronymic, Surname: a.Surname,
		Login: a.Login, Role: a.Role,
	}
}

func (h *AccountHandler) setSession(c *gin.Context, userID uint) error {
	token, err := middleware.GenerateToken(userID, h.jwtSecret)
	if err != nil {
		return err
	}
	c.SetSameSite(http.SameSiteLaxMode)
	// secure=false для локальной разработки без https; в проде — true
	c.SetCookie(sessionCookie, token, sessionMaxAge, "/", "", false, true)
	return nil
}

func clearSession(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(sessionCookie, "", -1, "/", "", false, true)
}

// writeAccountError превращает ошибки сервиса в HTTP-ответы.
func writeAccountError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, service.ErrLoginTaken):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, service.ErrLoginTooShort), errors.Is(err, service.ErrPasswordTooShort):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, service.ErrAccountNotFound):
		clearSession(c)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	default:
		internalError(c, err)
	}
}

// Register регистрирует нового пользователя
// @Summary      Регистрация аккаунта
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        input  body      dto.AccountCreateRequest  true  "Данные регистрации"
// @Success      201    {object}  dto.AccountResponse
// @Failure      400    {object}  map[string]string
// @Failure      409    {object}  map[string]string
// @Router       /register [post]
func (h *AccountHandler) Register(c *gin.Context) {
	var req dto.AccountCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные регистрации"})
		return
	}

	account := &model.Account{
		Name: req.Name, Patronymic: req.Patronymic, Surname: req.Surname,
		Login: req.Login, Password: req.Password,
	}
	if err := h.serv.AddNewAccount(account); err != nil {
		writeAccountError(c, err)
		return
	}
	c.JSON(http.StatusCreated, toAccountResponse(account))
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

	account, err := h.serv.Authenticate(req.Login, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		internalError(c, err)
		return
	}

	if err := h.setSession(c, account.ID); err != nil {
		internalError(c, err)
		return
	}
	c.JSON(http.StatusOK, toAccountResponse(account))
}

// Logout завершает сессию пользователя
// @Summary      Выход из аккаунта
// @Tags         account
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /logout [post]
// @Security     CookieAuth
func (h *AccountHandler) Logout(c *gin.Context) {
	clearSession(c)
	c.JSON(http.StatusOK, gin.H{"message": "Вы вышли из аккаунта"})
}

// Me возвращает профиль текущего пользователя
// @Summary      Профиль
// @Description  Данные текущего пользователя (Vue использует для проверки сессии)
// @Tags         account
// @Produce      json
// @Success      200  {object}  dto.AccountResponse
// @Failure      401  {object}  map[string]string
// @Router       /me [get]
// @Security     CookieAuth
func (h *AccountHandler) Me(c *gin.Context) {
	userID, ok := currentUserID(c)
	if !ok {
		return
	}
	account, err := h.serv.GetByID(userID)
	if err != nil {
		writeAccountError(c, err)
		return
	}
	c.JSON(http.StatusOK, toAccountResponse(account))
}

// UpdateMe изменяет профиль
// @Summary      Изменить профиль
// @Description  Пустой password — пароль не меняется
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        input  body      dto.AccountUpdateRequest  true  "Новые данные"
// @Success      200    {object}  dto.AccountResponse
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Failure      409    {object}  map[string]string
// @Router       /me [put]
// @Security     CookieAuth
func (h *AccountHandler) UpdateMe(c *gin.Context) {
	userID, ok := currentUserID(c)
	if !ok {
		return
	}
	var req dto.AccountUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные профиля"})
		return
	}

	account, err := h.serv.UpdateProfile(userID, service.ProfileUpdate{
		Surname: req.Surname, Name: req.Name, Patronymic: req.Patronymic,
		Login: req.Login, Password: req.Password,
	})
	if err != nil {
		writeAccountError(c, err)
		return
	}
	c.JSON(http.StatusOK, toAccountResponse(account))
}

// DeleteMe удаляет (деактивирует) аккаунт
// @Summary      Удалить аккаунт
// @Tags         account
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /me [delete]
// @Security     CookieAuth
func (h *AccountHandler) DeleteMe(c *gin.Context) {
	userID, ok := currentUserID(c)
	if !ok {
		return
	}
	if err := h.serv.DeleteAccount(userID); err != nil {
		writeAccountError(c, err)
		return
	}
	clearSession(c)
	c.JSON(http.StatusOK, gin.H{"message": "Аккаунт удалён"})
}
