package models

import (
	"HR-monitor/pkg/enums"
	"time"
)

type Resume struct {
	ID               int       `json:"id" db:"id"`
	VacancyID        int       `json:"vacancy_id" db:"vacancy_id"`
	CurrentStageID   int       `json:"current_stage_id" db:"current_stage_id"`
	CandidateName    string    `json:"candidate_name" db:"candidate_name"`
	CandidateContact string    `json:"candidate_contact" db:"candidate_contact"`
	Source           string    `json:"source" db:"source"`
	Description      string    `json:"description" db:"description"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	SLADeadline      time.Time `json:"sla_deadline" db:"sladeadline"`
}

type ResumeSort struct {
	Field     enums.ResumeSortField     `json:"sort_field"`
	Direction enums.ResumeSortDirection `json:"sort_direction"`
}

type ResumeFilters struct {
	VacancyID *int       `json:"vacancy_id"`
	StageID   *int       `json:"stage_id"`
	DateFrom  *time.Time `json:"date_from"`
	DateTo    *time.Time `json:"date_to"`
}

type Stage struct {
	ID    int               `json:"id" db:"id"`
	Stage enums.ResumeStage `json:"stage" db:"stage"`
}

type ResumeHistory struct {
	ID        int       `json:"id" db:"id"`
	ResumeID  int       `json:"resume_id" db:"resume_id"`
	StageID   int       `json:"stage_id" db:"stage_id"`
	StartTime time.Time `json:"start_time" db:"start_time"`
	EndTime   time.Time `json:"end_time" db:"end_time"`
	ChangedBy int       `json:"changed_by" db:"changed_by"`
}

type ResumeStats struct {
	AverageStageTime            float64        `json:"average_stage_time"`
	StageDistribution           map[string]int `json:"stage_distribution"`
	SourceDistribution          map[string]int `json:"source_distribution"`
	AverageCandidatesPerVacancy float64        `json:"average_candidates_per_vacancy"`
	SLAViolationsCount          int            `json:"sla_violations_count"`
}
