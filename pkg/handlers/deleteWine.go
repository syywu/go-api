package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/syywu/go-api.git/pkg/mocks"
)

func DeleteWine(w http.ResponseWriter, r *http.Request) {
	// read dynamic id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// loops over mocks
	for index, wine := range mocks.Wines {
		if wine.Id == id {
			// send response and delete book
			mocks.Wines = append(mocks.Wines[:index], mocks.Wines[index+1:]...)
		}
	}
}
