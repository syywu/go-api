package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	var wine models.Wine

	if res := h.DB.First(&wine, id); res.Error != nil {
		fmt.Println(res.Error)
	}

	// update and send res back when id matches
	wine.Producer = updatedWine.Producer
	wine.Name = updatedWine.Name
	wine.Region = updatedWine.Region
	wine.Rating = updatedWine.Rating
	wine.Description = updatedWine.Description

	h.DB.Save(&wine)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")

}
