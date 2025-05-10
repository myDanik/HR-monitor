package models

import (
	"HR-monitor/pkg/enums"
	"time"
)

type User struct {
	ID           int            `json:"id" db:"id"`
	Email        string         `json:"email" db:"email"`
	PasswordHash string         `json:"password_hash" db:"password_hash"`
	Role         enums.UserRole `json:"role" db:"role"`
	CreatedAt    time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" db:"updated_at"`
	Name         string         `json:"name" db:"name"`
}

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
