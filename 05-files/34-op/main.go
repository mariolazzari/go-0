package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Mkdir("testDir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("testDir/testFile.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString("Ciao, Mario!")
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile("testDir/testFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
