package main

import (
	"Indexer-Prueba/API/controller"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/search", controller.SearchItems)

	http.ListenAndServe(":3080", r)
}
