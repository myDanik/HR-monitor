package handlers

import (
	"HR-monitor/pkg/enums"
	"HR-monitor/pkg/models"
	"HR-monitor/pkg/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type VacancyHandler struct {
	vacancyService *service.VacancyService
}

func NewVacancyHandler(vacancyService *service.VacancyService) *VacancyHandler {
	return &VacancyHandler{
		vacancyService: vacancyService,
	}
}

func (h *VacancyHandler) CreateVacancy(w http.ResponseWriter, r *http.Request) {
	var vacancy models.Vacancy
	err := json.NewDecoder(r.Body).Decode(&vacancy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.vacancyService.CreateVacancy(r.Context(), vacancy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *VacancyHandler) GetVacancyByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	vacancyID, err := strconv.Atoi(id)
	vacancy, err := h.vacancyService.GetVacancyByID(r.Context(), vacancyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vacancy)
}

func (h *VacancyHandler) DeleteVacancy(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	resumeID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	err = h.vacancyService.DeleteVacancy(r.Context(), resumeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *VacancyHandler) ChangeVacancyStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	vacancyID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	statusStr := r.URL.Query().Get("status")
	if statusStr == "" {
		http.Error(w, "Status is required", http.StatusBadRequest)
		return
	}

	status := enums.VacancyStatus(statusStr)
	if status != enums.Open && status != enums.Closed {
		http.Error(w, "Invalid status", http.StatusBadRequest)
		return
	}

	err = h.vacancyService.ChangeVacancyStatus(r.Context(), vacancyID, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
