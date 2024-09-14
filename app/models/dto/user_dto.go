package dto

type UserDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type EmailDto struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordDTO struct {
	Password string `json:"password" binding:"required"`
	Token    string `json:"token" binding:"required"`
}
