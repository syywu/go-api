package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Wine struct {
	producerName string `json:"producerName`
	name         string `json:"name`
	region       string `json"region`
	expertRating int    `json:"rating"`
	description  string `json:"desc"`
}

func winelistHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("This is the winelist API")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/winelist", winelistHandler)
	log.Println("API is running")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}
