package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/syywu/go-api.git/pkg/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/winelist", handlers.GetAllWines)
	log.Println("API is running")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}
