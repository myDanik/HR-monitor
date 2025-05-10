package postgres

import (
	"HR-monitor/pkg/models"
	"HR-monitor/pkg/repository"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresResumeRepository struct {
	db *pgxpool.Pool
}

func NewPostgresResumeRepository(db *pgxpool.Pool) repository.ResumeRepository {
	return &postgresResumeRepository{db: db}
}

func (r *postgresResumeRepository) CreateResume(ctx context.Context, resume models.Resume) error {
	query := `
	INSERT INTO resumes (vacancy_id, current_stage_id, candidate_name, candidate_contact, source, description, created_at, updated_at, sladeadline)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id
	`
	err := r.db.QueryRow(ctx, query,
		resume.VacancyID,
		resume.CurrentStageID,
		resume.CandidateName,
		resume.CandidateContact,
		resume.Source,
		resume.Description,
		resume.CreatedAt,
		resume.UpdatedAt,
		resume.SLADeadline,
	).Scan(&resume.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *postgresResumeRepository) GetResumeByID(ctx context.Context, id int) (models.Resume, error) {
	query := `
	SELECT id, vacancy_id, current_stage_id, candidate_name, candidate_contact, source, description, created_at, updated_at, sladeadline
	FROM resumes
	WHERE id = $1
	`
	var resume models.Resume
	err := r.db.QueryRow(ctx, query, id).Scan(
		&resume.ID,
		&resume.VacancyID,
		&resume.CurrentStageID,
		&resume.CandidateName,
		&resume.CandidateContact,
		&resume.Source,
		&resume.Description,
		&resume.CreatedAt,
		&resume.UpdatedAt,
		&resume.SLADeadline,
	)
	if err != nil {
		return models.Resume{}, err
	}
	return resume, nil
}

func (r *postgresResumeRepository) GetResumes(ctx context.Context, filters models.ResumeFilters, sort models.ResumeSort) ([]models.Resume, error) {
	query := `
	SELECT id, vacancy_id, current_stage_id, candidate_name, candidate_contact, source, description, created_at, updated_at, sladeadline
	FROM resumes
	WHERE 1=1
	`
	args := []interface{}{}
	argCount := 1

	if filters.VacancyID != nil {
		query += fmt.Sprintf(" AND vacancy_id = %d", argCount)
		args = append(args, *filters.VacancyID)
		argCount++
	}
	if filters.StageID != nil {
		query += fmt.Sprintf(" AND current_stage_id = %d", argCount)
		args = append(args, *filters.StageID)
		argCount++
	}
	if filters.DateFrom != nil {
		query += fmt.Sprintf(" AND created_at >= '%d'", argCount)
		args = append(args, *filters.DateFrom)
		argCount++
	}
	if filters.DateTo != nil {
		query += fmt.Sprintf(" AND created_at <= '%d'", argCount)
		args = append(args, *filters.DateTo)
		argCount++
	}
	if sort.Field != "" {
		query += fmt.Sprintf(" ORDER BY %s %s", sort.Field, sort.Direction)
	}
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var resumes []models.Resume
	for rows.Next() {
		var resume models.Resume
		err := rows.Scan(
			&resume.ID,
			&resume.VacancyID,
			&resume.CurrentStageID,
			&resume.CandidateName,
			&resume.CandidateContact,
			&resume.Source,
			&resume.Description,
			&resume.CreatedAt,
			&resume.UpdatedAt,
			&resume.SLADeadline,
		)
		if err != nil {
			return nil, err
		}
		resumes = append(resumes, resume)
	}
	return resumes, nil

}

func (r *postgresResumeRepository) UpdateResumeByID(ctx context.Context, resume models.Resume) error {
	query := `
	UPDATE resumes
	SET vacancy_id = $2,
	current_stage_id = $3,
	candidate_name = $4,
	candidate_contact = $5,
	source = $6,
	description = $7,
	created_at = $8,
	updated_at = $9,
	sladeadline = $10
	WHERE id = $1
	`
	_, err := r.db.Exec(ctx, query,
		resume.ID,
		resume.VacancyID,
		resume.CurrentStageID,
		resume.CandidateName,
		resume.CandidateContact,
		resume.Source,
		resume.Description,
		resume.CreatedAt,
		resume.UpdatedAt,
		resume.SLADeadline,
	)
	if err != nil {
		return err
	}
	return nil

}

func (r *postgresResumeRepository) DeleteResumeByID(ctx context.Context, id int) error {
	query := `
	DELETE FROM resumes
	WHERE id = $1
	`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil

}

func (r *postgresResumeRepository) MoveResumeToStage(ctx context.Context, resumeID, stageID int) error {
	updateQuery := `
	UPDATE resumes
	SET current_stage_id = $2, updated_at = NOW()
	WHERE id = $1
	`
	_, err := r.db.Exec(ctx, updateQuery, resumeID, stageID)
	if err != nil {
		return err
	}
	return nil

}

func (r *postgresResumeRepository) AddResumeHistory(ctx context.Context, resume models.Resume, userID int) error {
	query := `
	INSERT INTO resume_histories (resume_id, stage_id, start_time, end_time, changed_by)
	VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(ctx, query, resume.ID, resume.CurrentStageID, resume.UpdatedAt, time.Now(), userID)
	if err != nil {
		return err
	}
	return nil
}
