package models

import "time"

type SLARule struct {
	ID            int       `json:"id"`
	VacancyID     int       `json:"vacancy_id"`
	StageID       int       `json:"stage_id"`
	DurationHours int       `json:"duration_hours"`
	CreatedBy     int       `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
