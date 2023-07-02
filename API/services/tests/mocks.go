package services

import "Indexer-Prueba/API/models"

type AuthHandlerMock struct {
}

func (a *AuthHandlerMock) ValidateAuthDbUser() (models.Credentials, error) {
	c := models.Credentials{}
	return c, nil
}

func (a *AuthHandlerMock) UserExists(credentials models.Credentials) error {
	return nil
}

type DbHandlerMock struct {
}

func (a *DbHandlerMock) SearchQuery(credentials models.Credentials, query models.Search) (models.ZSResponse, error) {
	resp := models.ZSResponse{}

	src := models.Source{Message_ID: "ABC123", From: "pau@enron.com", Subject: "Urgente"}
	hit := models.Hits{Source: src}

	resp.Hits.Hits = append(resp.Hits.Hits, hit)
	return resp, nil
}
