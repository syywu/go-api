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

var Wines = []Wine{
	{
		producerName: "Olivier Leflaive",
		name:         "St Aubin 1er Cru Dents de Chien",
		region:       "France",
		expertRating: 92,
		description:  "Gorgeous nose of buttery lemon, orange zest, honeysuckle, vanilla and melon. The texture is crisp and refreshing, with a vanilla caramel density shot through with lime. Creaminess and oak underpin the fruit.",
	},
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
