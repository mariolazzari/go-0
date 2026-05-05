package main

import (
	"fmt"
	"time"
)

// funzione che potrebbe fallire
func performTask(i int, errCh chan<- error) {
	// simuliamo un errore per i pari
	if i%2 == 0 {
		errCh <- fmt.Errorf("Errore con l'input %d", i)
		return
	}

	// Se non ci sono errori, simuliamo un operazione con un ritardo
	time.Sleep(2 * time.Second)
	fmt.Printf("Task %d completato con successo\n", i)
}

func main() {
	numTasks := 5
	errCh := make(chan error, numTasks)

	for i := range numTasks {
		go performTask(i, errCh)
	}

	for i := range numTasks {
		select {
		case err := <-errCh:
			if err != nil {
				fmt.Printf("Ricevuto un errore: %s\n", err)
			}
		case <-time.After(5 * time.Second):
			fmt.Printf("Task %d non ha completato in tempo\n", i)
		}
	}
}
