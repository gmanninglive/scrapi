package handlers

import (
	"encoding/json"
	"net/http"
)

func (h handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hello World!")
}
