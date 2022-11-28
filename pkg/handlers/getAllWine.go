package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/syywu/go-api.git/pkg/models"
)

func (h handler) GetAllWines(w http.ResponseWriter, r *http.Request) {
	var wine []models.Wine

	if res := h.DB.Find(&wine); res.Error != nil {
		fmt.Println(res.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&wine)
}
