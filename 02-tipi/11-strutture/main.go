package main

import "fmt"

type Persona struct {
	Nome  string
	Età   int
	Email string
}

func (p Persona) Saluta() string {
	return "Ciao " + p.Nome
}

func main() {
	p := Persona{
		Nome:  "Mario",
		Età:   51,
		Email: "mario.lazzari@gmail.com",
	}

	saluto := p.Saluta()
	fmt.Println(saluto)
}
