package controller

import (
	service "Indexer-Prueba/API/services"
	"html/template"
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
	tmpl := template.Must(template.ParseFiles("templates.html"))

	for _, mail := range data {
		err = tmpl.Execute(response, mail)
		if err != nil {
			log.Fatalf("Error al recorrer los mails")
		}
	}
}
