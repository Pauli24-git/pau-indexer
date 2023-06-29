package mailconverter

import (
	"Indexer-Prueba/models"
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Converter struct {
	Mails       models.Records
	RecordMutex sync.Mutex
}

func ReadMails(ruta string) []byte {
	var c Converter
	c.Mails.Index = "mails"
	var filesList []string
	ch := make(chan map[string]string)
	wg := new(sync.WaitGroup)

	err := GetAllFilePaths(&filesList, ruta)

	if err != nil {
		fmt.Println("Error al obtener los archivos")
	}

	batches := DivideInBatches(&filesList, 10000)

	for _, file := range batches {
		wg.Add(1)
		go ReadFiles(file, ch, wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	c.RecordMutex.Lock()
	for mail := range ch {
		c.Mails.Records = append(c.Mails.Records, mail)
	}
	c.RecordMutex.Unlock()
	jsonData, err := json.Marshal(c.Mails)

	if err != nil {
		fmt.Println("Error al convertir los mails a Json")
	}
	return jsonData
}

func DivideInBatches(filesList *[]string, size int) [][]string {
	batchSize := size
	batches := make([][]string, 0, (len(*filesList)+batchSize-1)/batchSize)

	for batchSize < len(*filesList) {
		*filesList, batches = (*filesList)[batchSize:], append(batches, (*filesList)[0:batchSize:batchSize])
	}
	batches = append(batches, *filesList)

	return batches
}

func ReadFiles(paths []string, ch chan map[string]string, wg *sync.WaitGroup) {

	camposValidos := "To_From_Cc_Subject_Bcc_X-Filename_Message-ID_Date_Mime-Version_Content_Type_Content-Transfer-Encoding_X-From_X-To_X-cc_X-bcc_X-Folder_X-Origin_X-FileName"

	defer wg.Done()
	for _, file := range paths {
		mapita := make(map[string]string)
		file, err := os.Open(file)
		if err != nil {
			fmt.Println("Error abriendo los archivos")
		}
		defer file.Close()
		var value string
		var key string
		var keyAnterior string

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			linea := scanner.Text()
			if key == "Content" {
				mapita[key] += linea
				continue
			}
			limite := ":"
			primerTermino, segundoTermino, tieneSeparador := strings.Cut(linea, limite)

			if tieneSeparador {
				if strings.Contains(camposValidos, primerTermino) {
					key = primerTermino
					value = segundoTermino
				} else {
					key = keyAnterior
					mapita[key] += linea
					continue
				}
			} else {
				key = keyAnterior
				mapita[key] += linea
				continue
			}

			mapita[strings.TrimSpace(key)] = strings.TrimSpace(value)
			keyAnterior = key

			if keyAnterior == "X-FileName" {
				mapita["Content"] = ""
				key = "Content"
			}

		}
		ch <- mapita
	}
}

func GetAllFilePaths(list *[]string, ruta string) error {

	err := filepath.Walk(ruta, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			*list = append(*list, path)
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
