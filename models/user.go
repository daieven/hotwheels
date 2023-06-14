package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null default '' index VARCHAR(50) UNIQUE username"`
	Password string `gorm:"not null default '' VARCHAR(100) password"`
	IsAdmin  bool   `gorm:"not null default false BOOL is_admin"`
}
