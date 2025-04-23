package repository

import (
	"context"
	"time"
)

type StatsRepository interface {
	GetAverageStageTime(ctx context.Context, stageID int) (time.Duration, error)
	GetStageDistribution(ctx context.Context) (map[string]int, error)
	GetSourceDistribution(ctx context.Context) (map[string]int, error)
	GetAverageCandidatesPerVacancy(ctx context.Context) (float64, error)
	GetSLAViolationsCount(ctx context.Context) (int, error)
}