package db

import (
	"Indexer-Prueba/API/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserExists(t *testing.T) {
	//Emular ZincSearc respuesta
	expected := MockCredentialsResponse()
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(expected)
	}))
	defer svr.Close()

	z := NewZincAuth(svr.URL)
	err := z.UserExists(models.Credentials{})

	if err != nil {
		t.Errorf("No deberia dar error, volver a revisar los mocks.")
	}
}

func MockCredentialsResponse() []byte {
	resp := models.LoginResponse{Validated: true}

	jsonMockData, _ := json.Marshal(resp)

	return jsonMockData
}
