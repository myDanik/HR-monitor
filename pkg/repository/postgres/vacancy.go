package postgres

import (
	"HR-monitor/pkg/enums"
	"HR-monitor/pkg/models"
	"HR-monitor/pkg/repository"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresVacancyRepository struct {
	db *pgxpool.Pool
}

func NewPostgresVacancyRepository(db *pgxpool.Pool) repository.VacancyRepository {
	return &postgresVacancyRepository{db: db}
}

func (r *postgresVacancyRepository) CreateVacancy(ctx context.Context, vacancy models.Vacancy) error {
	query := `
	INSERT INTO vacancies (title, description, status, created_by, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
	`
	err := r.db.QueryRow(ctx, query,
		vacancy.Title,
		vacancy.Description,
		vacancy.Status,
		vacancy.CreatedBy,
		vacancy.CreatedAt,
		vacancy.UpdatedAt,
	).Scan(&vacancy.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *postgresVacancyRepository) GetVacancyByID(ctx context.Context, id int) (models.Vacancy, error) {
	query := `
	SELECT id, title, description, status, created_by, created_at, updated_at
	FROM vacancies
	WHERE id = $1
	`
	var vacancy models.Vacancy
	err := r.db.QueryRow(ctx, query, id).Scan(
		&vacancy.ID,
		&vacancy.Title,
		&vacancy.Description,
		&vacancy.Status,
		&vacancy.CreatedBy,
		&vacancy.CreatedAt,
		&vacancy.UpdatedAt,
	)
	if err != nil {
		return models.Vacancy{}, err
	}
	return vacancy, nil
}

func (r *postgresVacancyRepository) UpdateVacancy(ctx context.Context, vacancy models.Vacancy) error {
	query := `
	UPDATE vacancies
	SET title = $2,
	description = $3,
	status = $4,
	created_by = $5,
	created_at = $6,
	updated_at = $7
	WHERE id = $1
	`
	_, err := r.db.Exec(ctx, query,
		vacancy.ID,
		vacancy.Title,
		vacancy.Description,
		vacancy.Status,
		vacancy.CreatedBy,
		vacancy.CreatedAt,
		vacancy.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil

}

func (r *postgresVacancyRepository) DeleteVacancy(ctx context.Context, id int) error {
	query := `
	DELETE FROM vacancies
	WHERE id = $1
	`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil

}

func (r *postgresVacancyRepository) ChangeVacancyStatus(ctx context.Context, id int, status enums.VacancyStatus) error {
	query := `
	UPDATE vacancies
	SET status = $2
	WHERE id = $1
	`
	_, err := r.db.Exec(ctx, query, id, status)
	if err != nil {
		return err
	}
	return nil
}
