package main

import (
	"fmt"
	"pack/utenti"
)

func main() {
	utente := utenti.NuovoUtente("Mario", "mario,lazzari@gmail.com")
	fmt.Println(utente.Nome)
}
