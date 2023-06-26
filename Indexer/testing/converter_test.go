package mailconverter

import (
	mailconverter "Indexer-Prueba/mail_converter"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"testing"
)

func TestGetAllFilePaths(t *testing.T) {

	ruta := filepath.Join(t.TempDir(), "file.txt")

	file1, err := os.Create(ruta)
	defer file1.Close()
	if err != nil {
		t.Fatalf("Error al crear archivo")
	}

	var rutas []string
	mailconverter.GetAllFilePaths(&rutas, ruta)

	if rutas[0] != ruta {
		t.Errorf("Se esperaba otra ruta")
	}

}

func BenchmarkGetAllFilePaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ruta := filepath.Join(b.TempDir(), "file.txt")

		file1, err := os.Create(ruta)
		defer file1.Close()
		if err != nil {
			b.Fatalf("Error al crear archivo")
		}

		var rutas []string
		mailconverter.GetAllFilePaths(&rutas, ruta)
		if rutas[0] != ruta {
			b.Errorf("Se esperaba otra ruta")
		}
	}
}

func TestReadFiles(t *testing.T) {
	var testPath []string
	mockChannel := make(chan map[string]string)
	mockWG := new(sync.WaitGroup)

	mockMail := []string{
		"To: hola.pauli@gmail.com\n" + "From: Pau\n" + "Subject: Recordatorio personal\n" + "X-FileName: Recordatorio",
	}
	dirTemporal, err := ioutil.TempDir("", "testdir")
	if err != nil {
		t.Fatal("Failed to create temporary directory:", err)
	}

	ruta := filepath.Join(dirTemporal, "file"+strconv.Itoa(1))
	_, err = os.Create(ruta)
	if err != nil {
		t.Fatalf("Error al crear los archivos")
	}

	data := []byte(mockMail[0])
	err = ioutil.WriteFile(ruta, data, 0644)

	if err != nil {
		t.Fatalf("Failed to write data to file %s: %s", ruta, err)
	}
	testPath = append(testPath, ruta)

	defer os.RemoveAll(dirTemporal)

	mockWG.Add(1)
	go mailconverter.ReadFiles(testPath, mockChannel, mockWG)

	go func() {
		mockWG.Wait()
		close(mockChannel)
	}()

	for mail := range mockChannel {
		if mail["To"] != "hola.pauli@gmail.com" {
			t.Errorf("Se esperaba otro valor para la Key: To")
		}

		if mail["From"] != "Pau" {
			t.Errorf("Se esperaba otro valor para la Key: From")
		}

		if mail["Subject"] != "Recordatorio personal" {
			t.Errorf("Se esperaba otro valor para la Key: Subject")
		}
		if mail["X-FileName"] != "Recordatorio" {
			t.Errorf("Se esperaba otro valor para la Key: X-FileName")
		}
	}
}

func BenchmarkReadFiles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var testPath []string
		mockChannel := make(chan map[string]string)
		mockWG := new(sync.WaitGroup)

		mockMail := []string{
			"To: hola.pauli@gmail.com\n" + "From: Pau\n" + "Subject: Recordatorio personal\n" + "X-FileName: Recordatorio",
		}
		dirTemporal, err := ioutil.TempDir("", "testdir")
		if err != nil {
			b.Fatal("Failed to create temporary directory:", err)
		}

		ruta := filepath.Join(dirTemporal, "file"+strconv.Itoa(1))
		_, err = os.Create(ruta)
		if err != nil {
			b.Fatalf("Error al crear los archivos")
		}

		data := []byte(mockMail[0])
		err = ioutil.WriteFile(ruta, data, 0644)

		if err != nil {
			b.Fatalf("Failed to write data to file %s: %s", ruta, err)
		}
		testPath = append(testPath, ruta)

		defer os.RemoveAll(dirTemporal)

		mockWG.Add(1)
		go mailconverter.ReadFiles(testPath, mockChannel, mockWG)

		go func() {
			mockWG.Wait()
			close(mockChannel)
		}()

		for mail := range mockChannel {
			if mail["To"] != "hola.pauli@gmail.com" {
				b.Errorf("Se esperaba otro valor para la Key: To")
			}

			if mail["From"] != "Pau" {
				b.Errorf("Se esperaba otro valor para la Key: From")
			}

			if mail["Subject"] != "Recordatorio personal" {
				b.Errorf("Se esperaba otro valor para la Key: Subject")
			}
			if mail["X-FileName"] != "Recordatorio" {
				b.Errorf("Se esperaba otro valor para la Key: X-FileName")
			}
		}
	}
}

func TestDivideInBatches(t *testing.T) {
	var filesListMock *[]string
	size := 2

	filesListMock = &[]string{"file1", "file2", "file3", "file4", "file5"}

	returnData := mailconverter.DivideInBatches(filesListMock, size)

	if len(returnData) != 3 {
		t.Errorf("Se esperaba otro valor. Se obtuvo %d, se esperaba 3", len(returnData))
	}

	if len(returnData[0]) != 2 {
		t.Errorf("Se esperaba otro valor. Se obtuvo %d, se esperaba 2", len(returnData[0]))
	}

	if len(returnData[1]) != 2 {
		t.Errorf("Se esperaba otro valor. Se obtuvo %d, se esperaba 2", len(returnData[1]))
	}

	if len(returnData[2]) != 1 {
		t.Errorf("Se esperaba otro valor. Se obtuvo %d, se esperaba 1", len(returnData[0]))
	}
}

func BenchmarkDivideInBatches(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var filesListMock *[]string
		size := 2

		filesListMock = &[]string{"file1", "file2", "file3", "file4", "file5"}

		returnData := mailconverter.DivideInBatches(filesListMock, size)

		if len(returnData) != 3 {
			b.Errorf("Se esperaba otro valor. Se obtuvo %d, se esperaba 3", len(returnData))
		}

		if len(returnData[0]) != 2 {
			b.Errorf("Se esperaba otro valor. Se obtuvo %d, se esperaba 2", len(returnData[0]))
		}

		if len(returnData[1]) != 2 {
			b.Errorf("Se esperaba otro valor. Se obtuvo %d, se esperaba 2", len(returnData[1]))
		}

		if len(returnData[2]) != 1 {
			b.Errorf("Se esperaba otro valor. Se obtuvo %d, se esperaba 1", len(returnData[0]))
		}
	}
}

func TestReadMails(t *testing.T) {
	var rutas []string
	var mockRecords []map[string]string

	mockMail := []string{
		"To: pauli@gmail.com\n" + "From: Pau\n" + "Subject: Recordatorio personal\n" + "X-FileName: Recordar\n",
		"To: enron@hotmail.com\n" + "From: Rick\n" + "Subject: Expensas\n" + "X-FileName: Importante\n",
	}

	dirTemporal, err := ioutil.TempDir("", "testdir")
	if err != nil {
		t.Fatal("Failed to create temporary directory:", err)
	}
	defer os.RemoveAll(dirTemporal)

	for i := 0; i < len(mockMail); i++ {
		ruta := filepath.Join(dirTemporal, "file"+strconv.Itoa(i))
		_, err = os.Create(ruta)
		if err != nil {
			t.Fatalf("Error al crear los archivos")
		}
		rutas = append(rutas, ruta)
	}

	for i := 0; i < len(mockMail); i++ {
		data := []byte(mockMail[i])
		err = ioutil.WriteFile(rutas[i], data, 0644)
	}

	if err != nil {
		t.Fatalf("Failed to write data to file %s: %s", rutas, err)
	}

	file0 := make(map[string]string)
	file0["To"] = "pauli@gmail.com"
	file0["From"] = "Pau"
	file0["Subject"] = "Recordatorio personal"
	file0["X-FileName"] = "Recordar"
	file0["Content"] = ""

	mockRecords = append(mockRecords, file0)

	file1 := make(map[string]string)
	file1["To"] = "enron@hotmail.com"
	file1["From"] = "Rick"
	file1["Subject"] = "Expensas"
	file1["X-FileName"] = "Importante"
	file1["Content"] = ""

	mockRecords = append(mockRecords, file1)
	var c mailconverter.Converter
	c.Mails.Index = "mail"
	c.Mails.Records = mockRecords
	mockJsonData, err := json.Marshal(c.Mails)

	jsonData := mailconverter.ReadMails(dirTemporal)

	areEqual := bytes.Equal(mockJsonData, jsonData)

	if !areEqual {
		t.Errorf("Se esperaba otro valor")
	}
}

func BenchmarkReadMails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var rutas []string
		var mockRecords []map[string]string

		mockMail := []string{
			"To: pauli@gmail.com\n" + "From: Pau\n" + "Subject: Recordatorio personal\n" + "X-FileName: Recordar\n",
			"To: enron@hotmail.com\n" + "From: Rick\n" + "Subject: Expensas\n" + "X-FileName: Importante\n",
		}

		dirTemporal, err := ioutil.TempDir("", "testdir")
		if err != nil {
			b.Fatal("Failed to create temporary directory:", err)
		}
		defer os.RemoveAll(dirTemporal)

		for i := 0; i < len(mockMail); i++ {
			ruta := filepath.Join(dirTemporal, "file"+strconv.Itoa(i))
			_, err = os.Create(ruta)
			if err != nil {
				b.Fatalf("Error al crear los archivos")
			}
			rutas = append(rutas, ruta)
		}

		for i := 0; i < len(mockMail); i++ {
			data := []byte(mockMail[i])
			err = ioutil.WriteFile(rutas[i], data, 0644)
		}

		if err != nil {
			b.Fatalf("Failed to write data to file %s: %s", rutas, err)
		}

		file0 := make(map[string]string)
		file0["To"] = "pauli@gmail.com"
		file0["From"] = "Pau"
		file0["Subject"] = "Recordatorio personal"
		file0["X-FileName"] = "Recordar"
		file0["Content"] = ""

		mockRecords = append(mockRecords, file0)

		file1 := make(map[string]string)
		file1["To"] = "enron@hotmail.com"
		file1["From"] = "Rick"
		file1["Subject"] = "Expensas"
		file1["X-FileName"] = "Importante"
		file1["Content"] = ""

		mockRecords = append(mockRecords, file1)
		var c mailconverter.Converter
		c.Mails.Index = "mail"
		c.Mails.Records = mockRecords
		mockJsonData, err := json.Marshal(c.Mails)

		jsonData := mailconverter.ReadMails(dirTemporal)

		areEqual := bytes.Equal(mockJsonData, jsonData)

		if !areEqual {
			b.Errorf("Se esperaba otro valor")
		}
	}
}
