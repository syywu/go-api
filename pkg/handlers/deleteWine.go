package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/syywu/go-api.git/pkg/mocks"
)

func (h handler) DeleteWine(w http.ResponseWriter, r *http.Request) {
	// read dynamic id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// loops over mocks
	for index, wine := range mocks.Wines {
		if wine.Id == id {
			// send response and delete book
			mocks.Wines = append(mocks.Wines[:index], mocks.Wines[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
		}
	}
}
