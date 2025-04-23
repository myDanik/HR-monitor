package models

import (
	"HR-monitor/pkg/enums"
	"time"
)

type Resume struct {
	ID int `json:"id"`
	VacancyID int `json:"vacancy_id"`
	CurrentStageID int `json:"current_stage_id"`
	CondidateName string `json:"condidate_name"`
	CondiadateContact string `json:"condiadate_contact"`
	Source string `json:"source"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	SLADeadline time.Duration `json:"sla_deadline"`
}

type ResumeSort struct {
	Field enums.ResumeSortField
	Direction enums.ResumeSortDirection
}

type ResumeFilters struct {
	VacancyID *int `json:"vacancy_id"`
	StageID *int `json:"stage_id"`
	DateFrom *time.Time `json:"date_from"`
	DateTo *time.Time `json:"date_to"`
}

type Stage struct {
	ID int `json:"id"`
	Stage enums.ResumeStage `json:"stage"`
}

type ResumeHistory struct {
	ID int `json:"id"`
	ResumeID int `json:"resume_id"`
	StageID int `json:"stage_id"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	ChangedBy int `json:"changed_by"`
}

