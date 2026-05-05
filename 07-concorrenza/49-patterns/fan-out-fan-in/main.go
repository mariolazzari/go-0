package main

import (
	"fmt"
)

// La funzione calcolo prende un id intero e un canale out in cui inviare i dati.
// Calcola il prodotto di id e i, poi invia il risultato nel canale out.
func calcolo(id int, out chan<- int) {
	for i := range 5 {
		out <- id * i
	}
}

func main() {
	// Creazione di due canali, c1 e c2.
	c1 := make(chan int)
	c2 := make(chan int)

	// Avvio di due goroutine separate. Ognuna esegue la funzione calcolo
	// con diversi valori di id e invia i risultati a un canale separato.
	// Questo è un esempio di "fan-out".
	go calcolo(2, c1)
	go calcolo(3, c2)

	// Creazione di un terzo canale, c.
	c := make(chan int)

	// Avvio di una terza goroutine che riceve i risultati dai canali c1 e c2.
	// Non appena un risultato è disponibile in uno dei due canali,
	// lo invia al canale c. Questo è un esempio di "fan-in".
	go func() {
		for {
			select {
			case v := <-c1:
				c <- v
			case v := <-c2:
				c <- v
			}
		}
	}()

	// Il ciclo for legge e stampa i primi 10 valori dal canale c.
	for range 10 {
		fmt.Println(<-c)
	}

}
