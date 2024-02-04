package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = ":8080"

func main() {
	log.Println("starting server...")
	r := mux.NewRouter()

	r.HandleFunc("/health-check", HealthCheck)
	r.HandleFunc("/breeds", BreedCountry)
	r.HandleFunc("/breeds", GetBreedById)
	r.HandleFunc("/breeds", GetBreedByName)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
