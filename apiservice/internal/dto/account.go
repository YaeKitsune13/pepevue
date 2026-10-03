package dto

type AccountResponse struct {
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Surname    string `json:"surname"`
	Login      string `json:"login"`
	Role       string `json:"role"`
}

type AccountCreateRequest struct {
	Name       string `json:"name" binding:"required"`
	Patronymic string `json:"patronymic"` // в PHP-форме отчество необязательное
	Surname    string `json:"surname" binding:"required"`
	Login      string `json:"login" binding:"required"`
	Password   string `json:"password" binding:"required,min=6"`
}

type AccountLoginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AccountUpdateRequest — изменение профиля. Пустой password = не менять пароль.
type AccountUpdateRequest struct {
	Name       string `json:"name" binding:"required"`
	Patronymic string `json:"patronymic"`
	Surname    string `json:"surname" binding:"required"`
	Login      string `json:"login" binding:"required"`
	Password   string `json:"password" binding:"omitempty,min=6"`
}
