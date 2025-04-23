package postgres

import (
	"HR-monitor/pkg/repository"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresStatsRepository struct{
	db *pgxpool.Pool
}

func NewPostgresStatsRepository(db *pgxpool.Pool) repository.StatsRepository {
	return &postgresStatsRepository{db: db}
}

func (r *postgresStatsRepository) GetAverageStageTime(ctx context.Context, stageID int) (float64, error) {
	query := `
	SELECT AVG(EXTRACT(EPOCH FROM (end_time - start_time)) / 3600) AS avg_hours
	FROM resume_histories
	WHERE end_time IS NOT NULL
	AND stage_id = $1
	`
	var avgHours float64
	err := r.db.QueryRow(ctx, query, stageID).Scan(&avgHours)
	if err != nil {
		return 0, err
	}
	return avgHours, nil
}

func (r *postgresStatsRepository) GetStageDistribution(ctx context.Context) (map[string]int, error) {
	query := `
	SELECT s.name, COUNT(*) AS count
	FROM resumes r
	JOIN stages s ON r.current_stage_id = s.id
	GROUP BY s.name
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	distribution := make(map[string]int)
	for rows.Next() {
		var stageName string
		var resumesCount int
		err := rows.Scan(&stageName, &resumesCount)
		if err != nil {
			return nil, err
		}
		distribution[stageName] = resumesCount
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return distribution, nil
}

func (r *postgresStatsRepository) GetSourceDistribution(ctx context.Context) (map[string]int, error) {
	query := `
	SELECT source, COUNT(*) AS count
	FROM resumes
	GROUP BY source
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	distribution := make(map[string]int)
	for rows.Next() {
		var source string
		var resumesCount int
		err := rows.Scan(&source, &resumesCount)
		if err != nil {
			return nil, err
		}
		distribution[source] = resumesCount
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return distribution, nil
}

func (r *postgresStatsRepository) GetAverageCandidatesPerVacancy(ctx context.Context) (float64, error) {
	return 0, nil
}

func (r *postgresStatsRepository) GetSLAViolationsCount(ctx context.Context) (int, error) {
	return 0, nil
}