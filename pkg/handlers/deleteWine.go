package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/syywu/go-api.git/pkg/models"
)

func (h handler) DeleteWine(w http.ResponseWriter, r *http.Request) {
	// read dynamic id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// find first ID then delete thay book
	var wine models.Wine

	if res := h.DB.First(&wine, id); res.Error != nil {
		fmt.Println(res.Error)
	}

	h.DB.Delete(&wine)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")

}
