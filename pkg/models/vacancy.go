package models

import (
	"HR-monitor/pkg/enums"
	"time"
)

type Vacancy struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status enums.VacancyStatus `json:"status"`
	CreatedBy int `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

