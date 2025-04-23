package models

import (
	"HR-monitor/pkg/enums"
	"time"
)

type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role enums.UserRole `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
}

type Auth struct {
	Email string `json:"email"`
	Password string `json:"password"`
}




