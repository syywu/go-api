package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/syywu/go-api.git/initialiser"
	"github.com/syywu/go-api.git/pkg/handlers"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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
	err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
	if err != nil {
		panic(err)
	}

}
