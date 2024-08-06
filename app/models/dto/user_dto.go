package dto

type UserDTO struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}
