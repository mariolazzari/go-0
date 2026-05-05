package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Order represents a coffee order
type Order struct {
	ID        int
	CreatedAt time.Time
}

// Get random in range
func getRandomTime(from, to int) time.Duration {
	rand := time.Duration(from + rand.Intn(to))
	return rand * time.Millisecond
}

// barista function handles coffee preparation
func barista(id int, orders <-chan Order, ready chan<- int, errs chan<- error) {
	// Defer recovery to handle eventual panics
	defer func() {
		if r := recover(); r != nil {
			errs <- fmt.Errorf("barista %d recovered from panic: %v", id, r)
		}
	}()

	for order := range orders {
		fmt.Printf("Barista %d started order %d\n", id, order.ID)

		// Simulate preparation time between 1 and 3 seconds
		prepTime := getRandomTime(1000, 3000)
		time.Sleep(prepTime)

		// Simulate a random error (10% chance)
		if rand.Intn(10) == 0 {
			errs <- fmt.Errorf("barista %d: spilled coffee for order %d", id, order.ID)
			continue
		}

		// Send ready notification
		ready <- order.ID
	}
}

// customer function sends an order and waits for completion or timeout
func customer(id int, orders chan<- Order, readyChan <-chan int, errs <-chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	order := Order{ID: id, CreatedAt: time.Now()}
	orders <- order

	// Select with 5 seconds timeout
	select {
	case readyID := <-readyChan:
		if readyID == order.ID {
			fmt.Printf("Customer %d: received coffee %d! Happy.\n", id, readyID)
		}
	case err := <-errs:
		fmt.Printf("Customer %d: sorry, error occurred: %v\n", id, err)
	case <-time.After(5 * time.Second):
		fmt.Printf("Customer %d: order %d took too long. Leaving the shop!\n", id, order.ID)
	}
}

func main() {
	// Channels initialization
	orderChan := make(chan Order, 10)
	readyChan := make(chan int, 10)
	errorChan := make(chan error, 10)

	var wg sync.WaitGroup

	// Start 5 baristas
	for i := 1; i <= 5; i++ {
		go barista(i, orderChan, readyChan, errorChan)
	}

	// Simulation of 10 customers arriving at random intervals
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go customer(i, orderChan, readyChan, errorChan, &wg)

		// Random arrival interval between 500ms and 1.5s
		arrivalInterval := getRandomTime(500, 1500)
		time.Sleep(arrivalInterval)
	}

	// Wait for all customers to finish their experience
	wg.Wait()

	close(orderChan)
	fmt.Println("Cafeteria is closing. Goodbye!")
}
