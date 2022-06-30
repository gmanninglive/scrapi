package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gmanninglive/scrapi/pkg/db"
	"github.com/gmanninglive/scrapi/pkg/handlers"
	"github.com/gorilla/mux"
)

var API_PREFIX = "/api/v1/"
var __ADDR = ":" + os.Getenv("PORT")

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc(API_PREFIX+"hello", h.Hello).Methods(http.MethodGet)

	// Fetch and update db with reviews // Currently this appends records to the db. Need to work on replacing duplicates
	router.HandleFunc(API_PREFIX+"refresh", h.UpdateAll).Methods(http.MethodGet)

	// Get Reviews
	router.HandleFunc(API_PREFIX+"reviews", h.GetReviews).Methods(http.MethodGet)
	router.HandleFunc(API_PREFIX+"review/{id}", h.GetReview).Methods(http.MethodGet)

	log.Printf("Listening on port %s!\n", os.Getenv("PORT"))
	http.ListenAndServe(__ADDR, router)
}
