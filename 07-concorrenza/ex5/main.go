package main

import (
	"fmt"
	"sync"
)

// var counter int = 0
type SafeCounter struct {
	counter int
	mux     sync.Mutex
}

func saluta(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Ciao, %s!\n", name)
}

func main() {
	var wg1, wg2 sync.WaitGroup

	nomi := []string{"Alice", "Bruno", "Carlo", "Davide"}

	for _, nome := range nomi {
		wg1.Add(1)
		go saluta(nome, &wg1)
	}
	wg1.Wait()

	sc := SafeCounter{}
	for range 1000 {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			sc.mux.Lock()
			defer sc.mux.Unlock()
			sc.counter++
		}()
	}

	wg2.Wait()
	fmt.Println("Counter =", sc.counter)
}
