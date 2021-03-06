package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gmanninglive/scrapi/pkg/models"
	"github.com/gmanninglive/scrapi/pkg/scraper"
)

func (h handler) UpdateAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	reviewChan, done := make(chan models.Review), make(chan bool)

	// Run scraper concurrently
	go scraper.RefreshReviews(reviewChan, done)

	// Wait for messages on channels, close when done channel receives message
L:
	for {
		select {
		case review := <-reviewChan:
			if result := h.DB.Create(&review); result.Error != nil {
				w.WriteHeader(http.StatusExpectationFailed)
				json.NewEncoder(w).Encode(result.Error)
			}
			h.DB.Save(&review)
		case <-done:
			close(reviewChan)
			close(done)
			break L
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
