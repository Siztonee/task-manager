package dto

type CreateAuthRequest struct {
	Username string `json:"username" binding:"required,max=32"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}
