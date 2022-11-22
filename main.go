package main

import (
	"log"
	"net/http"
)

func winelistHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/winelist", winelistHandler)
	log.Println("API is running")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
