package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type wine struct {
	producerName string
	name         string
	region       string
	expertRating int
	description  string
}

func winelistHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("This is the winelist API")
}

func main() {
	http.HandleFunc("/winelist", winelistHandler)
	log.Println("API is running")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
