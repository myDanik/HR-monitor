package models

import (
	"HR-monitor/pkg/enums"
	"time"
)

type Vacancy struct {
	ID          int                 `json:"id" db:"id"`
	Title       string              `json:"title" db:"title"`
	Description string              `json:"description" db:"description"`
	Status      enums.VacancyStatus `json:"status" db:"status"`
	CreatedBy   int                 `json:"created_by" db:"created_by"`
	CreatedAt   time.Time           `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at" db:"updated_at"`
}
