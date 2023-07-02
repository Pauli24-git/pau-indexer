package services

import (
	service "Indexer-Prueba/API/services"
	"testing"
)

func TestSendQuery(t *testing.T) {

	authTrucho := AuthHandlerMock{}
	dbTrucha := DbHandlerMock{}

	serviceConTruchadas := service.Search_Service{DB: &dbTrucha, AuthZinc: &authTrucho}
	respuesta, err := serviceConTruchadas.SendQuery("", "")

	if err != nil {
		t.Errorf("No deberia dar error")
	}

	for _, r := range respuesta {
		if r.Message_ID != "ABC123" {
			t.Errorf("Se esperaba otro valor para: Message_ID")
		}
		if r.From != "pau@enron.com" {
			t.Errorf("Se esperaba otro valor para: From")
		}
		if r.Subject != "Urgente" {
			t.Errorf("Se esperaba otro valor para: Urgente")
		}
	}
}
