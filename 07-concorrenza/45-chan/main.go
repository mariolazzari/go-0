package main

import (
	"fmt"
	"time"
)

func worker(msg string, ch chan string) {
	time.Sleep(time.Second)
	ch <- msg
}

func main() {
	messages := make(chan string)

	go worker("Lavoro completato", messages)
	msg := <-messages
	fmt.Println(msg)
}
