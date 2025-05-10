package postgres

import (
	"HR-monitor/pkg/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresSLARepository struct {
	db *pgxpool.Pool
}

func NewPostgresSLARepository(db *pgxpool.Pool) *postgresSLARepository {
	return &postgresSLARepository{db: db}
}

func (r *postgresVacancyRepository) CreateSLARule(ctx context.Context, rule models.SLARule) error {
	query := `
	INSERT INTO sla_rules (vacancy_id, stage_id, duration_hours, created_by, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
	`
	err := r.db.QueryRow(ctx, query,
		rule.VacancyID,
		rule.StageID,
		rule.DurationHours,
		rule.CreatedBy,
		rule.CreatedAt,
		rule.UpdatedAt,
	).Scan(&rule.ID)
	if err != nil {
		return err
	}
	return nil

}

func (r *postgresSLARepository) GetSLARulesByVacancyID(ctx context.Context, vacancyID int) ([]models.SLARule, error) {
	query := `
	SELECT id, vacancy_id, stage_id, duration_hours, created_by, created_at, updated_at
	FROM sla_rules
	WHERE vacancy_id = $1
	`
	rows, err := r.db.Query(ctx, query, vacancyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []models.SLARule
	for rows.Next() {
		var rule models.SLARule
		err := rows.Scan(
			&rule.ID,
			&rule.VacancyID,
			&rule.StageID,
			&rule.DurationHours,
			&rule.CreatedBy,
			&rule.CreatedAt,
			&rule.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return rules, nil

}
