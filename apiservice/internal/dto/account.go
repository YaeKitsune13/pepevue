package dto

type AccountResponse struct {
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Surname    string `json:"surname"`
	Login      string `json:"login"`
}

type AccountCreateRequest struct {
	Name       string `json:"name" binding:"required"`
	Patronymic string `json:"patronymic" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Login      string `json:"login" binding:"required"`
	Password   string `json:"password" binding:"required,min=6"`
}

type AccountLoginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
