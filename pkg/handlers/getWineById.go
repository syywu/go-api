package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/syywu/go-api.git/pkg/models"
)

func (h handler) GetWineById(w http.ResponseWriter, r *http.Request) {
	// read id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var wine models.Wine

	if res := h.DB.First(&wine, id); res.Error != nil {
		fmt.Println(res.Error)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&wine)

}
