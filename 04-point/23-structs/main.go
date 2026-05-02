package main

import "fmt"

type Persona struct {
	Nome string
	Età  int
}

type Impiegato struct {
	Persona
	Posizione string
}

func (p *Persona) compleanno() {
	p.Età++
}

func main() {
	p := Persona{Nome: "Mario", Età: 51}

	fmt.Println(p.Nome)
	fmt.Println(p.Età)

	p.compleanno()
	fmt.Println(p.Età)

	e := Impiegato{
		Persona: Persona{
			Nome: "Mario",
			Età:  50,
		},
		Posizione: "Programmatore",
	}

	e.compleanno()
	fmt.Println(e.Nome)
	fmt.Println(e.Età)

	// struttura anonima
	punto := struct{ x, y int }{x: 10, y: 10}
	fmt.Println(punto)
}
