package repository

import (
	"HR-monitor/pkg/models"
	"context"
)

type SLARepository interface {
	CreateSLARule(ctx context.Context, rule models.SLARule) error
	GetSLARulesByVacancyID(ctx context.Context, vacancyID int) ([]models.SLARule, error)
	UpdateSLARule(ctx context.Context, rule models.SLARule) error
	DeleteSLARuleByID(ctx context.Context, id int) error
}

