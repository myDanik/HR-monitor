package models

import (
	"HR-monitor/pkg/enums"
)

type RegisterRequest struct {
	Email    string
	Password string
	Name     string
	Role     enums.UserRole
}

type LoginRequest struct {
	Email    string
	Password string
}

type AuthResponse struct {
	AccessToken  string
	RefreshToken string
	User         User
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
