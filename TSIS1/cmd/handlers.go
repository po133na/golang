package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/po133na/golang/pkg/json.go"
)

type Response struct {
	breeds []Breeds `json:"breeds"`
}

func GetBreedByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	breeds := params["breedname"]

	breeds = strings.ToLower(breeds)

	for _, breed := range json.Breeds {
		if strings.ToLower(breed.BreedName) == breeds {
			json.NewEncoder(w).Encode(breeds)
			return
		}
	}
	http.Error(w, "Cat type is not found", http.StatusNotFound)
}

func GetBreedById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	breedId := params["id"]

	for _, breed := range json.Breeds {
		if strconv.Itoa(breed.Id) == breedId {
			json.NewEncoder(w).Encode(breed)
			return
		}
	}

	http.Error(w, "Cat type is not found", http.StatusNotFound)
}

func BreedCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	country := params["country"]

	country = strings.ToLower(country)
	var neededBreed []json.Breeds

	for _, breed := range Breeds {
		if strings.ToLower(breed.Country) == breed {
			neededBreed = append(neededBreed, breed)
		}
	}

	if len(neededBreed) > 0 {
		json.NewEncoder(w).Encode(neededBreed)
	} else {
		http.Error(w, "Any breed from needed country is not found", http.StatusNotFound)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	healthStatus := "App is healthy!\n\n"
	healthCheckResponse := healthStatus
	w.Write([]byte(healthCheckResponse))
}
