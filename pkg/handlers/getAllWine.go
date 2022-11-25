package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/syywu/go-api.git/pkg/mocks"
)

func GetAllWines(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Wines)
}