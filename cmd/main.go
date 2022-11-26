package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/syywu/go-api.git/pkg/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/winelist", handlers.GetAllWines).Methods(http.MethodGet)
	router.HandleFunc("/winelist/{id}", handlers.GetWineById).Methods(http.MethodGet)
	router.HandleFunc("/winelist/{id}", handlers.UpdateWine).Methods(http.MethodPatch)
	router.HandleFunc("/winelist", handlers.AddWine).Methods(http.MethodPost)

	log.Println("API is running")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}
