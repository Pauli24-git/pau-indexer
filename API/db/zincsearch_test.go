package db

import (
	"Indexer-Prueba/API/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchQuery(t *testing.T) {
	//Emular ZincSearc respuesta
	expected := MockZSResponse()
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(expected)
	}))
	defer svr.Close()

	z := NewZinc(svr.URL)
	resp, err := z.SearchQuery(models.Credentials{}, models.Search{})

	if err != nil {
		t.Errorf("No deberia dar error, volver a revisar los mocks.")
	}

	for _, src := range resp.Hits.Hits {
		if src.Source.Message_ID != "ABC123" {
			t.Errorf("Se esperaba otro valor para: Message_ID")
		}
		if src.Source.From != "pau@enron.com" {
			t.Errorf("Se esperaba otro valor para: From")
		}
		if src.Source.Subject != "Urgente" {
			t.Errorf("Se esperaba otro valor para: Subject")
		}
	}
}

func MockZSResponse() []byte {
	resp := models.ZSResponse{}

	src := models.Source{Message_ID: "ABC123", From: "pau@enron.com", Subject: "Urgente"}
	hit := models.Hits{Source: src}

	resp.Hits.Hits = append(resp.Hits.Hits, hit)

	jsonMockData, _ := json.Marshal(resp)

	return jsonMockData
}
