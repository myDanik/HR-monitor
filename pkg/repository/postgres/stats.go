package postgres

import (
	"HR-monitor/pkg/repository"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresStatsRepository struct {
	db *pgxpool.Pool
}

func NewPostgresStatsRepository(db *pgxpool.Pool) repository.StatsRepository {
	return &postgresStatsRepository{db: db}
}

func (r *postgresStatsRepository) GetAverageStageTime(ctx context.Context) (float64, error) {
	query := `
	SELECT AVG(EXTRACT(EPOCH FROM (end_time - start_time)) / 3600) AS avg_hours
	FROM resume_histories
	WHERE end_time IS NOT NULL
	GROUP BY stage_id

	`
	var avgHours float64
	err := r.db.QueryRow(ctx, query).Scan(&avgHours)
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
	query := `
	SELECT AVG(candidate_count) as avg_candidates_per_vacancy
	FROM (
		SELECT v.id, COUNT(r.id) as candidate_count
		FROM vacancies v
		LEFT JOIN resumes r ON v.id = r.vacancy_id
		GROUP BY v.id
	) as vacancy_counts
	`
	var avgCandidatesPerVacancy float64
	err := r.db.QueryRow(ctx, query).Scan(&avgCandidatesPerVacancy)
	if err != nil {
		return avgCandidatesPerVacancy, err
	}
	return avgCandidatesPerVacancy, nil

}

func (r *postgresStatsRepository) GetHistoricalSLAViolationsCount(ctx context.Context) (int, error) {
	query := `
	SELECT COUNT(*) AS historical_sla_violations_count
	FROM (
		SELECT 
			rh.resume_id,
			rh.stage_id,
			EXTRACT(EPOCH FROM (rh.end_time - rh.start_time)) / 3600 AS hours_spent,
			sr.duration_hours AS allowed_hours
		FROM resume_histories rh
		JOIN resumes r ON rh.resume_id = r.id
		JOIN sla_rules sr 
			ON rh.stage_id = sr.stage_id 
			AND r.vacancy_id = sr.vacancy_id
		WHERE 
			rh.end_time IS NOT NULL
	) AS violations
	WHERE hours_spent > allowed_hours
	`
	var slaViolationsCount int
	err := r.db.QueryRow(ctx, query).Scan(&slaViolationsCount)
	if err != nil {
		return 0, err
	}
	return slaViolationsCount, nil
}

func (r *postgresStatsRepository) GetCurrentSLAViolationsCount(ctx context.Context) (int, error) {
	// query := `
	// SELECT COUNT(*) AS current_sla_violations_count
	// FROM (
	// 	SELECT
	// 		r.id AS resume_id,
	// 		r.current_stage_id AS stage_id,
	// 		EXTRACT(EPOCH FROM (NOW() - r.updated_at)) / 3600 AS hours_spent,
	// 		sr.duration_hours AS allowed_hours
	// 	FROM resumes r
	// 	JOIN sla_rules sr ON r.current_stage_id = sr.stage_id AND r.vacancy_id = sr.vacancy_id
	// 	WHERE r.updated_at IS NOT NULL
	// ) AS current_violations
	// WHERE hours_spent > allowed_hours

	// `
	query := `
	SELECT COUNT(*) AS current_sla_violations_count
	FROM resumes r
	WHERE r.sla_deadline < NOW()
	`

	var slaViolationsCount int
	err := r.db.QueryRow(ctx, query).Scan(&slaViolationsCount)
	if err != nil {
		return 0, err
	}
	return slaViolationsCount, nil
}
