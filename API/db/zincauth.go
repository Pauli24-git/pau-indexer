package db

import (
	"Indexer-Prueba/API/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func NewZincAuth(url string) ZincAuthHandler {
	z := ZincAuthHandler{url}
	return z
}

type ZincAuthHandler struct {
	url string
}

func (z *ZincAuthHandler) ValidateAuthDbUser() (models.Credentials, error) {
	err := godotenv.Load()
	cred := models.Credentials{}
	if err != nil {
		return cred, fmt.Errorf("Error al cargar el archivo .env: %w", err)
	}

	valueUser, boolUser := os.LookupEnv("username")
	if boolUser {
		if valueUser == "" {
			return cred, fmt.Errorf("Error al cargar el archivo .env: %w", err)
		}
	} else {
		return cred, fmt.Errorf("Error al cargar el archivo .env: %w", err)
	}

	valuePass, boolPass := os.LookupEnv("password")
	if boolPass {
		if valuePass == "" {
			return cred, fmt.Errorf("USERNAME no encontrado dentro de las variables de entorno")
		}
	} else {
		return cred, fmt.Errorf("Error: PASSWORD no encontrada dentro de las variables de entorno")
	}
	cred.Id = valueUser
	cred.Password = valuePass
	err = z.UserExists(cred)
	if err != nil {
		return cred, err
	}

	return cred, nil
}

func (z *ZincAuthHandler) UserExists(credentials models.Credentials) error {
	var response models.LoginResponse
	method := "POST"

	jsonData, err := json.Marshal(credentials)
	if err != nil {
		log.Printf("Error en la conversion a JSON")
	}

	req, err := http.NewRequest(method, z.url+"/api/login", bytes.NewBuffer(jsonData))
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
