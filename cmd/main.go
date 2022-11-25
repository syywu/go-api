package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Wine struct {
	Producer    string
	Name        string
	Region      string
	Rating      int
	Description string
}

var Wines = []Wine{
	{
		Producer:    "Olivier Leflaive",
		Name:        "St Aubin 1er Cru Dents de Chien",
		Region:      "France",
		Rating:      92,
		Description: "Gorgeous nose of buttery lemon, orange zest, honeysuckle, vanilla and melon. The texture is crisp and refreshing, with a vanilla caramel density shot through with lime. Creaminess and oak underpin the fruit.",
	},
}

func winelistHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Wines)
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
