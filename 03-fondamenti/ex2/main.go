package main

import (
	"fmt"
	"sort"
)

var CatalogoLibri map[string]int

func AggiungiLibro(titolo string, pagine int) {
	CatalogoLibri[titolo] = pagine
}

func RimuoviLibro(titolo string) {
	delete(CatalogoLibri, titolo)
}

func AggiornaLibro(titolo string, pagine int) {
	if _, exists := CatalogoLibri[titolo]; exists {
		CatalogoLibri[titolo] = pagine
	}
}

func StampaCatalogo() {
	var chiavi []string
	for k := range CatalogoLibri {
		chiavi = append(chiavi, k)
	}

	sort.Strings(chiavi)
	for _, key := range chiavi {
		fmt.Printf("%s - %d\n", key, CatalogoLibri[key])
	}
}

func main() {
	CatalogoLibri = make(map[string]int)

	CatalogoLibri["Programmazione Go"] = 375
	CatalogoLibri["Go Avanzato"] = 520
	CatalogoLibri["Go per Principianti"] = 200

	AggiungiLibro("Go Concurrency", 300)
	RimuoviLibro("Go per Principianti")
	AggiornaLibro("Go Avanzato", 530)
	StampaCatalogo()

}
