package repository

import (
	"HR-monitor/pkg/models"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByID(ctx context.Context, id int) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	DeleteUserByID(ctx context.Context, id int) error
}

