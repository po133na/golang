package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("starting server...")
	router := mux.NewRouter()

	log.Println("Creating routes")
	router.HandleFunc("/breeds", handlers.GetBreeds).Methods("GET")
	router.HandleFunc("/breeds/breed_name/{breed_name}", handlers.GetBreedByName).Methods("GET")
	router.HandleFunc("/breeds/{id}", handlers.GetBreedById).Methods("GET")
	router.HandleFunc("/health-checking", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/breeds/country/{country}", handlers.BreedByCountry).Methods("GET")
	http.Handle("/", router)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
