package service

import (
	"HR-monitor/pkg/auth"
	"HR-monitor/pkg/enums"
	"HR-monitor/pkg/models"
	"HR-monitor/pkg/repository"
	"context"
	"time"
)

type VacancyService struct {
	vacancyRepo repository.VacancyRepository
}

func NewVacancyService(vacancyRepo repository.VacancyRepository) *VacancyService {
	return &VacancyService{
		vacancyRepo: vacancyRepo,
	}
}

func (s *VacancyService) CreateVacancy(ctx context.Context, vacancy models.Vacancy) error {
	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}

	vacancy.CreatedBy = userID
	vacancy.CreatedAt = time.Now()
	vacancy.UpdatedAt = time.Now()
	return s.vacancyRepo.CreateVacancy(ctx, vacancy)
}

func (s *VacancyService) GetVacancyByID(ctx context.Context, id int) (models.Vacancy, error) {
	return s.vacancyRepo.GetVacancyByID(ctx, id)
}

func (s *VacancyService) UpdateVacancy(ctx context.Context, vacancy models.Vacancy) error {
	vacancy.UpdatedAt = time.Now()
	return s.vacancyRepo.UpdateVacancy(ctx, vacancy)
}

func (s *VacancyService) DeleteVacancy(ctx context.Context, id int) error {
	return s.vacancyRepo.DeleteVacancy(ctx, id)
}

func (s *VacancyService) ChangeVacancyStatus(ctx context.Context, id int, status enums.VacancyStatus) error {
	return s.vacancyRepo.ChangeVacancyStatus(ctx, id, status)
}
