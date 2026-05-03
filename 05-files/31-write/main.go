package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString("Ciao, Mario!")
	if err != nil {
		log.Fatal(err)
	}
}
