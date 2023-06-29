package main

import (
	mailconverter "Indexer-Prueba/mail_converter"
	zincsearch "Indexer-Prueba/zinc_search"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	ruta := os.Args[1]
	var cpuprofile = flag.String("cpuprofile", "cpuProf", "write cpu profile to `file`")
	var memprofile = flag.String("memprofile", "memProf", "write memory profile to `file`")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("No se pudo crear el CPU Profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("No se pudo crear el CPU Profile: ", err)
		}
		defer pprof.StopCPUProfile()

		json := mailconverter.ReadMails(ruta)

		newZS, err := zincsearch.NewZincsearch()
		if err != nil {
			fmt.Println(err)
		}
		err = newZS.UserExists()
		if err != nil {
			fmt.Println(err)
		}

		newZS.SendBulkMails(&json)

		if *memprofile != "" {
			f, err := os.Create(*memprofile)
			if err != nil {
				log.Fatal("No se pudo crear el Memory Profile: ", err)
			}
			defer f.Close()
			runtime.GC()
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatal("No se pudo crear el Memory Profile: ", err)
			}
		}
	}
}
