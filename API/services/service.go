package service

import (
	"Indexer-Prueba/API/db"
	"Indexer-Prueba/API/models"
	"log"
)

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

func SendQuery(term string, field string) ([]models.Source, error) {
	var data []models.Source
	newZS, err := db.NewZincsearch()
	if err != nil {
		log.Printf("Error con el login a ZincSearch")
	}

	query, err := CreateSearchQuery(term, field)
	if err != nil {
		log.Printf("Error al recibir la consulta")
	}

	ZSResponse, err := newZS.SearchQuery(query)
	if err != nil {
		log.Printf("Error al recibir la respuesta de Zincsearch")
	}

	for _, src := range ZSResponse.Hits.Hits {
		data = append(data, src.Source)
	}
	return data, nil
}
