package main

import "fmt"

type Studente struct {
	Nome      string
	Età       int
	Punteggio int
}

func (s Studente) Saluta(nome string) string {
	return "Ciao, piacere di vederti " + nome
}

func main() {
	mario := Studente{
		Nome:      "Mario",
		Età:       51,
		Punteggio: 10,
	}

	fmt.Println(mario.Saluta(mario.Nome))
	fmt.Println("Età:", mario.Età)
	fmt.Println("Punteggio:", mario.Punteggio)
}
