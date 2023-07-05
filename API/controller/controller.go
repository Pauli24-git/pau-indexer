package controller

import (
	"Indexer-Prueba/API/db"
	service "Indexer-Prueba/API/services"
	"html/template"
	"log"
	"net/http"
)

func SearchItems(response http.ResponseWriter, request *http.Request) {
	urlBase := "http://localhost:4080"

	term := request.URL.Query().Get("term")
	field := request.URL.Query().Get("field")

	auth := db.NewZincAuth(urlBase)
	db := db.NewZinc(urlBase)

	s := service.Search_Service{DB: &db, AuthZinc: &auth}

	data, err := s.SendQuery(term, field)
	if err != nil {
		log.Printf("Error al enviar la solicitud")
	}

	if len(data) == 0 {
		tmpl := template.Must(template.ParseFiles("templateNoMatch.html"))
		err = tmpl.Execute(response, "")
		if err != nil {
			log.Printf("Error al ejecutar la visualizacion del template")
		}
	} else {
		tmpl := template.Must(template.ParseFiles("templates.html"))
		for _, mail := range data {
			err = tmpl.Execute(response, mail)
			if err != nil {
				log.Printf("Error al recorrer los mails")
			}
		}
	}
}

func MethodNotAllowed(response http.ResponseWriter, request *http.Request) {
	http.Error(response, "Method Not Allowed", 405)
}

func NotFound(response http.ResponseWriter, request *http.Request) {
	http.Error(response, "Page not found", 404)
}
