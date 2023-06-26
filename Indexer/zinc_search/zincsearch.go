package zincsearch

import (
	"Indexer-Prueba/models"
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

	nuevoZincSearch := Zincsearch{Username: valueUser, Password: valuePass}

	return &nuevoZincSearch, nil
}

func (z Zincsearch) UserExists() error {
	var response models.LoginResponse
	method := "POST"
	credentials := models.Credentials{Id: z.Username, Password: z.Password}

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

func (z Zincsearch) SendBulkMails(jsonData *[]byte) {
	method := "POST"

	req, err := http.NewRequest(method, "http://localhost:4080/api/_bulkv2", bytes.NewBuffer(*jsonData))
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

	if res.StatusCode == http.StatusOK {
		fmt.Println("Los mails fueron enviados correctamente a Zinc")
	}
}
