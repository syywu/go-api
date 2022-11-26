package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetWineById(w http.ResponseWriter, r *http.Request) {
	// read id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// iterate over mock.wines
	// if the same then return wine
}
