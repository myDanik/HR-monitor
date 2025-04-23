package repository

import (
	"HR-monitor/pkg/models"
	"context"
)

type ResumeRepository interface {
	CreateResume(ctx context.Context, resume models.Resume) error
	GetResumeByID(ctx context.Context, id int) (models.Resume, error)
	GetResumes(ctx context.Context, filters models.ResumeFilters, sort models.ResumeSort) ([]models.Resume, error)
	UpdateResumeByID(ctx context.Context, resume models.Resume) error
	DeleteResumeByID(ctx context.Context, id int) error
	MoveResumeToStage(ctx context.Context, resumeID, stageID int) error
}

