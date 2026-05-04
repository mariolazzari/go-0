package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (sc *SafeCounter) Inc(key string) {
	sc.mux.Lock()
	// solo una goroutine alla volta può accedere a sc.v
	sc.v[key]++
	sc.mux.Unlock()
}

func (sc *SafeCounter) Value(key string) int {
	sc.mux.Lock()
	defer sc.mux.Unlock()

	// solo una goroutine alla volta può accedere a sc.v
	return sc.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	var wg sync.WaitGroup

	/* 	for range 1000 {
	wg.Add(1)
	go func() {
		c.Inc("myKey")
		wg.Done()
	}()
	}*/

	for range 1000 {
		wg.Go(func() {
			c.Inc("myKey")
		})
	}

	wg.Wait()

	fmt.Println(c.v["myKey"])
}
