package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ciao Mario!")
}

func main() {
	http.HandleFunc("/", handlerRoot)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
