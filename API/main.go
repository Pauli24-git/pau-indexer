package main

import (
	"Indexer-Prueba/API/controller"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/search", controller.SearchItems)
	r.MethodNotAllowed(controller.MethodNotAllowed)
	r.NotFound(controller.NotFound)

	err := http.ListenAndServe(":3080", r)
	if err != nil {
		log.Fatal(err)
	}
}
