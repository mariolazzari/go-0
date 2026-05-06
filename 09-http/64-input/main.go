package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHanlder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	nome := r.FormValue("nome")
	email := r.FormValue("email")

	// validazione
	if nome == "" || email == "" {
		fmt.Fprintf(w, "Nome e email sono obbligatori")
	}

	fmt.Fprintf(w, "Nome: %s\nEmail: %s", nome, email)
}

func main() {
	http.HandleFunc("/", formHanlder)
	http.ListenAndServe(":8080", nil)
}
