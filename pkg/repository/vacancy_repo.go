package repository

import (
	"HR-monitor/pkg/enums"
	"HR-monitor/pkg/models"
	"context"
)

type VacancyRepository interface {
	CreateVacancy(ctx context.Context, vacancy models.Vacancy) error
	GetVacancyByID(ctx context.Context, id int) (models.Vacancy, error)
	UpdateVacancy(ctx context.Context, vacancy models.Vacancy) error
	DeleteVacancy(ctx context.Context, id int) error
	ChangeVacancyStatus(ctx context.Context, id int, status enums.VacancyStatus) error
}
