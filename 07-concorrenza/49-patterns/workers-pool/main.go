package main

import (
	"fmt"
	"time"
)

// ogni worked ha un id univico, un canale per ricevere lavori e uno per inviare risultati
func worker(id int, lavori <-chan int, risultati chan<- int) {
	// Il ciclo for continua ad estrarre i lavori dal canale "lavori" finché non è vuoto e chiuso.
	for j := range lavori {
		fmt.Println("worker", id, "lavoro iniziato", j)
		// Simula un lavoro lungo un secondo.
		time.Sleep(time.Second)
		fmt.Println("worker", id, "lavoro finito", j)
		// Invia il risultato del lavoro al canale "risultati".
		risultati <- j * 2
	}
}

func main() {
	const numLav = 5
	// Crea un canale per i lavori e un canale per i risultati.
	// Ogni canale ha spazio per "numLav" elementi.
	lavori := make(chan int, numLav)
	risultati := make(chan int, numLav)

	// Avvia tre worker, identificati da ID 1, 2, 3.
	for w := 1; w <= 3; w++ {
		// Ogni goroutine ha il suo ID unico "w" e condivide i canali "lavori" e "risultati".
		go worker(w, lavori, risultati)
	}

	// Inserisce i lavori nel canale "lavori".
	for j := 1; j <= numLav; j++ {
		lavori <- j
	}

	close(lavori)

	for a := 1; a <= numLav; a++ {
		<-risultati
	}
}
