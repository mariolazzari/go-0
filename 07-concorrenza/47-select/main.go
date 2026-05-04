package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}
}
