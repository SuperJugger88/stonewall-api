package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	//Phone     uint64           `json:"phone"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Verified  bool      `json:"verified"`
}
