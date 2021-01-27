package dto

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,max=20,min=6"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
