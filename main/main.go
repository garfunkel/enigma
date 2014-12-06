package main

import (
	"github.com/garfunkel/enigma"
	"flag"
	"os"
	"io"
	"log"
)

func main() {
	settingsPath := flag.String("s", "settings.json", "path to JSON settings file")

	flag.Parse()

	enigma, err := enigma.New(*settingsPath)

	if err != nil {
		log.Fatal(err)
	}

	if flag.NArg() == 0 {
		io.Copy(enigma, os.Stdin)
	} else {
		var handle *os.File

		for _, path := range flag.Args() {
			handle, err = os.Open(path)

			if err != nil {
				log.Fatal(err)
			}

			io.Copy(enigma, handle)
		}
	}
}
