package service

import (
	"HR-monitor/pkg/auth"
	"HR-monitor/pkg/models"
	"HR-monitor/pkg/repository"
	"context"
	"time"
)

type SLAService struct {
	SLARepo repository.SLARepository
}

func NewSLAService(SLARepo repository.SLARepository) *SLAService {
	return &SLAService{
		SLARepo: SLARepo,
	}
}

func (s *SLAService) CreateSLARule(ctx context.Context, rule models.SLARule) error {
	rule.CreatedAt = time.Now()
	rule.UpdatedAt = time.Now()
	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}
	rule.CreatedBy = userID
	return s.SLARepo.CreateSLARule(ctx, rule)
}

func (s *SLAService) GetSLARulesByVacancyID(ctx context.Context, vacancyID int) ([]models.SLARule, error) {
	return s.SLARepo.GetSLARulesByVacancyID(ctx, vacancyID)
}

func (s *SLAService) UpdateSLARule(ctx context.Context, rule models.SLARule) error {
	rule.UpdatedAt = time.Now()
	return s.SLARepo.UpdateSLARule(ctx, rule)
}

func (s *SLAService) DeleteSLARule(ctx context.Context, id int) error {
	return s.SLARepo.DeleteSLARuleByID(ctx, id)
}
