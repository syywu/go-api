package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/syywu/go-api.git/initialiser"
	"github.com/syywu/go-api.git/pkg/handlers"
)

func Init() {
	initialiser.LoadEnvVariables()
	initialiser.Connection()
}

func main() {
	DB := initialiser.Connection()
	h := handlers.New(DB)
	router := mux.NewRouter()
	router.HandleFunc("/winelist", h.GetAllWines).Methods(http.MethodGet)
	router.HandleFunc("/winelist/{id}", h.GetWineById).Methods(http.MethodGet)
	router.HandleFunc("/winelist/{id}", h.UpdateWine).Methods(http.MethodPut)
	router.HandleFunc("/winelist", h.AddWine).Methods(http.MethodPost)
	router.HandleFunc("/winelist/{id}", h.DeleteWine).Methods(http.MethodDelete)

	log.Println("API is running")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}
