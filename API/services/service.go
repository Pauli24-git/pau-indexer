package service

import (
	"Indexer-Prueba/API/models"
	"fmt"
	"log"
)

// Esta funcion recibe como parametros dos strings (term y field). Con estos datos, crea la consulta de busqueda con el nombre 'query' que es del tipo un models.Search. Devuelve 'query' y error.
func CreateSearchQuery(term string, field string) (models.Search, error) {
	sortFields := []string{"-@timestamp"}
	const from = 0
	const maxResults = 20
	sourceFields := []string{}

	query := models.Search{
		Search_type: "matchphrase",
		Query: models.QueryObj{
			Term:  term,
			Field: field,
		},
		Sort_fields: sortFields,
		From:        from,
		Max_results: maxResults,
		Source:      sourceFields,
	}
	return query, nil
}

type Search_Service struct {
	AuthZinc models.ZincAuthHandler
	DB       models.DBHandler
}

// Funcion del tipo Search_Service. Esta funcion recibe dos strings(term y field). Dentro de la misma se evalua la validacion del usuario, se crea la query de busqueda, se obtiene el resultado de la query y se guarda en un objeto del tipo []models.Source
func (s Search_Service) SendQuery(term string, field string) ([]models.Source, error) {
	var data []models.Source
	credentials, err := s.AuthZinc.ValidateAuthDbUser()
	if err != nil {
		log.Printf("Error con el login a ZincSearch")
	}

	query, err := CreateSearchQuery(term, field)
	if err != nil {
		log.Printf("Error al recibir la consulta")
	}

	ZSResponse, err := s.DB.SearchQuery(credentials, query)
	if err != nil {
		log.Printf("Error al recibir la respuesta de Zincsearch")
	}

	if ZSResponse.Took == 0 {
		fmt.Println("No se encontró un valor que coincida con la busqueda. Vuelva a intentar")
	}

	for _, src := range ZSResponse.Hits.Hits {
		data = append(data, src.Source)
	}
	return data, nil
}
