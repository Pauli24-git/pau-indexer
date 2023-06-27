package main

import (
	"Indexer-Prueba/API/controller"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Get("/search", controller.SearchItems)

	http.ListenAndServe(":3080", r)
}
