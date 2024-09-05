package models

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email       string       `gorm:"size:100;not null;unique" json:"email"`
	Password    string       `gorm:"size:100;not null;" json:"password"`
	ActivatedAt sql.NullTime `gorm:"default:null" json:"activated_at"`
}
