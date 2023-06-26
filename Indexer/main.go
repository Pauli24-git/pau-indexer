package main

import (
	mailconverter "Indexer-Prueba/mail_converter"
	zincsearch "Indexer-Prueba/zinc_search"
	"fmt"
)

func main() {
	//ruta := os.Args[1]

	json := mailconverter.ReadMails("maildir")

	newZS, err := zincsearch.NewZincsearch()
	if err != nil {
		fmt.Println(err)
	}
	err = newZS.UserExists()
	if err != nil {
		fmt.Println(err)
	}

	newZS.SendBulkMails(&json)
}
