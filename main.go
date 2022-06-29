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
 var API_SERVER_ADDR = os.Getenv("API_SERVER_ADDR")

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc(API_PREFIX + "hello", h.Hello).Methods(http.MethodGet)


	http.ListenAndServe(API_SERVER_ADDR, router)
	log.Println("Listening on port" + API_SERVER_ADDR)
}
