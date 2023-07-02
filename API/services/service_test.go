package service

import (
	"Indexer-Prueba/API/models"
	"testing"
)

func TestSendQuery(t *testing.T) {

	authMock := models.AuthHandlerMock{}
	dbMock := models.DbHandlerMock{}

	ServiceWithMock := Search_Service{DB: &dbMock, AuthZinc: &authMock}
	response, err := ServiceWithMock.SendQuery("", "")

	if err != nil {
		t.Errorf("No deberia dar error, volver a revisar los mocks.")
	}

	for _, r := range response {
		if r.Message_ID != "ABC123" {
			t.Errorf("Se esperaba otro valor para: Message_ID")
		}
		if r.From != "pau@enron.com" {
			t.Errorf("Se esperaba otro valor para: From")
		}
		if r.Subject != "Urgente" {
			t.Errorf("Se esperaba otro valor para: Subject")
		}
	}
}
