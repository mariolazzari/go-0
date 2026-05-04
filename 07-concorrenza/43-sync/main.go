package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	now := time.Now()
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Elapsed:", time.Since(now))
}
