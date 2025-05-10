package handlers

import (
	"HR-monitor/pkg/models"
	"HR-monitor/pkg/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type SLAHandler struct {
	SLAService *service.SLAService
}

func NewSLAService(SLAService *service.SLAService) *SLAHandler {
	return &SLAHandler{
		SLAService: SLAService,
	}
}

func (h *SLAHandler) CreateSLARule(w http.ResponseWriter, r *http.Request) {
	var rule models.SLARule
	err := json.NewDecoder(r.Body).Decode(&rule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.SLAService.CreateSLARule(r.Context(), rule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (h *SLAHandler) GetSLARulesByVacancyID(w http.ResponseWriter, r *http.Request) {
	vacancyID := r.URL.Query().Get("vacancy_id")
	if vacancyID == "" {
		http.Error(w, "Vacancy ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(vacancyID)
	if err != nil {
		http.Error(w, "Invalid vacancy ID", http.StatusBadRequest)
		return
	}

	rules, err := h.SLAService.GetSLARulesByVacancyID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rules)

}

func (h *SLAHandler) UpdateSLARule(w http.ResponseWriter, r *http.Request) {
	var rule models.SLARule
	err := json.NewDecoder(r.Body).Decode(&rule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.SLAService.UpdateSLARule(r.Context(), rule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (h *SLAHandler) DeleteSLARule(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	ruleID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	err = h.SLAService.DeleteSLARule(r.Context(), ruleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

