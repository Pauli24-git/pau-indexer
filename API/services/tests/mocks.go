package services

import "Indexer-Prueba/API/models"

type AuthHandlerMock struct {
}

func (a *AuthHandlerMock) ValidateAuthDbUser() (models.Credentials, error) {
	c := models.Credentials{}
	c.Id = "pepe"
	c.Password = "pass de pepe"
	return c, nil
}

func (a *AuthHandlerMock) UserExists(credentials models.Credentials) error {
	return nil
}

type DbHandlerMock struct {
}

func (a *AuthHandlerMock) SearchQuery(credentials models.Credentials, query models.Search) (models.ZSResponse, error) {
	resp := models.ZSResponse{}
	src := models.Source{}
	resp.Hits.Source = src
	return resp, nil
}
