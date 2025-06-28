package service

import (
	"context"
	"time"

	"HR-monitor/pkg/auth"
	"HR-monitor/pkg/models"
	"HR-monitor/pkg/repository"
)

type ResumeService struct {
	resumeRepo repository.ResumeRepository
	statsRepo  repository.StatsRepository
	SLARepo    repository.SLARepository
}

func NewResumeService(resumeRepo repository.ResumeRepository, statsRepo repository.StatsRepository, SLARepo repository.SLARepository) *ResumeService {
	return &ResumeService{
		resumeRepo: resumeRepo,
		statsRepo:  statsRepo,
		SLARepo:    SLARepo,
	}
}

func (s *ResumeService) CreateResume(ctx context.Context, resume models.Resume) error {
	resume.CurrentStageID = 1
	resume.CreatedAt = time.Now()
	resume.UpdatedAt = time.Now()

	slaRule, err := s.SLARepo.GetSLARuleByStageAndVacancy(ctx, resume.CurrentStageID, resume.VacancyID)
	if err == repository.ErrNotFound {
		resume.SLADeadline = time.Now().Add(24 * time.Hour)
	} else if err != nil {
		return err
	} else {
		resume.SLADeadline = time.Now().Add(time.Duration(slaRule.DurationHours) * time.Hour)
	}

	return s.resumeRepo.CreateResume(ctx, resume)
}

func (s *ResumeService) GetResumeByID(ctx context.Context, id int) (models.Resume, error) {
	return s.resumeRepo.GetResumeByID(ctx, id)
}

func (s *ResumeService) GetResumes(ctx context.Context, filters models.ResumeFilters, sort models.ResumeSort) ([]models.Resume, error) {
	return s.resumeRepo.GetResumes(ctx, filters, sort)
}

func (s *ResumeService) MoveResumeToStage(ctx context.Context, resumeID, stageID int) error {
	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}
	resume, err := s.GetResumeByID(ctx, resumeID)
	if err != nil {
		return err
	}
	err = s.resumeRepo.AddResumeHistory(ctx, resume, userID)
	if err != nil {
		return err
	}

	slaRule, err := s.SLARepo.GetSLARuleByStageAndVacancy(ctx, stageID, resume.VacancyID)
	if err == repository.ErrNotFound {
		resume.SLADeadline = time.Now().Add(24 * time.Hour)
	} else if err != nil {
		return err
	} else {
		resume.SLADeadline = time.Now().Add(time.Duration(slaRule.DurationHours) * time.Hour)
	}

	resume.CurrentStageID = stageID
	resume.UpdatedAt = time.Now()

	return s.resumeRepo.UpdateResumeByID(ctx, resume)
}

func (s *ResumeService) GetResumeStats(ctx context.Context) (*models.ResumeStats, error) {
	avgStageTime, err := s.statsRepo.GetAverageStageTime(ctx)
	if err != nil {
		return nil, err
	}
	stageDistribution, err := s.statsRepo.GetStageDistribution(ctx)
	if err != nil {
		return nil, err
	}
	sourceDistribution, err := s.statsRepo.GetSourceDistribution(ctx)
	if err != nil {
		return nil, err
	}
	avgCandidates, err := s.statsRepo.GetAverageCandidatesPerVacancy(ctx)
	if err != nil {
		return nil, err
	}
	slaViolations, err := s.statsRepo.GetCurrentSLAViolationsCount(ctx)
	if err != nil {
		return nil, err
	}

	return &models.ResumeStats{
		AverageStageTime:            avgStageTime,
		StageDistribution:           stageDistribution,
		SourceDistribution:          sourceDistribution,
		AverageCandidatesPerVacancy: avgCandidates,
		SLAViolationsCount:          slaViolations,
	}, nil
}

func (s *ResumeService) UpdateResume(ctx context.Context, resume models.Resume) error {
	resume.UpdatedAt = time.Now()
	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}
	err = s.resumeRepo.AddResumeHistory(ctx, resume, userID)
	if err != nil {
		return err
	}
	return s.resumeRepo.UpdateResumeByID(ctx, resume)
}
