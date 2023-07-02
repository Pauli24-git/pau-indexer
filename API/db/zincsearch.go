package db

import (
	"Indexer-Prueba/API/models"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func NewZinc(url string) ZincSearch {
	c := ZincSearch{url}
	return c
}

type ZincSearch struct {
	url string
}

func (z *ZincSearch) SearchQuery(credentials models.Credentials, query models.Search) (models.ZSResponse, error) {
	method := "POST"
	var ZSResponse models.ZSResponse

	jsonData, err := json.Marshal(query)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return ZSResponse, err
	}

	req, err := http.NewRequest(method, z.url+"/api/mail/_search", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error al realizar la solicitud")
	}

	auth := credentials.Id + ":" + credentials.Password
	basicAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	// Set the Authorization header with Basic Authentication
	req.Header.Set("Authorization", "Basic "+basicAuth)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error al realizar la solicitud HTTP")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error al leer el cuerpo del response")
	}

	if res.StatusCode == http.StatusOK {
		fmt.Println("Se recibio correctamente la respuesta")
	}

	err = json.Unmarshal(data, &ZSResponse)
	if err != nil {
		log.Printf("Error al parsear el archivo JSON")
	}
	return ZSResponse, nil
}
