package controller

import (
	service "Indexer-Prueba/API/services"
	"encoding/json"
	"log"
	"net/http"
)

func SearchItems(response http.ResponseWriter, request *http.Request) {

	term := request.URL.Query().Get("term")
	field := request.URL.Query().Get("field")

	// term := chi.URLParam(request, "term")
	// field := chi.URLParam(request, "field")

	data, err := service.SendQuery(term, field)
	if err != nil {
		log.Fatalf("Error ")
	}

	json.NewEncoder(response).Encode(data)

}
