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
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	valueUser, boolUser := os.LookupEnv("username")
	if boolUser {
		if valueUser == "" {
			return nil, fmt.Errorf("failed to load .env file: %w", err)
		}
	} else {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	valuePass, boolPass := os.LookupEnv("password")
	if boolPass {
		if valuePass == "" {
			return nil, fmt.Errorf("username not found in environment variable")
		}
	} else {
		return nil, fmt.Errorf("password not found in environment variables")
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

	req, err := http.NewRequest(method, "http://localhost:4080/api/login", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Fatal(err)
	}

	if !response.Validated {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	auth := z.Username + ":" + z.Password
	basicAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	// Set the Authorization header with Basic Authentication
	req.Header.Set("Authorization", "Basic "+basicAuth)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		fmt.Println("Se recibio correctamente la respuesta")
	}

	err = json.Unmarshal(data, &ZSResponse)
	if err != nil {
		log.Fatal(err)
	}
	return ZSResponse, nil
}
