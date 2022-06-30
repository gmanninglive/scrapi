package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gmanninglive/scrapi/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) GetReviews(w http.ResponseWriter, r *http.Request) {
	var reviews []models.Review
	var limit int = 100
	query := r.URL.Query()

	if query.Has("limit") {
		limit, _ = strconv.Atoi(query.Get("limit"))
	}

	w.Header().Add("Content-Type", "application/json")
	if result := h.DB.Limit(limit).Find(&reviews); result.Error != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(result.Error)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reviews)
}

func (h handler) GetReview(w http.ResponseWriter, r *http.Request) {
	var reviews []models.Review
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	w.Header().Add("Content-Type", "application/json")

	if result := h.DB.Find(&reviews, id); result.Error != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(result.Error)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reviews)
}
