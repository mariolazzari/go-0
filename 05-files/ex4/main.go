package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Crea una nuova directory chiamata "Compito" e controlla sempre gli errori.
	err := os.Mkdir("Compito", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Crea un nuovo file chiamato "testo.txt" all'interno della directory "Compito".
	file, err := os.Create("Compito/testo.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Scrivi la frase "Buon lavoro con Go!" nel file "testo.txt".
	// Ricorda di chiudere il file dopo aver completato la scrittura.
	_, err = file.WriteString("Buon lavoro con Go!")
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Leggi il contenuto del file "testo.txt" e stampa il contenuto sul terminale.
	data, err := os.ReadFile("Compito/testo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	// Leggi e stampa i nomi di tutti i file presenti nella directory "Compito".
	// Aiutati con ciclo for per iterare sui file.
	files, err := os.ReadDir("Compito")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
