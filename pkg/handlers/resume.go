package handlers

import (
	"HR-monitor/pkg/enums"
	"HR-monitor/pkg/models"
	"HR-monitor/pkg/service"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type ResumeHandler struct {
	resumeService *service.ResumeService
}

func NewResumeHandler(resumeService *service.ResumeService) *ResumeHandler {
	return &ResumeHandler{
		resumeService: resumeService,
	}
}

func (h *ResumeHandler) GetResumes(w http.ResponseWriter, r *http.Request) {
	var filters models.ResumeFilters
	var sort models.ResumeSort

	vacancyID := r.URL.Query().Get("vacancy_id")
	if vacancyID != "" {
		id, err := strconv.Atoi(vacancyID)
		if err != nil {
			http.Error(w, "invalid vacancy id", http.StatusBadRequest)
			return
		}
		filters.VacancyID = &id
	}

	stageID := r.URL.Query().Get("stage_id")
	if stageID != "" {
		id, err := strconv.Atoi(stageID)
		if err != nil {
			http.Error(w, "invalid stage id", http.StatusBadRequest)
			return
		}
		filters.StageID = &id
	}

	dateFrom := r.URL.Query().Get("date_from")
	if dateFrom != "" {
		date, err := time.Parse(time.DateOnly, dateFrom)
		if err != nil {
			http.Error(w, "invalid date format", http.StatusBadRequest)
			return
		}
		filters.DateFrom = &date
	}

	dateTo := r.URL.Query().Get("date_to")
	if dateTo != "" {
		date, err := time.Parse(time.DateOnly, dateTo)
		if err != nil {
			http.Error(w, "invalid date format", http.StatusBadRequest)
			return
		}
		filters.DateTo = &date
	}

	sortField := r.URL.Query().Get("sort_field")
	if sortField != "" {
		sort.Field = enums.ResumeSortField(sortField)
	}

	sortDirection := r.URL.Query().Get("sort_direction")
	if sortDirection != "" {
		sort.Direction = enums.ResumeSortDirection(sortDirection)
	}

	resumes, err := h.resumeService.GetResumes(r.Context(), filters, sort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resumes)
}

func (h *ResumeHandler) CreateResume(w http.ResponseWriter, r *http.Request) {
	var resume models.Resume
	err := json.NewDecoder(r.Body).Decode(&resume)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.resumeService.CreateResume(r.Context(), resume)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ResumeHandler) MoveResumeToStage(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ResumeID int `json:"resume_id"`
		StageID  int `json:"stage_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.resumeService.MoveResumeToStage(r.Context(), request.ResumeID, request.StageID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ResumeHandler) GetResumeStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.resumeService.GetResumeStats(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
