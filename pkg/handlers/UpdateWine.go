package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/syywu/go-api.git/pkg/mocks"
	"github.com/syywu/go-api.git/pkg/models"
)

func (h handler) UpdateWine(w http.ResponseWriter, r *http.Request) {
	// read dynamic id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//read req body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedWine models.Wine
	json.Unmarshal(body, &updatedWine)

	// iterates over wines
	for index, wine := range mocks.Wines {
		if wine.Id == id {
			// update and send res back when id matches
			wine.Producer = updatedWine.Producer
			wine.Name = updatedWine.Name
			wine.Region = updatedWine.Region
			wine.Rating = updatedWine.Rating
			wine.Description = updatedWine.Description
			// update mocks
			mocks.Wines[index] = wine

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")

		}
	}
}
