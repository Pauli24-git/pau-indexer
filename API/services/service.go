package service

import (
	"Indexer-Prueba/API/db"
	"Indexer-Prueba/API/models"
	"fmt"
)

func CreateSearchQuery(term string, field string) (models.Search, error) {
	sortFields := []string{"-@timestamp"}
	from := 0
	maxResults := 20
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
		fmt.Println(err)
	}

	query, err := CreateSearchQuery(term, field)
	ZSResponse, err := newZS.SearchQuery(query)

	for _, src := range ZSResponse.Hits.Hits {
		data = append(data, src.Source)
	}
	return data, nil
}
