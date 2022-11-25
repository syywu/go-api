package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/syywu/go-api.git/pkg/models"
)

func AddWine(w http.ResponseWriter, r *http.Request) {
	// read req body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var wine models.Wine
	json.Unmarshal(body, &wine)

	// add to wines []
	// send back 201 status code
}
