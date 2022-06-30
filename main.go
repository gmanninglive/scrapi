package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gmanninglive/scrapi/pkg/db"
	"github.com/gmanninglive/scrapi/pkg/handlers"
	"github.com/gorilla/mux"
)


 var API_PREFIX = "/api/v1/"
 var PORT = os.Getenv("PORT")

func main() {
	if !strings.HasPrefix(PORT, ":"){
		PORT = fmt.Sprintf(":%s", PORT)
	}
	
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc(API_PREFIX + "hello", h.Hello).Methods(http.MethodGet)

	// Fetch and update db with reviews // Currently this appends records to the db. Need to work on replacing duplicates
	router.HandleFunc(API_PREFIX + "refresh", h.UpdateAll).Methods(http.MethodGet)

	// Get Reviews
	router.HandleFunc(API_PREFIX + "reviews", h.GetReviews).Methods(http.MethodGet)
	router.HandleFunc(API_PREFIX + "review/{id}", h.GetReview).Methods(http.MethodGet)

	log.Println("Listening on port" + PORT)
	http.ListenAndServe(PORT, router)
}
