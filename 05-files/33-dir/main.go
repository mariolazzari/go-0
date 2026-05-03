package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// creazione dir
	err := os.Mkdir("example_dir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// leggere contenuto dir
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
