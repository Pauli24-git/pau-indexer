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
	"os"

	"github.com/joho/godotenv"
)

type Zincsearch struct {
	Username string
	Password string
}

func NewZincsearch() (*Zincsearch, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error al cargar el archivo .env: %w", err)
	}

	valueUser, boolUser := os.LookupEnv("username")
	if boolUser {
		if valueUser == "" {
			return nil, fmt.Errorf("Error al cargar el archivo .env: %w", err)
		}
	} else {
		return nil, fmt.Errorf("Error al cargar el archivo .env: %w", err)
	}

	valuePass, boolPass := os.LookupEnv("password")
	if boolPass {
		if valuePass == "" {
			return nil, fmt.Errorf("USERNAME no encontrado dentro de las variables de entorno")
		}
	} else {
		return nil, fmt.Errorf("Error: PASSWORD no encontrada dentro de las variables de entorno")
	}

	err = UserExists(valueUser, valuePass)
	if err != nil {
		return nil, err
	}

	nuevoZincSearch := Zincsearch{Username: valueUser, Password: valuePass}

	return &nuevoZincSearch, nil
}

func UserExists(username string, password string) error {
	var response models.LoginResponse
	method := "POST"
	credentials := models.Credentials{Id: username, Password: password}

	jsonData, err := json.Marshal(credentials)
	if err != nil {
		log.Printf("Error en la conversion a JSON")
	}

	req, err := http.NewRequest(method, "http://localhost:4080/api/login", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error al realizar la solicitud")
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error al realizar la solicitud HTTP")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error al leer el cuerpo del response")
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Printf("Error al parsear el archivo JSON")
	}

	if !response.Validated {
		log.Printf("Error, el usuario no es válido")
	}
	return nil
}

func (z Zincsearch) SearchQuery(query models.Search) (models.ZSResponse, error) {
	method := "POST"
	var ZSResponse models.ZSResponse

	jsonData, err := json.Marshal(query)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return ZSResponse, err
	}

	req, err := http.NewRequest(method, "http://localhost:4080/api/mail/_search", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error al realizar la solicitud")
	}

	auth := z.Username + ":" + z.Password
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
