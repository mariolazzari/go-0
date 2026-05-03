package main

import (
	"fmt"
	"sort"
	"strings"
)

// Definizione di Libro
type Libro struct {
	Titolo string
	Autore string
	Anno   int
}

// Metodo Presentazione per Libro
func (l Libro) Presentazione() string {
	return fmt.Sprintf("Titolo: %s\nAutore: %s\nAnno  : %d\n", l.Titolo, l.Autore, l.Anno)
}

// Interfaccia Biblioteca
type Biblioteca interface {
	AggiungiLibro(libro *Libro)
	PresentaLibri() string
}

// Implementazione concreta
type BibliotecaConcreta struct {
	Libri []*Libro
}

// AggiungiLibro
func (b *BibliotecaConcreta) AggiungiLibro(libro *Libro) {
	b.Libri = append(b.Libri, libro)
}

// PresentaLibri
func (b *BibliotecaConcreta) PresentaLibri() string {
	var presentazioni []string

	for _, libro := range b.Libri {
		presentazioni = append(presentazioni, libro.Presentazione())
	}

	sort.Strings(presentazioni)

	return strings.Join(presentazioni, "\n")
}

// Esempio utilizzo
func main() {
	b := &BibliotecaConcreta{}

	b.AggiungiLibro(&Libro{
		Titolo: "Viaggio al termine della notte",
		Autore: "Celine",
		Anno:   1932,
	})

	b.AggiungiLibro(&Libro{
		Titolo: "1984",
		Autore: "George Orwell",
		Anno:   1949,
	})

	fmt.Println(b.PresentaLibri())
}
