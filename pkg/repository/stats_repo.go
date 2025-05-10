package repository

import (
	"context"
)

type StatsRepository interface {
	GetAverageStageTime(ctx context.Context) (float64, error)
	GetStageDistribution(ctx context.Context) (map[string]int, error)
	GetSourceDistribution(ctx context.Context) (map[string]int, error)
	GetAverageCandidatesPerVacancy(ctx context.Context) (float64, error)
	GetHistoricalSLAViolationsCount(ctx context.Context) (int, error)
	GetCurrentSLAViolationsCount(ctx context.Context) (int, error)
}
